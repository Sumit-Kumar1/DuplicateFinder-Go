package service

import (
	"dupfinder/internal/models"
	"log/slog"
	"os"
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

func getTotalCPU() (*float64, error) {
	vals, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	var total float64 = 0

	for i := range vals {
		total += vals[i]
	}

	return &total, nil
}
