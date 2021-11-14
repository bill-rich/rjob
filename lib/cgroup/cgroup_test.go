// +build unit

package cgroup

import (
	"testing"
)

type testCase struct {
	name           string
	cgroupConfig   CgroupConfig
	expectedResult string
	expectedError  bool
}

func TestGetCpu(t *testing.T) {
	testCases := []testCase{
		{name: "CpuMax", cgroupConfig: CgroupConfig{Cpu: 50}, expectedResult: "50000 100000", expectedError: false},
		{name: "CpuMax100", cgroupConfig: CgroupConfig{Cpu: 100}, expectedResult: "MAX 100000", expectedError: false},
		{name: "CpuMaxOverMax", cgroupConfig: CgroupConfig{Cpu: 110}, expectedResult: "", expectedError: true},
		{name: "CpuMaxUnderMin", cgroupConfig: CgroupConfig{Cpu: 0}, expectedResult: "", expectedError: true},
	}
	for _, tcase := range testCases {
		result, err := tcase.cgroupConfig.getCpuMax()
		testCgroupCase(t, tcase, result, err)
	}
}

func TestGetMem(t *testing.T) {
	testCases := []testCase{
		{name: "MemMax", cgroupConfig: CgroupConfig{Memory: 2}, expectedResult: "2048", expectedError: false},
		{name: "MemMaxUnderMin", cgroupConfig: CgroupConfig{Memory: -10}, expectedResult: "", expectedError: true},
	}
	for _, tcase := range testCases {
		result, err := tcase.cgroupConfig.getMemMax()
		testCgroupCase(t, tcase, result, err)
	}
}

func TestGetIo(t *testing.T) {
	testCases := []testCase{
		{name: "IoMax", cgroupConfig: CgroupConfig{Io: 50}, expectedResult: "default 50", expectedError: false},
		{name: "IoMaxUnderMin", cgroupConfig: CgroupConfig{Io: 9}, expectedResult: "", expectedError: true},
		{name: "IoMaxOverMax", cgroupConfig: CgroupConfig{Io: 110}, expectedResult: "", expectedError: true},
	}
	for _, tcase := range testCases {
		result, err := tcase.cgroupConfig.getIoWeight()
		testCgroupCase(t, tcase, result, err)
	}
}
func testCgroupCase(t *testing.T, test testCase, result string, err error) {
	t.Helper()
	switch {
	case test.expectedError && err == nil:
		t.Errorf("test %s: error expected, but was not received", test.name)
		return
	case !test.expectedError && err != nil:
		t.Errorf("test %s: received error: %s", test.name, err)
		return
	}
	if result != test.expectedResult {
		t.Errorf("test %s: expected: %s, got: %s", test.name, test.expectedResult, result)
	}
}
