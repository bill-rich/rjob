// +build integration

package cgroup

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

// TestMountFS checks that the cgroup hierarchy is successfully mounted and
// unmounted.
func TestMountFS(t *testing.T) {
	// Check of cgroup is already mounted
	mounted, err := testIsCgroupMounted(CgroupMountDir)
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
		Path:   filepath.Join(CgroupMountDir, "testGroup"),
	}

	if err := config.Create(); err != nil {
		t.Error(err)
	}

	testVerifyCgroupSettings(t, config)

	config.Delete()
	testVerifyCgroupDeleted(t, config)
	testUnmount(CgroupMountDir, t)

}

func testMount(path string, t *testing.T) {
	t.Helper()
	if err := Mount(CgroupMountDir); err != nil {
		t.Fatal(err)
	}
	mounted, err := testIsCgroupMounted(CgroupMountDir)
	if !mounted {
		t.Fatal("cgroup hierarchy failed to mount")
	}
	if err != nil {
		t.Fatal(err)
	}
}

func testUnmount(path string, t *testing.T) {
	t.Helper()
	if err := Umount(CgroupMountDir); err != nil {
		t.Error(err)
	}

	mounted, err := testIsCgroupMounted(CgroupMountDir)
	if mounted {
		t.Errorf("cgroup hierarchy is still mounted after unmounting")
	}
	if err != nil {
		t.Error(err)
	}
}

func testVerifyCgroupSettings(t *testing.T, config CgroupConfig) {
	t.Helper()
	testVerifyCgroupSetting(t, filepath.Join(config.Path, cpuMaxFile), config.getCpuMax)
	testVerifyCgroupSetting(t, filepath.Join(config.Path, memoryMaxFile), config.getMemMax)
	testVerifyCgroupSetting(t, filepath.Join(config.Path, ioWeightFile), config.getIoWeight)

}

func testVerifyCgroupSetting(t *testing.T, targetFile string, expectedFunc func() (string, error)) {
	t.Helper()
	expected, err := expectedFunc()
	if err != nil {
		t.Errorf("unable to get expected limits: %s", err)
		return
	}
	data, err := os.ReadFile(targetFile)
	if err != nil {
		t.Errorf("unable to read settings from %s: %s", targetFile, err)
		return
	}
	if bytes.Compare(bytes.TrimSuffix(data, []byte("\n")), []byte(expected)) != 0 {
		t.Errorf("%s setting incorrect. expected: %s, got %s", targetFile, string(expected), string(data))
	}
}

func testVerifyCgroupDeleted(t *testing.T, config CgroupConfig) {
	t.Helper()
	_, err := os.Stat(config.Path)
	if !errors.Is(err, os.ErrNotExist) {
		if err == nil {
			t.Errorf("cgroup still exists at %s", config.Path)
		}
		t.Errorf("error checking if cgroup was deleted: %s", err)
	}
}

func testIsCgroupMounted(path string) (bool, error) {
	_, err := os.Stat(filepath.Join(path, "cgroup.procs"))
	switch {
	case err == nil:
		return true, nil
	case errors.Is(err, os.ErrNotExist):
		return false, nil
	}

	return false, fmt.Errorf("error checking if cgroup hierarchy is mounted: %s", err)
}
