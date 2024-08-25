package handler

import "dupfinder/internal/models"

type Servicer interface {
	SysInfo() (*models.SystemInfo, error)
	CurrentUsage() (*models.Usage, error)
	Duplicate(string) (map[string][]models.File, error)
}
