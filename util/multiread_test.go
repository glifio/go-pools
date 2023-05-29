package util

import (
	"errors"
	"testing"
)

func TestMultireadBasic(t *testing.T) {
	// Define some task functions
	succeed := func() (interface{}, error) {
		return "success", nil
	}
	fail := func() (interface{}, error) {
		return nil, errors.New("failure")
	}

	// Test a successful case
	tasks := []TaskFunc{succeed, succeed}
	results, err := Multiread(tasks)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(results) != 2 || results[0] != "success" || results[1] != "success" {
		t.Errorf("Unexpected results: %v", results)
	}

	// Test a case where one task returns an error
	tasks = []TaskFunc{succeed, fail}
	results, err = Multiread(tasks)
	if err == nil || err.Error() != "failure" {
		t.Errorf("Expected error not found: %v", err)
	}
	if results != nil {
		t.Errorf("Results should be nil when error occurs, but got: %v", results)
	}

	// Test a case where all tasks return an error
	tasks = []TaskFunc{fail, fail}
	results, err = Multiread(tasks)
	if err == nil || err.Error() != "failure" {
		t.Errorf("Expected error not found: %v", err)
	}
	if results != nil {
		t.Errorf("Results should be nil when error occurs, but got: %v", results)
	}
}
