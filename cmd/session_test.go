package cmd

import (
	"testing"

	"github.com/comogo/testify/assert"
)

func kill(name string) {
	KillSession(name)
}

func exists(name string) bool {
	sessionNames := ListSessions()
	found := false

	for _, sessionName := range sessionNames {
		if sessionName == name {
			found = true
			break
		}
	}
	return found
}

func TestNewSession(t *testing.T) {
	sessionName := "foobar"

	assert.Equal(t, exists(sessionName), false)

	err := NewSession(sessionName)
	defer kill(sessionName)

	assert.Equal(t, err, nil)

	err = NewSession(sessionName)

	assert.NotEqual(t, err, nil)
	assert.Equal(t, exists(sessionName), true)
}

func TestRenameSession(t *testing.T) {
	sessionName := "foobar"
	newName := "foobaz"

	NewSession(sessionName)
	defer kill(sessionName)
	defer kill(newName)

	err := RenameSession(sessionName, newName)

	assert.Equal(t, err, nil)
	assert.Equal(t, exists(sessionName), false)
	assert.Equal(t, exists(newName), true)
}

func TestListSessions(t *testing.T) {
	sessionsLength := len(ListSessions())

	assert.Equal(t, exists("foobar"), false)

	NewSession("foobar")
	defer kill("foobar")

	sessions := ListSessions()

	assert.Includes(t, sessions, "foobar")
	assert.Len(t, sessions, sessionsLength+1)
}

func TestKillSession(t *testing.T) {
	NewSession("foobar")
	defer kill("foobar")

	assert.Includes(t, ListSessions(), "foobar")

	KillSession("foobar")

	assert.NotIncludes(t, ListSessions(), "foobar")
}
