package models

type SystemInfo struct {
	HostName string
	OS       string
	CPUArch  string
	CPU      CPUInfo
	Mem      MemoryInfo
	Disks    []DiskInfo
}

type CPUInfo struct {
	ModelName       string
	ModelIdentifier string
	ProcessorName   string
	ProcessorSpeed  float64
	Cores           int
}

type MemoryInfo struct {
	Total uint64 `json:"total"`
}

type DiskInfo struct {
	Device string    `json:"device"`
	Mount  string    `json:"mountPoint"`
	FsType string    `json:"fsType,omitempty"`
	Usage  DiskUsage `json:"usage"`
}

type DiskUsage struct {
	Total       uint64  `json:"total"`
	Used        uint64  `json:"used"`
	Free        uint64  `json:"free"`
	UsedPercent float64 `json:"usedPercent"`
}
