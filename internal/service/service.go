package service

import (
	"dupfinder/internal/models"
	"log/slog"
	"math/rand"
	"os"
	"runtime"
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
	var (
		info models.SystemInfo
	)

	info.Arch = runtime.GOARCH
	info.OS = runtime.GOOS

	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	info.HostName = host

	//cpu, err := getCPUInfo()
	//if err != nil {
	//	return nil, err
	//}

	_, err = getSystemInfo()
	if err != nil {
		s.Log.Error("error while getting system information: ", "error", err)

		return nil, err
	}

	return &info, nil
}

func (s *Service) CurrentUsage(args ...string) (*models.Usage, error) {
	return &models.Usage{CPU: models.Measure{
		Value: (rand.Float64() * 90) + 10,
		UOM:   "FLOP/s",
	}, RAM: models.Measure{
		Value: (rand.Float64() * 90) + 10,
		UOM:   "MT/s",
	}, GPU: models.Measure{
		Value: (rand.Float64() * 90) + 10,
		UOM:   "ms",
	}}, nil
}
