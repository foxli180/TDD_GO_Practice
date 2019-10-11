package _select

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T)  {
	slowUrl := "http://www.whe.edu.cn"
	fastUrl := "http://www.baidu.com"
	want := fastUrl
	got, _ := Racer(slowUrl, fastUrl)
	if got != want {
		t.Errorf("got '%s', want '%s'", got, want)
	}
}

func TestRacer2(t *testing.T) {

	t.Run("get fatest url", func(t *testing.T) {
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)

		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL

		want := fastURL
		got, _ := Racer(fastURL, slowURL)
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	})
	
	t.Run("return an error if no response in 10 seconds", func(t *testing.T) {
		serverA := makeDelayedServer(11 * time.Second)
		serverB := makeDelayedServer(12 * time.Second)

		defer serverA.Close()
		defer serverB.Close()

		_, err := Racer(serverA.URL, serverB.URL)

		if err == nil {
			t.Error("expected an error but did not get one")
		}
	})

	t.Run("return an error if no response in timedout", func(t *testing.T) {
		slowServer := makeDelayedServer(25 * time.Millisecond)

		defer slowServer.Close()

		_, err := ConfigurableRacer(slowServer.URL, slowServer.URL, 20 * time.Millisecond)

		if err == nil {
			t.Error("expected an error but did not get one")
		}
	})
}

func makeDelayedServer(duration time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			time.Sleep(duration)
			w.WriteHeader(http.StatusOK)
		}))
}