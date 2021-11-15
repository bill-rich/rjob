package common

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"google.golang.org/grpc/credentials"
)

func GetCreds(caPath, certPath, keyPath string) (*credentials.TransportCredentials, error) {
	cert, err := tls.LoadX509KeyPair(certPath, keyPath)
	if err != nil {
		return nil, fmt.Errorf("error loading creds: %v", err)
	}

	rootCA, err := ioutil.ReadFile(caPath)
	if err != nil {
		return nil, fmt.Errorf("error reading CA file (%s): %s", caPath, err)
	}
	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(rootCA) {
		return nil, fmt.Errorf("error loading root CA")
	}

	tlsConfig := tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{cert},
		RootCAs:      certPool,
	}
	creds := credentials.NewTLS(&tlsConfig)
	return &creds, nil
}
