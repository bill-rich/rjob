package api

import (
	context "context"
	"fmt"
	"strings"

	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

type ApiServer struct {
	UnimplementedJobsServer
	Config Config
	Jobs   []Job
}

type Config struct {
	// TBD
}

type Job struct {
	Owner    string
	Command  string
	Output   []byte
	Status   string
	ExitCode int
}

func (server *ApiServer) Start(ctx context.Context, input *StartJobInput) (*StartJobResponse, error) {
	command := input.Command
	if !server.AuthorizeJob(ctx, command) {
		return nil, fmt.Errorf("user not authorized")
	}
	cmd := strings.Split(command, " ")
	fmt.Printf("%+v:", cmd)

	resp := StartJobResponse{}
	return &resp, nil
}

func (server *ApiServer) Stop(context.Context, *StopJobInput) (*StopJobResponse, error) {
	return nil, nil
}

func (server *ApiServer) Status(context.Context, *StatusInput) (*StatusResponse, error) {
	return nil, nil
}

func (server *ApiServer) Monitor(*MonitorJobInput, Jobs_MonitorServer) error {
	return nil
}

func (server *ApiServer) AuthorizeJob(ctx context.Context, command string) bool {
	// TBD
	return true
}

func (server *ApiServer) List(ctx context.Context, _ *Empty) (*ListJobsResponse, error) {
	fmt.Printf("running list")
	p, ok := peer.FromContext(ctx)
	if ok {
		tlsInfo := p.AuthInfo.(credentials.TLSInfo)
		fmt.Printf("%+v", tlsInfo.State.PeerCertificates[0].Subject.CommonName)
	}
	resp := &ListJobsResponse{
		JobInfo: []*JobInfo{
			&JobInfo{
				TaskId: "123",
			},
		},
	}
	return resp, nil
}
