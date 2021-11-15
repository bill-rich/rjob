package common

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"

	"github.com/bill-rich/rjob/lib/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func StartServer() error {
	caCert := "ssl/ca.crt"
	serverCert := "ssl/server.crt"
	serverKey := "ssl/server.key"

	//caCert := "badssl/ca-cert.pem"
	//serverCert := "badssl/server-cert.pem"
	//serverKey := "badssl/server-key.pem"

	certificate, err := tls.LoadX509KeyPair(
		serverCert,
		serverKey,
	)

	certPool := x509.NewCertPool()
	bs, err := ioutil.ReadFile(caCert)
	if err != nil {
		log.Fatalf("failed to read client ca cert: %s", err)
	}

	ok := certPool.AppendCertsFromPEM(bs)
	if !ok {
		log.Fatal("failed to append client certs")
	}

	lis, err := net.Listen("tcp", "127.0.0.1:9898")
	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	tlsConfig := &tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	}

	serverOption := grpc.Creds(credentials.NewTLS(tlsConfig))
	server := grpc.NewServer(serverOption)

	jobServer := api.ApiServer{
		Jobs:   []api.Job{},
		Config: api.Config{},
	}

	api.RegisterJobsServer(server, &jobServer)

	server.Serve(lis)
	return nil
}
