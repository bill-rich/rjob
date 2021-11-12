// +build integration

package cgroup

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"testing"
)

// TestMountFS checks that the cgroup hierarchy is successfully mounted and
// unmounted.
func TestMountFS(t *testing.T) {
	// Check of cgroup is already mounted
	mounted, err := testCgroupMounted(CgroupMountDir)
	if mounted {
		t.Fatalf("cgroup hierarchy is already mounted at %s", CgroupMountDir)
	}
	if err != nil {
		t.Fatal(err)
	}

	// Mount cgroup
	testMount(CgroupMountDir, t)

	// Unmount cgroup
	testUnmount(CgroupMountDir, t)
}

// TestCreateCgroup mounts cgroup hierarchy, creates a new cgroup, verifies the
// limits are properly set, and cleans up.
func TestCreateCgroup(t *testing.T) {
	testMount(CgroupMountDir, t)

	config := CgroupConfig{
		Name:   "testGroup",
		Cpu:    90,
		Memory: 1024,
		Io:     90,
		Path:   fmt.Sprintf("%s/%s", CgroupMountDir, "testGroup"),
	}

	if err := config.Create(); err != nil {
		t.Error(err)
	}

	config.testVerifyCgroupSettings(t)

	config.Delete()
	config.testVerifyCgroupDeleted(t)
	testUnmount(CgroupMountDir, t)

}

func testMount(path string, t *testing.T) {
	if err := Mount(CgroupMountDir); err != nil {
		t.Fatal(err)
	}
	mounted, err := testCgroupMounted(CgroupMountDir)
	if !mounted {
		t.Fatal("cgroup hierarchy failed to mount")
	}
	if err != nil {
		t.Fatal(err)
	}
}

func testUnmount(path string, t *testing.T) {
	if err := Umount(CgroupMountDir); err != nil {
		t.Error(err)
	}

	mounted, err := testCgroupMounted(CgroupMountDir)
	if mounted {
		t.Errorf("cgroup hierarchy is still mounted after unmounting")
	}
	if err != nil {
		t.Error(err)
	}
}

func (config *CgroupConfig) testVerifyCgroupSettings(t *testing.T) {
	config.testVerifyCgroupCpuSetting(t)
}

func (config *CgroupConfig) testVerifyCgroupCpuSetting(t *testing.T) {
	cpuData, err := os.ReadFile(fmt.Sprintf("%s/cpu.max", config.Path))
	if err != nil {
		t.Errorf("unable to read CPU settings: %s", err)
	}
	if strings.Compare(strings.TrimSuffix(string(cpuData), "\n"), config.getCpuMax()) != 0 {
		t.Errorf("cpu.max setting incorrect. expected: %s, got %s", config.getCpuMax(), string(cpuData))
	}
}

func (config *CgroupConfig) testVerifyCgroupMemorySetting(t *testing.T) {
	memData, err := os.ReadFile(fmt.Sprintf("%s/memory.max", config.Path))
	if err != nil {
		t.Errorf("unable to read Memory settings: %s", err)
	}
	if strings.Compare(strings.TrimSuffix(string(memData), "\n"), config.getMemMax()) != 0 {
		t.Errorf("memory.max setting incorrect. expected: %s, got %s", config.getMemMax(), string(memData))
	}
}

func (config *CgroupConfig) testVerifyCgroupIoSetting(t *testing.T) {
	ioData, err := os.ReadFile(fmt.Sprintf("%s/io.weight", config.Path))
	if err != nil {
		t.Errorf("unable to read IO settings: %s", err)
	}
	if strings.Compare(strings.TrimSuffix(string(ioData), "\n"), config.getIoWeight()) != 0 {
		t.Errorf("io.weight setting incorrect. expected: %s, got %s", config.getIoWeight(), string(ioData))
	}
}

func (config *CgroupConfig) testVerifyCgroupDeleted(t *testing.T) {
	_, err := os.Stat(config.Path)
	if !errors.Is(err, os.ErrNotExist) {
		if err == nil {
			t.Errorf("cgroup still exists at %s", config.Path)
		}
		t.Errorf("error checking if cgroup was deleted: %s", err)
	}
}

func testCgroupMounted(path string) (bool, error) {
	_, err := os.Stat(fmt.Sprintf("%s/cgroup.procs", path))
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	}

	return false, fmt.Errorf("error checking if cgroup hierarchy is mounted: %s", err)
}
