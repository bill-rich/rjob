package command

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"github.com/bill-rich/rjob/lib/cgroup"
	log "github.com/sirupsen/logrus"
)

// TODO: Add logging to goroutine
// TODO: Figure out re-execing processes for cgroup
// TODO: Redirect/capture stderr

const (
	OutRefresh       = 1
	JobStatusStopped = "STOPPED"
	JobStatusRunning = "RUNNING"
)

type JobConfig struct {
	Command      string
	Args         []string
	CgroupName   string
	CgroupConfig *cgroup.CgroupConfig
	Owner        string

	Cmd *exec.Cmd

	Output   *OutputBuffer
	Status   string
	ExitCode int
}

// Wrap creates the required cgroup for the new job, and uses rjob's reexec
// option to start a job within a cgroup to avoid any escaping.
func (job *JobConfig) Start() error {
	job.CgroupConfig.Create()

	self, err := os.Executable()
	if err != nil {
		return fmt.Errorf("cannot get path of rjob executable: %s", err)
	}
	originalCommand := job.Command
	job.Command = self
	job.Args = append([]string{"reexec", job.CgroupName, originalCommand}, job.Args...)
	log.Debug("Wrapping command as ", job.Command, " ", strings.Join(job.Args, " "))

	if err := job.Run(); err != nil {
		return err
	}

	return nil
}

// Run will execute a command in a job. If CgroupName is included in the config,
// the current process will first move itself to that cgroup. This is to ensure
// that no processes escape the cgroup.
func (job *JobConfig) Run() error {
	if job.CgroupName != "" {
		if err := job.ChangeCgroup(); err != nil {
			return fmt.Errorf("unable to move process into cgroup: %s", err)
		}
	}

	// Kick off a goroutine to run the job and update the output.
	go func() {
		job.Cmd = exec.Command(job.Command, job.Args...)
		// Allowing which namespaces to use would be a nice future feature.
		job.Cmd.SysProcAttr = &syscall.SysProcAttr{Cloneflags: syscall.CLONE_NEWPID | syscall.CLONE_NEWNS | syscall.CLONE_NEWNET}

		job.Output = &OutputBuffer{}

		out, _ := job.Cmd.StdoutPipe()
		err := job.Cmd.Start()
		if err != nil {
			log.Error(err)
		}
		job.Status = JobStatusRunning

		for {
			b := make([]byte, 1)
			_, err := out.Read(b)
			job.Output.Write(b)
			if err != nil {
				break
			}
		}

		log.Infof("Job (%s) has finished.", job.CgroupName)
		return
	}()

	// TODO: Replace this with an actual check that job has started.
	time.Sleep(OutRefresh * time.Second)

	return nil
}

// Kill kills the job.
func (job *JobConfig) Kill() error {
	return job.Cmd.Process.Kill()
}

// ChangeCgroup moves the current process into the specified cgroup.
func (job *JobConfig) ChangeCgroup() error {
	pid := os.Getpid()
	cgroupProcs := filepath.Join(cgroup.CgroupMountDir, job.CgroupName, "cgroup.procs")
	f, err := os.OpenFile(cgroupProcs, os.O_APPEND|os.O_WRONLY, 0555)
	if err != nil {
		return fmt.Errorf("unable to open %s: %s", cgroupProcs, err)
	}
	defer f.Close()

	_, err = f.WriteString(fmt.Sprintf("%d", pid))
	if err != nil {
		return fmt.Errorf("unable to write PID to %s: %s", cgroupProcs, err)
	}
	return nil
}

// UpdateStatus checks if the job is still running. If it is not, the exit code,
// and status are updated. The method is only called when the jobs are being
// examined.
func (job *JobConfig) UpdateStatus() {
	if !job.IsRunning() {
		job.Status = JobStatusStopped
		job.ExitCode = job.Cmd.ProcessState.ExitCode()
	}
}

// IsRunning returns true if the job is still running.
func (job *JobConfig) IsRunning() bool {
	if out, err := job.Cmd.StdoutPipe(); err != nil && out != nil {
		return true
	}
	return false
}

// PrintJPrintJobOutput will print job output as it becomes available.
func (job *JobConfig) PrintJobOutput() {
	bytesCurrent := 0
	var newData string
	for job.Status == JobStatusRunning || job.Output.HasNew(bytesCurrent) {
		newData, bytesCurrent = job.Output.Read(bytesCurrent)
		if newData != "" {
			fmt.Print(newData)
			time.Sleep(OutRefresh * time.Second)
		}
	}
}
