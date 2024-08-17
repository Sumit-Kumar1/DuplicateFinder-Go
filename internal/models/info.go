package models

type SystemInfo struct {
	HostName *string   `json:"hostname"`
	OS       *string   `json:"os"`
	Arch     *string   `json:"architecture"`
	Disk     *DiskInfo `json:"disk"`
	CPU      *CPUInfo  `json:"cpu"`
	RAM      *RAMInfo  `json:"ram"`
	GPU      *GPUInfo  `json:"gpu"`
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
	Caption           string `json:"Caption"`
	DeviceID          string `json:"DeviceID"`
	Manufacturer      string `json:"Manufacturer"`
	MaxClockSpeed     int    `json:"MaxClockSpeed"`
	Name              string `json:"Name"`
	SocketDesignation string `json:"SocketDesignation"`
}

type GPUInfo struct {
}
