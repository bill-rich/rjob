// +build integration

package command

import (
	"regexp"
	"testing"
	"time"
)

func TestCommand(t *testing.T) {
	config := JobConfig{
		Command:    "echo",
		Args:       []string{"test"},
		CgroupName: "",
	}

	if err := config.Run(); err != nil {
		t.Fatal(err)
	}

	for i := 0; i <= 10 && config.Status == JobStatusRunning; i++ {
		// TODO: Add timeout
		time.Sleep(1 * time.Second)
	}

	output, _ := config.Output.Read(0)
	re := regexp.MustCompile("test")
	if !re.MatchString(output) {
		t.Fatalf("unexpected job output. expected: test, got: %s", output)
	}

}

func TestCommandKill(t *testing.T) {
	config := JobConfig{
		Command:    "sleep",
		Args:       []string{"10"},
		CgroupName: "",
	}

	if err := config.Run(); err != nil {
		t.Fatal(err)
	}

	if err := config.Kill(); err != nil {
		t.Fatal(err)
	}

	for i := 0; i <= 3; i++ {
		config.UpdateStatus()
		if config.Status != JobStatusRunning {
			break
		}
		time.Sleep(OutRefresh * time.Second)
	}

	if config.Status == JobStatusRunning {
		t.Fatalf("job still running after 3 seconds")
	}

}

// TODO: Add test for cgroup placement
// TODO: Add test for namespace placement
