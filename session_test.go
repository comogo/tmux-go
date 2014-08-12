package tmux

import "testing"

func TestCreateSession(t *testing.T) {
	sessions := ListSessions()

	sessions_count := len(sessions)

	session, err := CreateSession("foo")

	if err != nil {
		t.Fatal("Cannot create a session")
	}

	defer session.Destroy()

	sessions = ListSessions()

	if len(sessions) != sessions_count+1 {
		t.Fatalf("Expect to have %d sessions but found %d", sessions_count+1, len(sessions))
	}

	founded := false
	for _, s := range sessions {
		if s.Name == session.Name {
			founded = true
			break
		}
	}

	if !founded {
		t.Fatalf("Expected to have a session(%s)", session)
	}
}

func TestListSessions(t *testing.T) {
	session, _ := CreateSession("foo")
	defer session.Destroy()

	sessions := ListSessions()

	if len(sessions) < 1 {
		t.Fatal("Expected to have at least one session")
	}

	founded := false
	for _, s := range sessions {
		if s.Name == session.Name {
			founded = true
			break
		}
	}

	if !founded {
		t.Fatalf("Expected to have a session(%s)", session)
	}
}

func TestExists(t *testing.T) {
	if Exists("foo") {
		t.Fatal("Expect to not exists 'foo' session")
	}

	session, _ := CreateSession("foo")
	defer session.Destroy()

	if !Exists("foo") {
		t.Fatal("Expect to exists 'foo' session")
	}
}

func TestSession(t *testing.T) {
	session, _ := CreateSession("foo")
	defer session.Destroy()

	// Test session.Exists()
	if !session.Exists() {
		t.Fatal("Expect to exists 'foo' session")
	}

	// Test session.Rename()
	session.Rename("bar")

	if Exists("foo") {
		t.Fatal("Expect to not exists 'foo' session")
	}

	if !Exists("bar") {
		t.Fatal("Expect to exists 'bar' session")
	}

	// Test session.Destroy
	session.Destroy()

	if session.Exists() {
		t.Fatal("Expect to not exists 'bar' session")
	}
}
