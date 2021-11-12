// +build unit

package cgroup

import (
	"testing"
)

func TestGetCpuMax(t *testing.T) {
	config := CgroupConfig{
		Cpu: 50,
	}
	result := config.getCpuMax()
	if result != "50000 100000" {
		t.Errorf("incorrect CPU max value for 50%%. got: %s, expected: 50000 100000", result)
	}
}

func TestGetCpuMax100(t *testing.T) {
	config := CgroupConfig{
		Cpu: 100,
	}
	result := config.getCpuMax()
	if result != "MAX 100000" {
		t.Errorf("incorrect CPU max value for 100%%. got: %s, expected: MAX 100000", result)
	}
}

func TestGetCpuMaxOverMax(t *testing.T) {
	config := CgroupConfig{
		Cpu: 110,
	}
	result := config.getCpuMax()
	if result != "MAX 100000" {
		t.Errorf("incorrect CPU over max value for >100(110%%). got: %s, expected: MAX 100000", result)
	}
}

func TestGetCpuMaxUnderMin(t *testing.T) {
	config := CgroupConfig{
		Cpu: 0,
	}
	result := config.getCpuMax()
	if result != "1000 100000" {
		t.Errorf("incorrect CPU max value for <1(0%%). got: %s, expected: 1000 100000", result)
	}
}

func TestGetIoWeight(t *testing.T) {
	config := CgroupConfig{
		Io: 50,
	}
	result := config.getIoWeight()
	if result != "50" {
		t.Errorf("incorrect IO weight value for 50%%. got: %s, expected: 50", result)
	}
}

func TestGetIoWeightUnderMin(t *testing.T) {
	config := CgroupConfig{
		Io: 9,
	}
	result := config.getIoWeight()
	if result != "10" {
		t.Errorf("incorrect IO weight value for <10(9%%). got: %s, expected: 10", result)
	}
}

func TestGetIoWeightOverMax(t *testing.T) {
	config := CgroupConfig{
		Io: 110,
	}
	result := config.getIoWeight()
	if result != "100" {
		t.Errorf("incorrect IO weight value for >100(110%%). got: %s, expected: 100", result)
	}
}

func TestGetMemMax(t *testing.T) {
	config := CgroupConfig{
		Memory: 2,
	}
	result := config.getMemMax()
	if result != "2048" {
		t.Errorf("incorrect memory max value for 2KB. got: %s, expected: 2048", result)
	}
}

func TestGetMemMaxUnderMin(t *testing.T) {
	config := CgroupConfig{
		Memory: -10,
	}
	result := config.getMemMax()
	if result != "max" {
		t.Errorf("incorrect memory max value for <1(-10). got: %s, expected: max", result)
	}
}
