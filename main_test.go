package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func checkError(err error, t *testing.T) {
	if err != nil {
		t.Errorf("An error occurred. %v", err)
	}
}

func TestIndexHandler(t *testing.T) {

	req, err := http.NewRequest("GET", "/?key=foo", nil)

	checkError(err, t)

	rr := httptest.NewRecorder()

	//Make the handler function satisfy http.Handler
	//https://lanreadelowo.com/blog/2017/04/03/http-in-go/
	http.HandlerFunc(indexHandler).
		ServeHTTP(rr, req)

	//Confirm the response has the right status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
	}

	//Confirm the returned json is what we expected
	//Manually build up the expected json string
	expected := string(`Neil Ghosh`)

	//The assert package checks if both JSON string are equal and for a plus, it actually confirms if our manually built JSON string is valid
	var result map[string]interface{}
	json.Unmarshal([]byte(rr.Body.String()), &result)
	actual := result["Name"]

	assert.Equal(t, expected, actual, "Response body differs")
}
