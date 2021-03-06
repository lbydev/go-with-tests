package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type StubPlayerStore struct {
	scores   map[string]int
	winCalls []string
	league   League
}

func (s *StubPlayerStore) GetPlayerScore(name string) int {
	score := s.scores[name]
	return score
}

func (s *StubPlayerStore) RecordWin(name string) {
	s.winCalls = append(s.winCalls, name)
}

func (s *StubPlayerStore) GetLeague() League {
	return s.league
}

// func TestGETPlayers(t *testing.T) {
// 	store := StubPlayerStore{
// 		map[string]int{
// 			"Pepper": 20,
// 			"Floyd":  10,
// 		},
// 		nil,
// 	}
// 	server := NewPlayerServer(&store)
// 	tests := []struct {
// 		name               string
// 		player             string
// 		expectedHTTPStatus int
// 		expectedScore      string
// 	}{
// 		{
// 			name:               "Returns Pepper's score",
// 			player:             "Pepper",
// 			expectedHTTPStatus: http.StatusOK,
// 			expectedScore:      "20",
// 		},
// 		{
// 			name:               "Returns Floyd's score",
// 			player:             "Floyd",
// 			expectedHTTPStatus: http.StatusOK,
// 			expectedScore:      "10",
// 		},
// 		{
// 			name:               "Returns 404 on missing players",
// 			player:             "Apollo",
// 			expectedHTTPStatus: http.StatusNotFound,
// 			expectedScore:      "0",
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			request := newGetScoreRequest(tt.player)
// 			response := httptest.NewRecorder()
// 			server.ServeHTTP(response, request)
// 			assertStatus(t, response.Code, tt.expectedHTTPStatus)
// 			assertResponseBody(t, response.Body.String(), tt.expectedScore)
// 		})
// 	}
// }
//
// func TestStoreWins(t *testing.T) {
// 	store := StubPlayerStore{
// 		map[string]int{},
// 		nil,
// 	}
//
// 	server := NewPlayerServer(&store)
//
// 	t.Run("it records wins on POST", func(t *testing.T) {
//
// 		player := "Peper"
//
// 		request := newPostWinRequest(player)
// 		response := httptest.NewRecorder()
//
// 		server.ServeHTTP(response, request)
//
// 		assertStatus(t, response.Code, http.StatusAccepted)
//
// 		if len(store.winCalls) != 1 {
// 			t.Fatalf("got %d calls to RecordWin want %d", len(store.winCalls), 1)
// 		}
// 		if store.winCalls[0] != player {
// 			t.Errorf("did not store correct winner got %q want %q", store.winCalls[0], player)
// 		}
// 	})
// }
func newGetScoreRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", name), nil)
	return req
}
func newPostWinRequest(name string) *http.Request {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return req
}

func newLeagueRequest() *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/league"), nil)
	return req
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}

func assertResponseBody(t *testing.T, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("response body is wrong, got '%s' want '%s'", got, want)
	}
}

func TestLegue(t *testing.T) {
	store := StubPlayerStore{}
	server := NewPlayerServer(&store)

	t.Run("it returns 200 on /league", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player

		err := json.NewDecoder(response.Body).Decode(&got)
		if err != nil {
			t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", response.Body, err)
		}
		assertStatus(t, response.Code, http.StatusOK)
	})

	t.Run("it returns the league table as JSON", func(t *testing.T) {

		wantedLeague := []Player{
			{"Cleo", 32},
			{"Chris", 20},
			{"Tiest", 14},
		}
		store := &StubPlayerStore{nil, nil, wantedLeague}
		server := NewPlayerServer(store)

		request, _ := http.NewRequest(http.MethodGet, "/league", nil)
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		var got []Player
		got = getLeagueFromResponse(t, response.Body)
		assertStatus(t, response.Code, http.StatusOK)
		assertContentType(t, response, jsonContentType)
		assertLeague(t, got, wantedLeague)
	})

}

func getLeagueFromResponse(t *testing.T, body io.Reader) (league []Player) {
	t.Helper()
	err := json.NewDecoder(body).Decode(&league)
	if err != nil {
		t.Fatalf("Unable to parse response from server '%s' into slice of Player, '%v'", body, err)
	}
	return
}

func assertLeague(t *testing.T, got, want []Player) {

	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func assertContentType(t *testing.T, response *httptest.ResponseRecorder, want string) {
	t.Helper()
	if response.Header().Get("content-type") != want {
		t.Errorf("response did not have content-type of application/json, got %v", response.Header())
	}
}
