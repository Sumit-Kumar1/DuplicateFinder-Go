package service

import (
	"testing"
)

const testFailed = "TEST[%d] Failed - %s\nExpected: %+v\nGot:%+v"

func Test_runCommand(t *testing.T) {
	tests := []struct {
		name    string
		cmd     string
		args    []string
		wantErr error
	}{
		{name: "random cmd", cmd: "pwsh", args: []string{"Get-WmiObject", "-Class", "Win32_Processor"}, wantErr: nil},
		{name: "g++ version", cmd: "pwsh", args: []string{"g++", "--version"}},
	}
	for i, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := runCommand(tt.cmd, tt.args...)

			if err != tt.wantErr {
				t.Errorf(testFailed, i, tt.name, tt.wantErr, err)
			}
		})
	}
}
