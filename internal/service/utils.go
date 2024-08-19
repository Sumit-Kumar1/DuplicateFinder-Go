package service

import (
	"dupfinder/internal/models"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"log"
	"os/exec"
)

// getSystemInfo get the system information from registry
func getSystemInfo() (*models.SystemInfo, error) {
	var (
		info models.SystemInfo
		err  error
	)

	return nil, nil
}

func basicInfo() (*models.SystemInfo, error) {
	var info models.SystemInfo

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, registrySysInfo, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}

	defer k.Close()

	info.SystemProductName, _, err = k.GetStringValue("SystemProductName")
	if err != nil {
		return nil, err
	}

	info.SystemManufacturer, _, err = k.GetStringValue("SystemManufacturer")
	if err != nil {
		return nil, err
	}

	info.BIOSVersion, _, err = k.GetStringValue("BIOSVersion")
	if err != nil {
		return nil, err
	}

	info.BIOSReleaseDate, _, err = k.GetStringValue("BIOSReleaseDate")
	if err != nil {
		return nil, err
	}

	return &info, nil
}

func winInfo() (*models.SystemInfo, error) {
	var info models.SystemInfo

	k, err := registry.OpenKey(registry.LOCAL_MACHINE, registryWinInfo, registry.QUERY_VALUE)
	if err != nil {
		return nil, err
	}

	defer k.Close()

	return &info, nil
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
