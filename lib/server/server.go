package server

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"net"

	"github.com/bill-rich/rjob/lib/api"
	"github.com/bill-rich/rjob/lib/cgroup"
	"github.com/bill-rich/rjob/lib/command"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ServerConfig struct {
	CaLocation   string
	KeyLocation  string
	CertLocation string

	ListenAddress string
	ListenPort    string
}

func (server *ServerConfig) StartServer() error {
	log.SetLevel(log.DebugLevel)
	log.Debugf("Starting job server at:%s", net.JoinHostPort(server.ListenAddress, server.ListenPort))

	cgroup.Mount(cgroup.CgroupMountDir)
	defer cgroup.Umount(cgroup.CgroupMountDir)

	// TODO: Use common.GetCreds() function.
	certificate, err := tls.LoadX509KeyPair(server.CertLocation, server.KeyLocation)
	if err != nil {
		return err
	}

	rootCA, err := ioutil.ReadFile(server.CaLocation)
	if err != nil {
		log.Fatalf("failed to read client ca cert: %s", err)
	}

	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM(rootCA)
	if !ok {
		log.Fatal("failed to append client certs")
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	}

	lis, err := net.Listen("tcp", net.JoinHostPort(server.ListenAddress, server.ListenPort))
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	serverOption := grpc.Creds(credentials.NewTLS(tlsConfig))
	grpcServer := grpc.NewServer(serverOption)

	jobServer := api.ApiServer{
		Jobs: map[string]command.JobConfig{},
	}

	api.RegisterJobsServer(grpcServer, &jobServer)
	log.Debugf("Job server started successfully")
	grpcServer.Serve(lis)
	return nil
}
