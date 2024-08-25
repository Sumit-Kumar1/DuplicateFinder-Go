package service

import (
	"dupfinder/internal/models"
	"reflect"
	"testing"
)

func Test_getDiskInfo(t *testing.T) {
	tests := []struct {
		name    string
		want    []models.DiskInfo
		wantErr error
	}{
		{"valid case", nil, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getDiskInfo()
			if err != tt.wantErr {
				t.Errorf("getDiskInfo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDiskInfo() = %v, want %v", got, tt.want)
			}
		})
	}
}
