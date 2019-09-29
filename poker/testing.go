package poker

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"os"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func AssertScoreEquals(got int, want int, t *testing.T) {
	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func AssertLeague(got League, wantedLeague League, t *testing.T) {
	t.Helper()
	if !reflect.DeepEqual(got, wantedLeague) {
		t.Errorf("got %v want %v", got, wantedLeague)
	}
}

func AssertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("didn't expect an error but got one, %v", err)
	}
}

func AssertPlayerWin(t *testing.T, store *StubPlayerStore, winner string) {
	t.Helper()
	if len(store.winCalls) != 1 {
		t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
	}
	if store.winCalls[0] != winner {
		t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], winner)
	}
}

func GetLeagueFromResponse(body io.Reader, t *testing.T) League {
	t.Helper()
	var got League
	err := json.NewDecoder(body).Decode(&got)
	if err != nil {
		t.Fatalf("unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}
	return got
}

func CreateTempFile(t *testing.T, initialData string) (*os.File, func()) {
	t.Helper()
	tmpfile, err := ioutil.TempFile("", "db")
	if err != nil {
		t.Fatalf("could not create tmp file %v", err)
	}

	tmpfile.Write([]byte(initialData))
	removeFile := func() {
		tmpfile.Close()
		os.Remove(tmpfile.Name())
	}

	return tmpfile, removeFile
}
