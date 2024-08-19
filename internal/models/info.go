package models

type SystemInfo struct {
	HostName           string
	OS                 string
	Arch               string
	BIOSReleaseDate    string
	BIOSVersion        string
	SystemManufacturer string
	SystemProductName  string
}

type RAMInfo struct {
	Attributes           uint32
	Capacity             uint64
	Caption              string
	ConfiguredClockSpeed uint32
	DataWidth            uint16
	Description          string
	DeviceLocator        string
	FormFactor           uint16
	Manufacturer         string
	MaxVoltage           uint32
	MemoryType           uint16
	MinVoltage           uint32
	Model                string
	Name                 string
	PartNumber           string
	SerialNumber         string
	Speed                uint32
	Status               string
	TypeDetail           uint16
}

type DiskInfo struct {
	Caption    string
	Model      string
	Partitions uint32
	Size       uint64
}

type CPUInfo struct {
	DeviceID          string `json:"DeviceID"`
	Manufacturer      string `json:"Manufacturer"`
	MaxClockSpeed     int    `json:"MaxClockSpeed"`
	Name              string `json:"Name"`
	SocketDesignation string `json:"SocketDesignation"`
}

type GPUInfo struct {
}
