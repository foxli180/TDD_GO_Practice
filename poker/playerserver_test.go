package poker

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetPlayers(t *testing.T) {

	store := StubPlayerStore{
		map[string]int{
			"Pepper": 20,
			"Floyd":  10,
		}, nil, nil,
	}
	server := NewPlayerServer(&store)

	t.Run("Request Pepper's score", func(t *testing.T) {
		request := newGetScoreRequest("Pepper")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "20"
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)

	})

	t.Run("Request Floyd's score", func(t *testing.T) {
		request := newGetScoreRequest("Floyd")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		got := response.Body.String()
		want := "10"
		assertStatus(t, response.Code, http.StatusOK)
		assertResponseBody(t, got, want)
	})

	t.Run("returns 404 on missing players", func(t *testing.T) {
		request := newGetScoreRequest("Apollo")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		got := response.Code
		want := http.StatusNotFound

		assertStatus(t, got, want)
	})

}

func TestStoreWins(t *testing.T) {
	store := StubPlayerStore{
		map[string]int{}, nil, nil,
	}

	server := NewPlayerServer(&store)

	t.Run("it records wins on POST", func(t *testing.T) {
		player := "Pepper"
		request := newPostWinRequest(player)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusAccepted)
		AssertPlayerWin(t, &store, player)

	})
}

func TestLeague(t *testing.T) {

	t.Run("it returns 200 on /league", func(t *testing.T) {
		wantedLeague := League{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(&store)

		request := newLeagueRequest()
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := GetLeagueFromResponse(response.Body, t)
		assertStatus(t, response.Code, http.StatusOK)
		AssertLeague(got, wantedLeague, t)
		assertContentType(t, response, jsonContentType)
	})
}

func TestFileSystemStore(t *testing.T) {
	t.Run("league from a reader", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)
		got := store.GetLeague()
		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}
		AssertLeague(got, want, t)
	})
	t.Run("/league read twice", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)
		got := store.GetLeague()
		want := League{
			{"Chris", 33},
			{"Cleo", 10},
		}
		AssertLeague(got, want, t)
		got = store.GetLeague()
		AssertLeague(got, want, t)

	})
	t.Run("get player score", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)
		got := store.GetPlayerScore("Chris")
		want := 33
		AssertScoreEquals(got, want, t)
	})

	t.Run("store wins for existing players", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)
		store.RecordWin("Chris")
		got := store.GetPlayerScore("Chris")
		want := 34
		AssertScoreEquals(got, want, t)
	})

	t.Run("store wins for new players", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
			{"Name": "Cleo", "Wins": 10},
			{"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, err := NewFileSystemPlayerStore(database)
		AssertNoError(t, err)
		store.RecordWin("Pepper")
		got := store.GetPlayerScore("Pepper")
		want := 1
		AssertScoreEquals(got, want, t)
	})

	t.Run("league sorted", func(t *testing.T) {
		database, cleanDatabase := CreateTempFile(t, `[
        {"Name": "Cleo", "Wins": 10},
        {"Name": "Chris", "Wins": 33}]`)
		defer cleanDatabase()
		store, _ := NewFileSystemPlayerStore(database)

		got := store.GetLeague()

		want := []Player{
			{"Chris", 33},
			{"Cleo", 10},
		}

		AssertLeague(got, want, t)

		// read again
		got = store.GetLeague()
		AssertLeague(got, want, t)
	})
}

func TestTape_Write(t *testing.T) {
	file, clean := CreateTempFile(t, "12345")
	defer clean()
	tape := &tape{file}
	tape.Write([]byte("abc"))
	file.Seek(0, 0)
	newFileContents, _ := ioutil.ReadAll(file)
	got := string(newFileContents)
	want := "abc"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, "/league", nil)
	return req
}

func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got status %d want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}

}

const jsonContentType = "application/json"

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of %s, got %v", want, response.HeaderMap)
	}
}
