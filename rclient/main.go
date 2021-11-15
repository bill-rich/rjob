package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	arg "github.com/alexflint/go-arg"
	"github.com/bill-rich/rjob/lib/api"
	"github.com/bill-rich/rjob/lib/common"
	"google.golang.org/grpc"
)

type ClientConfig struct {
	caPath   string
	keyPath  string
	certPath string

	server string // TODO: Split into host and port

	connection *grpc.ClientConn
	jobs       api.JobsClient
}

var args struct {
	Operation   string   `arg:"positional,required"`
	Target      string   `arg:"required"`
	CpuLimit    int      `default:"100"`
	MemoryLimit int      `default:"0"`
	IoLimit     int      `default:"100"`
	Command     string   `arg:"positional"`
	Args        []string `arg:"positional"`
	JobId       string
}

func main() {
	arg.MustParse(&args)
	client := ClientConfig{
		caPath:   "/tmp/rjob/ssl/ca.crt",
		certPath: "/tmp/rjob/ssl/client.crt",
		keyPath:  "/tmp/rjob/ssl/client.key",
		server:   args.Target,
	}

	if err := client.setup(); err != nil {
		fmt.Println(err)
		return
	}
	defer client.connection.Close()

	switch args.Operation {
	case "start":
		client.start()
	case "stop":
		client.stop()
	case "list":
		client.list()
	case "status":
		client.status()
	case "monitor":
		client.monitor()
	}

}

func (client *ClientConfig) start() {
	input := &api.StartJobInput{
		Command: args.Command,
		Args:    args.Args,
		Cpu:     int32(args.CpuLimit),
		Memory:  int32(args.MemoryLimit),
		Blkio:   int32(args.IoLimit),
	}
	jobId, err := client.jobs.Start(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(jobId.JobId)
}

func (client *ClientConfig) stop() {
	if args.JobId == "" {
		log.Fatal("job ID required to stop")
		return
	}
	input := &api.StopJobInput{
		JobId: args.JobId,
	}
	resp, err := client.jobs.Stop(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
		return
	}
	os.Exit(int(resp.ExitCode))
}

func (client *ClientConfig) list() {
	resp, err := client.jobs.List(context.TODO(), &api.Empty{})
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, jobInfo := range resp.JobInfo {
		fmt.Printf("%s:\t%s\texit_code:%d\n", jobInfo.TaskId, jobInfo.Status, jobInfo.ExitCode)
	}
}

func (client *ClientConfig) status() {
	input := &api.StatusInput{
		JobId: args.JobId,
	}
	resp, err := client.jobs.Status(context.TODO(), input)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%s\texit_code:%d", resp.Status, resp.ExitCode)
}

func (client *ClientConfig) monitor() {
	input := &api.MonitorJobInput{
		JobId: args.JobId,
	}
	monitorClient, err := client.jobs.Monitor(context.TODO(), input)
	if err != nil {
		log.Fatal(err)
	}
	for {
		chunk, err := monitorClient.Recv()
		if err != nil {
			if err == io.EOF {
				return
			}
			log.Fatal(err)
		}
		fmt.Print(chunk.Chunk)
		time.Sleep(1 * time.Second)
	}
}

func (client *ClientConfig) setup() error {
	creds, err := common.GetCreds(client.caPath, client.certPath, client.keyPath)
	if err != nil {
		return err
	}

	conn, err := grpc.Dial(client.server, grpc.WithTransportCredentials(*creds))
	if err != nil {
		return fmt.Errorf("error dialing: %v\n", err)
	}
	jobClient := api.NewJobsClient(conn)
	client.connection = conn
	client.jobs = jobClient
	return nil
}
