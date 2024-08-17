package service

import (
	"dupfinder/internal/models"
	"log/slog"
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

	info.Arch = getStrRef(runtime.GOARCH)
	info.OS = getStrRef(runtime.GOOS)

	host, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	info.HostName = &host

	device, err := getInfo("CPU", models.CPUInfo{})
	if err != nil {
		return nil, err
	}

	info.CPU = device

	return &info, nil
}

func (s *Service) CurrentUsage(args ...string) (*models.Usage, error) {
	return &models.Usage{
		CPU: models.Measure{Value: 85.04, UOM: []models.TranslatedText{{Text: "ms", Locale: models.LocaleEN}}},
		RAM: models.Measure{Value: 43.5, UOM: []models.TranslatedText{{Text: "ms", Locale: models.LocaleEN}}},
		GPU: models.Measure{Value: 10.2, UOM: []models.TranslatedText{{Text: "ms", Locale: models.LocaleEN}}},
	}, nil

}
