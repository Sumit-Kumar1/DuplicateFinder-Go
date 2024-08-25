package service

import (
	"crypto/md5"
	"crypto/sha256"
	"dupfinder/internal/models"
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
)

func getTotalCPU() (*float64, error) {
	vals, err := cpu.Percent(0, false)
	if err != nil {
		return nil, err
	}

	var total float64 = 0

	for i := range vals {
		total += vals[i]
	}

	return &total, nil
}

func getDiskInfo() ([]models.DiskInfo, error) {
	var dd []models.DiskInfo

	partitions, err := disk.Partitions(false)
	if err != nil {
		log.Fatal(err)
	}

	for _, partition := range partitions {
		d := models.DiskInfo{
			Device: partition.Device,
			Mount:  partition.Mountpoint,
			FsType: partition.Fstype,
		}

		usage, err := disk.Usage(partition.Mountpoint)
		if err != nil {
			log.Println(err)
			continue
		}

		d.Usage = models.DiskUsage{
			Total:       usage.Total / 1e9,
			Used:        usage.Used / 1e9,
			Free:        usage.Free / 1e9,
			UsedPercent: usage.UsedPercent,
		}

		dd = append(dd, d)
	}

	return dd, nil
}

// calculateHash calculates the MD5 hash of a file
func calculateHash(f *os.File) string {
	hash := md5.New()
	_, err := io.Copy(hash, f)
	if err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(hash.Sum(nil))
}

func generateID() string {
	hash := sha256.Sum256([]byte(fmt.Sprintf("%d", time.Now().UnixNano())))
	return hex.EncodeToString(hash[:])
}
