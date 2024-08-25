package service

import (
	"dupfinder/internal/models"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/mem"
)

type Service struct {
	Log *slog.Logger
}

func New(log *slog.Logger) *Service {
	return &Service{
		Log: log,
	}
}

func (s *Service) SysInfo() (*models.SystemInfo, error) {
	var memDivisor = 1024 * 1024 * 1024

	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	cpuInfo, err := cpu.Info()
	if err != nil {
		return nil, err
	}

	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	disks, err := getDiskInfo()
	if err != nil {
		return nil, err
	}

	systemInfo := &models.SystemInfo{
		HostName: host,
		CPUArch:  runtime.GOARCH,
		OS:       runtime.GOOS,
		CPU: models.CPUInfo{
			ModelName:       cpuInfo[0].ModelName,
			ModelIdentifier: cpuInfo[0].Model,
			ProcessorName:   cpuInfo[0].Family,
			ProcessorSpeed:  cpuInfo[0].Mhz,
			Cores:           int(cpuInfo[0].Cores),
		},
		Mem: models.MemoryInfo{
			Total: v.Total / uint64(memDivisor),
		},
		Disks: disks,
	}

	return systemInfo, nil
}

func (s *Service) CurrentUsage() (*models.Usage, error) {
	val, err := getTotalCPU()
	if err != nil {
		return nil, err
	}

	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}

	return &models.Usage{
		CPU: models.Measure{
			Value: *val,
			UOM:   "%",
		},
		RAM: models.Measure{
			Value: v.UsedPercent,
			UOM:   "%",
		},
		GPU: models.Measure{
			Value: float64(v.Shared) / (1024 * 1024 * 1024),
			UOM:   "%",
		},
	}, nil
}

func (s *Service) Duplicate(device string) (map[string][]models.File, error) {
	files := make(map[string][]models.File) // map to store the files and their hashes

	// Check if the device exists
	if _, err := os.Stat(device); os.IsNotExist(err) {
		s.Log.Error("Device %s does not exist", "device", device)
		return nil, err
	}

	f, err := os.Open(device)
	if err != nil {
		s.Log.Error("Error opening device", "error", err)
		return nil, err
	}

	defer f.Close()

	//logic for finding the duplicate
	err = filepath.Walk(device, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			f, err := os.Open(path)
			if err != nil {
				return err
			}

			defer f.Close()

			// Calculate the hash of the file
			hash := calculateHash(f)

			// Check if the hash is already in the map
			if files[hash] != nil {
				files[hash] = append(files[hash], models.File{ID: generateID(), Name: f.Name(), Path: path, Hash: hash})
			} else {
				files[hash] = []models.File{{Path: path, Hash: hash}}
			}
		}

		return nil
	})
	if err != nil {
		return nil, err
	}

	for key, vals := range files { //delete non duplicate entries from map
		if len(vals) <= 1 {
			delete(files, key)
		}
	}

	return files, nil
}
