package freeipa

import (
	"fmt"
)

var operationStatusHeader = []string{"ID", "Status", "Success", "Failure", "Error", "StartTime", "EndTime"}

type freeIpaOutOperation struct {
	ID        string          `json:"ID" yaml:"ID"`
	Status    string          `json:"Status" yaml:"Status"`
	Success   []successDetail `json:"Success" yaml:"Success"`
	Failure   []failureDetail `json:"Failure" yaml:"Failure"`
	Error     string          `json:"Error,omitempty" yaml:"Error,omitempty"`
	StartTime string          `json:"StartTime" yaml:"StartTime"`
	EndTime   string          `json:"EndTime,omitempty" yaml:"EndTime,omitempty"`
}

type successDetail struct {
	Environment string `json:"Environment" yaml:"Environment"`
}

type failureDetail struct {
	Environment string `json:"Environment" yaml:"Environment"`
	Details     string `json:"Details" yaml:"Details"`
}

func (f *freeIpaOutOperation) DataAsStringArray() []string {
	var successString string
	for _, success := range f.Success {
		successString += fmt.Sprintf("%s\n", success.Environment)
	}
	var failureString string
	for _, failure := range f.Failure {
		failureString += fmt.Sprintf("Environment: %s\n", failure.Environment)
		failureString += fmt.Sprintf("Details: %s\n\n", failure.Details)
	}
	return []string{f.ID, f.Status, successString, failureString, f.Error, f.StartTime, f.EndTime}
}
