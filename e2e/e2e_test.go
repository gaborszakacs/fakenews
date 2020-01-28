// +build e2e

package e2e_test

import (
	"os"
	"os/exec"
	"testing"
)

func TestE2E(t *testing.T) {
	err := os.Chdir("..")
	if err != nil {
		t.Fatalf("Failed to changed dir: %v", err)
	}
	cmd := exec.Command("go", "run", "main.go")
	output, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("Failed to run: %v: %s", err, string(output))
	}

	got := string(output)
	want := `Climate change helped spawn East Africa’s locust crisis
Hundreds of Amazon workers criticize the company's climate policy
`
	if got != want {
		t.Errorf("got:\n%s\nwant:\n%s\n", got, want)
	}
}
