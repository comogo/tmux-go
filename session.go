package tmux

import "github.com/comogo/tmux-go/cmd"

type Session struct {
	Name string
}

func (s Session) String() string {
	return s.Name
}

func (s Session) Kill() error {
	return cmd.KillSession(s.Name)
}

func (s *Session) Rename(name string) {
	err := cmd.RenameSession(s.Name, name)

	if err != nil {
		return
	}

	s.Name = name
}

func (s Session) Exists() bool {
	for _, session := range ListSessions() {
		if session.Name == s.Name {
			return true
		}
	}

	return false
}

func NewSession(name string) (*Session, error) {
	err := cmd.NewSession(name)

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

func ListSessions() []Session {
	names := cmd.ListSessions()

	sessions := make([]Session, len(names))

	for _, name := range names {
		sessions = append(sessions, Session{Name: name})
	}

	return sessions
}
