package cmd

import (
	"os/exec"
	"strings"
)

const (
	binary string = "tmux"
)

func command(commandName string, args ...string) ([]string, error) {
	fullArgs := []string{commandName}

	if len(args) > 0 {
		fullArgs = append(fullArgs, args...)
	}

	output, err := exec.Command(binary, fullArgs...).CombinedOutput()

	splittedOutput := make([]string, 0)

	for _, row := range strings.Split(string(output), "\n") {
		if len(row) > 0 {
			splittedOutput = append(splittedOutput, row)
		}
	}

	return splittedOutput, err
}

func NewSession(name string, args ...string) error {
	fullArgs := []string{"-s", name, "-d"}

	if len(args) > 0 {
		fullArgs = append(fullArgs, args...)
	}

	_, err := command("new-session", fullArgs...)

	return err
}

func RenameSession(oldSession, newSession string) error {
	_, err := command("rename-session", "-t", oldSession, newSession)

	return err
}

func ListSessions() []string {
	names, err := command("list-sessions", "-F", "#{session_name}")

	if err != nil {
		return make([]string, 0)
	}

	return names
}

func KillSession(name string) error {
	_, err := command("kill-session", "-t", name)

	return err
}
