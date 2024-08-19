package service

const (
	powershell           = "pwsh"
	cmd                  = "cmd"
	getWmiObj            = "Get-WmiObject"
	class                = "-Class"
	win32_Processor      = "Win32_Processor"
	win32_DiskDrive      = "Win32_DiskDrive"
	win32_Fan            = "Win32_Fan"
	win32_PhysicalMemory = "Win32_PhysicalMemory"
	registrySysInfo      = `SYSTEM\CurrentControlSet\Control\SystemInformation`
	registryWinInfo      = `SOFTWARE\Microsoft\Windows NT\CurrentVersion`
)
