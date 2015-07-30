package server_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/david1983xtc/colernio-backend/server"
)

var (
	srv *httptest.Server
)

func init() {
	srv = httptest.NewServer(server.Router)
}

func TestGETIndexRequest(t *testing.T) {
	reader := strings.NewReader("")

	request, err := http.NewRequest("GET", fmt.Sprintf("%s/", srv.URL), reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("Expected: %d, Received: %d", http.StatusOK, res.StatusCode)
	}
}

func TestPOSTIndexRequest(t *testing.T) {
	reader := strings.NewReader("")

	request, err := http.NewRequest("POST", fmt.Sprintf("%s/", srv.URL), reader)

	res, err := http.DefaultClient.Do(request)

	if err != nil {
		t.Error(err)
	}

	if res.StatusCode != http.StatusNotFound {
		t.Errorf("Expected: %d, Received: %d", http.StatusNotFound, res.StatusCode)
	}
}

// go test -bench .
func BenchmarkGETIndex(b *testing.B) {
	for n := 0; n < b.N; n++ {

	}
}
