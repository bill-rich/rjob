package server

import (
	"context"
	"testing"

	"github.com/bill-rich/rjob/lib/api"
	"github.com/bill-rich/rjob/lib/common"
	"google.golang.org/grpc"
)

const (
	fail    = "fail"
	success = "success"
)

// TODO: Make this test less hacky. Works for testing connectivity
func TestAuthenticationFail(t *testing.T) {
	testStartServer()

	for name, testCase := range map[string][]string{
		"GoodAuthentication": {"../../ssl/ca.crt", "../../ssl/client.crt", "../../ssl/client.key", success},
		"BadAuthentication":  {"../../ssl/ca.crt", "../../ssl/client3.crt", "../../ssl/client3.key", fail},
	} {
		creds, err := common.GetCreds(testCase[0], testCase[1], testCase[2])
		if err != nil {
			t.Fatal(err)
		}

		conn, err := grpc.Dial("localhost:9080", grpc.WithTransportCredentials(*creds))
		if err != nil {
			t.Error(err)
		}
		jobClient := api.NewJobsClient(conn)
		_, err = jobClient.List(context.TODO(), &api.Empty{})

		if err != nil && testCase[3] == success {
			t.Fatalf("%s: expected authentication to succes, but it didn't", name)
		}
		if err == nil && testCase[3] == fail {
			t.Fatalf("%s: expected authentication to fail, but it didn't", name)
		}

	}
}

func testStartServer() {
	srv := ServerConfig{
		CaLocation:   "ssl/ca.crt",
		KeyLocation:  "ssl/server.key",
		CertLocation: "ssl/server.crt",

		ListenAddress: "0.0.0.0",
		ListenPort:    "9080",
	}
	go func() {
		srv.StartServer()
	}()
}
