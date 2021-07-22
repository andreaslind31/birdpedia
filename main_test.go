//Go uses a convention to ascertains a test file when it has the pattern *_test.go
package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandler(t *testing.T) {
	// This is the request that's going to be passed to our handler.
	req, err := http.NewRequest("GET", "", nil)

	if err != nil {
		t.Fatal(err)
	}

	recorder := httptest.NewRecorder()

	// Create an HTTP handler from our handler function defined in our main.go file that we want to test
	hf := http.HandlerFunc(handler)

	// Serve the HTTP request to our recorder. This is the line that actually
	// executes our the handler that we want to test
	hf.ServeHTTP(recorder, req)

	if status := recorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	expected := `Hello world!`
	actual := recorder.Body.String()
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}

}

func TestRouter(t *testing.T) {

	r := newRouter()

	// Documentation : https://golang.org/pkg/net/http/httptest/#NewServer
	mockServer := httptest.NewServer(r)

	resp, err := http.Get(mockServer.URL + "/hello")

	if err != nil {
		t.Fatal(err)
	}

	// We want our status to be 200 (ok)
	if resp.StatusCode != http.StatusOK {
		t.Errorf("Status should be ok, got %d", resp.StatusCode)
	}

	// In the next few lines, the response body is read, and converted to a string
	defer resp.Body.Close()
	// read the body into a bunch of bytes (b)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatal(err)
	}
	// convert the bytes to a string
	respString := string(b)
	expected := "Hello world!"

	if respString != expected {
		t.Errorf("Response should be %s, got %s", expected, respString)
	}

}
