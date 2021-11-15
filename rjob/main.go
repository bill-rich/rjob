package main

import (
	"fmt"
	"os"
	"time"

	"github.com/bill-rich/rjob/lib/command"
	"github.com/bill-rich/rjob/rjob/common"
)

func main() {
	switch os.Args[1] {
	case "reexec":
		// TODO: Verify the number of args before assuming they exist.
		if len(os.Args) < 3 {
			fmt.Printf("Not enough arguments to call exec")
			return
		}
		cmdString := os.Args[3]
		cgroup := os.Args[2]
		var args []string
		if len(os.Args) > 3 {
			args = os.Args[4:]
		}
		job := command.JobConfig{
			CgroupName: cgroup,
			Command:    cmdString,
			Args:       args,
		}
		if err := job.Run(); err != nil {
			fmt.Print(err)
			return
		}

		bytesCurrent := 0
		for job.Status == "RUNNING" {
			if len(job.Output) > bytesCurrent {
				fmt.Printf(string(job.Output[bytesCurrent:]))
				bytesCurrent = len(job.Output)
				time.Sleep(1 * time.Second)
			}
		}
		fmt.Printf(string(job.Output[bytesCurrent:]))
	case "start":
		fmt.Printf("error: %s", common.StartServer())
	}
}
