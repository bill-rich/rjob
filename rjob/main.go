package main

import (
	"log"

	arg "github.com/alexflint/go-arg"
	"github.com/bill-rich/rjob/lib/command"
	"github.com/bill-rich/rjob/lib/server"
)

var args struct {
	Operation     string   `arg:"positional,required"`
	ListenAddress string   `default:"0.0.0.0"`
	ListenPort    string   `default:"9080"`
	Cgroup        string   `arg:"positional"`
	Command       string   `arg:"positional"`
	Args          []string `arg:"positional"`
}

func main() {
	arg.MustParse(&args)
	switch args.Operation {
	case "reexec":
		if args.Cgroup == "" || args.Command == "" {
			log.Fatalf("cgroup and command are required to call reexec")
		}
		job := command.JobConfig{
			CgroupName: args.Cgroup,
			Command:    args.Command,
			Args:       args.Args,
		}
		if err := job.Run(); err != nil {
			log.Fatal(err)
		}

		job.PrintJobOutput()
	case "start":
		server := server.ServerConfig{
			CaLocation:   "/tmp/rjob/ssl/ca.crt",
			CertLocation: "/tmp/rjob/ssl/server.crt",
			KeyLocation:  "/tmp/rjob/ssl/server.key",

			ListenAddress: args.ListenAddress,
			ListenPort:    args.ListenPort,
		}

		if err := server.StartServer(); err != nil {
			log.Fatal(err)
		}
	}
}
