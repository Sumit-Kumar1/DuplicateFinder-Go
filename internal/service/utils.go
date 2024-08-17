package service

import (
	"dupfinder/internal/models"
	"fmt"
	"log"
	"os/exec"
)

func getInfo[T models.RAMInfo | models.CPUInfo | models.DiskInfo | models.GPUInfo](device string, data T) (*T, error) {
	var (
		deviceClass string
	)

	switch device {
	case "CPU":
		deviceClass = win32_Processor
	case "RAM":
		deviceClass = win32_PhysicalMemory
	case "Disk":
		deviceClass = win32_DiskDrive
	default:
		deviceClass = win32_Processor
	}

	out, err := runCommand("pwsh", getWmiObj, class, deviceClass)
	if err != nil {
		return nil, err
	}

	valMap := parseOutput(out)

	log.Printf("Output: %+v", valMap)

	return nil, nil
}

func runCommand(cmd string, args ...string) ([]byte, error) {
	command := exec.Command(cmd, "-NoProfile", "-NonInteractive")

	stdin, err := command.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()

		st := ""
		for i := range args {
			st += " " + args[i]
		}

		fmt.Fprintln(stdin, st)
	}()

	out, err := command.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("output: %s\n", out)

	return out, nil
}

func parseOutput(output []byte) map[string]string {
	vals := make(map[string]string)

	//TODO: write parse logic here

	vals["Caption"] = "Intel64 Family 6 Model 158 Stepping 10"
	vals["DeviceID"] = "CPU0"
	vals["Manufacturer"] = "GenuineIntel"
	vals["MaxClockSpeed"] = "2592"
	vals["Name"] = "Intel (R) Core(TM) i7-9750H CPU @ 2.60GHz"
	vals["SocketDesignation"] = "U3E1"

	return vals
}

func getStrRef(str string) *string {
	return &str
}
