package test

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {
	result := m.Run()
	os.Exit(result)
}

func TestRunGame(t *testing.T) {
	data, err := os.Open("res/integration_test_data")

	if err != nil {
		t.Fatalf("Could not open file: %s", err)
	}

	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)
	var lines []string

	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	data.Close()

	for _, line := range lines {

		//

		p := exec.Command("")
		p.Stdin = strings.NewReader(line)
		p.Stdout = os.Stdout
		p.Stderr = os.Stderr

		err := p.Run()

		if err != nil {
			fmt.Println("runed line: ", line)
		}

	}
}
