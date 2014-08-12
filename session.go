package tmux

import (
	"os/exec"
	"strings"
)

type Session struct {
	Name string
}

func (s Session) String() string {
	return s.Name
}

func (s Session) Destroy() error {
	return exec.Command("tmux", "kill-session", "-t", s.Name).Run()
}

func (s *Session) Rename(name string) {
	err := exec.Command("tmux", "rename-session", "-t", s.Name, name).Run()

	if err != nil {
		s.Name = name
	}
}

func (s Session) Exists() bool {
	for _, session := range ListSessions() {
		if session.Name == s.Name {
			return true
		}
	}

	return false
}

func CreateSession(name string) (*Session, error) {
	cmd := exec.Command("tmux", "new-session", "-s", name, "-d")
	err := cmd.Run()

	return &Session{Name: name}, err
}

func Exists(name string) bool {
	for _, session := range ListSessions() {
		if session.Name == name {
			return true
		}
	}

	return false
}
func ListSessions() []*Session {
	output, err := exec.Command("tmux", "list-session", "-F", "#{session_name}").CombinedOutput()

	sessions := make([]*Session, 0)

	if err == nil {
		for _, row := range strings.Split(string(output), "\n") {
			sessions = append(sessions, &Session{Name: row})
		}
	}

	return sessions
}
