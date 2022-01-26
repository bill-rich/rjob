package api

import (
	context "context"
	"fmt"
	"path/filepath"
	"time"

	"github.com/bill-rich/rjob/lib/cgroup"
	"github.com/bill-rich/rjob/lib/command"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/peer"
)

type ApiServer struct {
	UnimplementedJobsServer
	Jobs map[string]command.JobConfig
}

func (server *ApiServer) Start(ctx context.Context, input *StartJobInput) (*StartJobResponse, error) {
	jobId := uuid.New()
	log.Debugf("New job requested. Id: %s", jobId)

	user := getUserFromContext(ctx)

	job := command.JobConfig{
		Command:    input.Command,
		Args:       input.Args,
		Owner:      user,
		CgroupName: jobId.String(),
		CgroupConfig: &cgroup.CgroupConfig{
			Name:   jobId.String(),
			Cpu:    int(input.Cpu),
			Memory: int(input.Memory),
			Io:     int(input.Blkio),
			Path:   filepath.Join(cgroup.CgroupMountDir, jobId.String()),
		},
	}
	if err := job.Start(); err != nil {
		log.Infof("Unable to run job %s: %s", jobId, err)
		return nil, err
	}

	resp := StartJobResponse{
		JobId: jobId.String(),
	}

	server.Jobs[jobId.String()] = job

	log.Debugf("User (%s) has started job %s.", user, jobId)
	return &resp, nil
}

func (server *ApiServer) Stop(ctx context.Context, input *StopJobInput) (*StopJobResponse, error) {
	log.Infof("Stopping job (%s)", input.JobId)
	if !server.AuthorizeJob(ctx, input.JobId) {
		return nil, fmt.Errorf("no job found with id: %s", input.JobId)
	}
	job := server.Jobs[input.JobId]

	if !job.IsRunning() {
		err := job.Kill()
		if err != nil {
			log.Errorf("Error stopping job (%s): %s", input.JobId, err)
			return nil, err
		}
	}

	for i := 0; i < 5 && job.Status == command.JobStatusRunning; i++ {
		job.UpdateStatus()
		time.Sleep(1 * time.Second)
	}

	if job.Status == command.JobStatusRunning {
		return nil, fmt.Errorf("job %s is still running", input.JobId)
	}

	response := &StopJobResponse{
		ExitCode: int32(job.ExitCode),
	}

	return response, nil
}

func (server *ApiServer) Status(ctx context.Context, input *StatusInput) (*StatusResponse, error) {
	log.Infof("Getting status of job (%s)", input.JobId)
	if !server.AuthorizeJob(ctx, input.JobId) {
		return nil, fmt.Errorf("no job found with id: %s", input.JobId)
	}
	job := server.Jobs[input.JobId]
	job.UpdateStatus()
	log.Infof("Job %s: %+v", input.JobId, job)
	response := &StatusResponse{
		Status:   job.Status,
		ExitCode: int32(job.ExitCode),
	}
	return response, nil
}

func (server *ApiServer) Monitor(input *MonitorJobInput, stream Jobs_MonitorServer) error {
	log.Infof("Starting monitoring on job (%s)", input.JobId)
	if !server.AuthorizeJob(stream.Context(), input.JobId) {
		return fmt.Errorf("no job found with id: %s", input.JobId)
	}
	job := server.Jobs[input.JobId]

	bytesCurrent := 0
	for job.Status == command.JobStatusRunning || job.Output.HasNew(bytesCurrent) {
		if job.Output.HasNew(bytesCurrent) {
			var newData string
			newData, bytesCurrent = job.Output.Read(bytesCurrent)
			chunk := &MonitorJobResponse{
				Chunk: newData,
			}
			stream.Send(chunk)
		}
		job.UpdateStatus()
		time.Sleep(1 * time.Second)
	}
	return nil
}

func (server *ApiServer) AuthorizeJob(ctx context.Context, jobId string) bool {
	user := getUserFromContext(ctx)

	job, ok := server.Jobs[jobId]
	if ok && job.Owner == user {
		return true
	}

	log.Infof("User (%s) is not authorized to interact with job (%s)", user, jobId)
	return false
}

func getUserFromContext(ctx context.Context) string {
	p, ok := peer.FromContext(ctx)
	if ok {
		tlsInfo := p.AuthInfo.(credentials.TLSInfo)
		return tlsInfo.State.PeerCertificates[0].Subject.CommonName
	}
	return ""
}

func (server *ApiServer) List(ctx context.Context, _ *Empty) (*ListJobsResponse, error) {
	user := getUserFromContext(ctx)
	ownedJobs := []*JobInfo{}
	for id, job := range server.Jobs {
		if job.Owner == user {
			newJob := &JobInfo{
				TaskId: id,
			}
			ownedJobs = append(ownedJobs, newJob)
		}
	}
	resp := &ListJobsResponse{
		JobInfo: ownedJobs,
	}
	return resp, nil
}
