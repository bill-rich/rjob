package cgroup

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
	"golang.org/x/sys/unix"
)

// Set a few constants for simplicity. These could be read from a configuration.
const (
	CgroupMountDir = "/tmp/rjob/cgroup"
	fileMode       = 0660
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

func (cg *CgroupConfig) setCpuLimit() error {
	cpuFile, err := os.OpenFile(fmt.Sprintf("%s/cpu.max", cg.Path), os.O_APPEND|os.O_WRONLY, fileMode)
	if err != nil {
		return fmt.Errorf("error opening %s/cpu.max: %s", cg.Path, err)
	}
	defer cpuFile.Close()

	if _, err := cpuFile.WriteString(cg.getCpuMax()); err != nil {
		return fmt.Errorf("error writing CPU limit to %s/cpu.max: %s", cg.Path, err)
	}

	return nil
}

func (cg *CgroupConfig) setMemoryLimit() error {
	memoryFile, err := os.OpenFile(fmt.Sprintf("%s/memory.max", cg.Path), os.O_APPEND|os.O_WRONLY, fileMode)
	if err != nil {
		return fmt.Errorf("error opening %s/memory.max: %s", cg.Path, err)
	}
	defer memoryFile.Close()

	// TODO: Change API to accept memory in bytes, or switch both the MB. KB is silly.
	if _, err := memoryFile.WriteString(cg.getMemMax()); err != nil {
		return fmt.Errorf("error writing memory limit to %s/memory.max: %s", cg.Path, err)
	}

	return nil
}

func (cg *CgroupConfig) setBlkIoLimit() error {
	ioFile, err := os.OpenFile(fmt.Sprintf("%s/io.weight", cg.Path), os.O_APPEND|os.O_WRONLY, fileMode)
	if err != nil {
		return fmt.Errorf("error opening %s/io.weight: %s", cg.Path, err)
	}
	defer ioFile.Close()

	if _, err := ioFile.WriteString(cg.getIoWeight()); err != nil {
		return fmt.Errorf("error writing io limit to %s/io.weight: %s", cg.Path, err)
	}

	return nil
}

func (cg *CgroupConfig) getCpuMax() string {
	switch {
	case cg.Cpu < 1:
		log.Infof("Minimum CPU setting is 1, got %d. Setting to 1.", cg.Cpu)
		cg.Cpu = 1
	case cg.Cpu > 100:
		log.Infof("Maximum CPU setting is 100, got %d. Setting to 100.", cg.Cpu)
		cg.Cpu = 100
		return "MAX 100000"
	case cg.Cpu == 100:
		return "MAX 100000"
	}

	cpuMax := 100000
	cpuLimit := float32(cpuMax) * (float32(cg.Cpu) / 100)
	return fmt.Sprintf("%d %d", int(cpuLimit), cpuMax)
}

func (cg *CgroupConfig) getMemMax() string {
	if cg.Memory < 1 {
		if cg.Memory < 0 {
			log.Infof("Minimum memory setting is 0 (no limit), got %d. Setting to 0.", cg.Memory)
			cg.Memory = 0
		}
		return "max"
	}
	return fmt.Sprintf("%d", cg.Memory*1024)
}

func (cg *CgroupConfig) getIoWeight() string {
	switch {
	case cg.Io < 10:
		log.Infof("Minimum IO setting is 10, got %d. Setting to 10.", cg.Io)
		cg.Io = 10
	case cg.Io > 100:
		log.Infof("Maximum IO setting is 100, got %d. Setting to 100.", cg.Io)
		cg.Io = 100
	}
	return fmt.Sprintf("%d", cg.Io)
}
