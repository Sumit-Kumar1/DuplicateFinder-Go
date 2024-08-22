package models

type SystemInfo struct {
	HostName string
	OS       string
	CPUArch  string
	CPU      CPUInfo
	Mem      MemoryInfo
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
