package tests

import (
	"encoding/json"
	"frontend-data/internal/server"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestNotFound(t *testing.T) {
	s := &server.Server{}
	server := httptest.NewServer(s.RegisterRoutes())
	defer server.Close()
	resp, err := http.Get(server.URL + "/nonexistant")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()

	// Assertions
	if resp.StatusCode != http.StatusNotFound {
		t.Errorf("expected status `%v`; got `%v`", http.StatusNotFound, resp.Status)
	}
	expected := http.StatusText(http.StatusNotFound)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("error reading response body. Err: %v", err)
	}

	before, _ := strings.CutSuffix(string(body), "\n")
	if expected != before {
		t.Errorf("expected response body to be `%v`; got `%v`", expected, string(body))
	}
}

func TestFound(t *testing.T) {
	s := &server.Server{}
	server := httptest.NewServer(s.RegisterRoutes())
	defer server.Close()
	resp, err := http.Get(server.URL + "/index")
	if err != nil {
		t.Fatalf("error making request to server. Err: %v", err)
	}
	defer resp.Body.Close()

	// Assertions
	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected status `%v`; got `%v`", http.StatusOK, resp.Status)
	}

	var data struct {
		FileName string   `json:"fileName"`
		Title    string   `json:"title"`
		Scripts  []string `json:"scripts"`
		Header   struct {
			Title   string `json:"title"`
			Type    string `json:"type"`
			Content string `json:"content"`
		} `json:"header"`
		Body []struct {
			Title   string `json:"title"`
			Content []struct {
				Type    string `json:"type"`
				Title   string `json:"title"`
				Content string `json:"content"`
			} `json:"content"`
		} `json:"body"`
		Footer struct {
			Type    string `json:"type"`
			Title   string `json:"title"`
			Content string `json:"content"`
			Link    string `json:"link"`
		} `json:"footer"`
		SidePanel []struct {
			Title string `json:"title"`
			Link  string `json:"link"`
			Image string `json:"image"`
		} `json:"sidePanel"`
	}

	body, _ := io.ReadAll(resp.Body)

	err = json.Unmarshal(body, &data)
	if err != nil {
		t.Errorf("Could not unmarshal data, got: `%s`", body)
		return
	}
}
