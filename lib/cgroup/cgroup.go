package cgroup

import (
	"fmt"
	"os"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

// Set a few constants for simplicity. These could be read from a configuration.
const (
	CgroupMountDir = "/tmp/rjob/cgroup"
	fileMode       = 0660
	cpuMaxFile     = "cpu.max"
	memoryMaxFile  = "memory.max"
	ioWeightFile   = "io.weight"
)

// CgroupConfig is used to create, delete, and manage a cgroup.
type CgroupConfig struct {
	Name string

	Cpu    int
	Memory int
	Io     int

	Path string
}

// Mount will mount the cgroup hierarchy at mountDir.
func Mount(mountDir string) error {
	log.Debugf("Mounting cgroup hierarchy at %s", mountDir)

	if err := os.MkdirAll(mountDir, fileMode); err != nil {
		return fmt.Errorf("error creating cgroup mount directory %s: %s", mountDir, err)
	}

	if err := unix.Mount("none", mountDir, "cgroup2", 0, ""); err != nil {
		os.Remove(mountDir)
		return fmt.Errorf("error mounting cgroup2 at %s: %s", mountDir, err)
	}

	log.Debugf("Cgroup hierarchy mounted successfully")
	return verifyCgroupRequirements()
}

// Umount unmounts the cgroup hierarchy mounted at mountDir.
func Umount(mountDir string) error {
	log.Debugf("Unmounting cgroup hierarchy at %s", mountDir)

	if err := unix.Unmount(mountDir, 0); err != nil {
		return fmt.Errorf("unable to unmount cgroup fs: %s", err)
	}

	log.Debugf("Cgroup hierarchy unmounted successfully")
	return nil
}

func verifyCgroupRequirements() error {
	// TODO: Check that required subsystems are enabled, control files exist, etc.
	return nil
}

// Create will create the configured cgroup and set the limits for cpu, memory,
// and io.
func (cg *CgroupConfig) Create() error {
	log.Debugf("Creating cgroup %s", cg.Name)

	if err := os.Mkdir(cg.Path, fileMode); err != nil {
		return fmt.Errorf("error making cgroup directory %s: %s", cg.Path, err)
	}

	// Set cgroup limits
	if err := cg.setCpuLimit(); err != nil {
		return err
	}
	if err := cg.setMemoryLimit(); err != nil {
		return err
	}
	if err := cg.setBlkIoLimit(); err != nil {
		return err
	}

	log.Debugf("Cgroup %s created successfully", cg.Name)
	return nil
}

// Delete will delete the configured cgroup by removing the directory. Processes
// int the cgroup must be killed or moved before deleting.
func (cg *CgroupConfig) Delete() error {
	log.Debugf("Deleting cgroup %s", cg.Name)

	if err := os.Remove(cg.Path); err != nil {
		return fmt.Errorf("unable to remove cgroup %s: %s", cg.Name, err)
	}
	log.Debugf("Cgroup %s deleted successfully", cg.Name)

	return nil
}

func setCgroupLimit(path, limit string) error {
	cgFile, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, fileMode)
	if err != nil {
		return fmt.Errorf("error opening %s: %s", path, err)
	}
	defer cgFile.Close()

	if _, err := cgFile.WriteString(limit); err != nil {
		return fmt.Errorf("error writing limit to %s: %s", path, err)
	}

	return nil
}

func (cg *CgroupConfig) setCpuLimit() error {
	cpuMax, err := cg.getCpuMax()
	if err != nil {
		return err
	}
	return setCgroupLimit(filepath.Join(cg.Path, cpuMaxFile), cpuMax)
}

func (cg *CgroupConfig) setMemoryLimit() error {
	memMax, err := cg.getMemMax()
	if err != nil {
		return err
	}
	return setCgroupLimit(filepath.Join(cg.Path, memoryMaxFile), memMax)
}

func (cg *CgroupConfig) setBlkIoLimit() error {
	ioWeight, err := cg.getIoWeight()
	if err != nil {
		return err
	}
	return setCgroupLimit(filepath.Join(cg.Path, ioWeightFile), ioWeight)
}

func (cg *CgroupConfig) getCpuMax() (string, error) {
	switch {
	case cg.Cpu < 1:
		return "", fmt.Errorf("minimum CPU setting is 1, got %d", cg.Cpu)
	case cg.Cpu > 100:
		return "", fmt.Errorf("maximum CPU setting is 100, got %d", cg.Cpu)
	case cg.Cpu == 100:
		return "MAX 100000", nil
	}

	cpuMax := 100000
	cpuLimit := float32(cpuMax) * (float32(cg.Cpu) / 100)
	return fmt.Sprintf("%d %d", int(cpuLimit), cpuMax), nil
}

func (cg *CgroupConfig) getMemMax() (string, error) {
	if cg.Memory < 1 {
		if cg.Memory < 0 {
			return "", fmt.Errorf("minimum memory setting is 0 (no limit), got %d", cg.Memory)
		}
		return "max", nil
	}
	return fmt.Sprintf("%d", cg.Memory*1024), nil
}

func (cg *CgroupConfig) getIoWeight() (string, error) {
	switch {
	case cg.Io < 10:
		return "", fmt.Errorf("minimum IO setting is 10, got %d", cg.Io)
	case cg.Io > 100:
		return "", fmt.Errorf("maximum IO setting is 100, got %d", cg.Io)
	}
	return fmt.Sprintf("default %d", cg.Io), nil
}
