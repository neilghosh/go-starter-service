		package main

		import (
			"fmt"
			"bytes"
			"encoding/json"
			"net/http"
			"net/http/httptest"
			"time"
			"testing"

			"github.com/stretchr/testify/assert"
		)

		func checkError(err error, t *testing.T) {
			if err != nil {
				t.Errorf("An error occurred. %v", err)
			}
		}


		func TestRestHandlerForUpdate(t *testing.T) {

		}

		func TestRestHandlerForPost(t *testing.T) {

			testTime := "2019-11-05T13:15:30Z"
			testContent := "someContent"
			testName := "test"


			var jsonData = []byte(`{"name":"`+testName+`","content": "`+testContent+`", "eventDate": "`+testTime+`"}`)
			req, err := http.NewRequest("POST", "/api", bytes.NewBuffer(jsonData))

			checkError(err, t)

			rr := httptest.NewRecorder()

			//Make the handler function satisfy http.Handler
			//https://lanreadelowo.com/blog/2017/04/03/http-in-go/
			http.HandlerFunc(restHandler).
				ServeHTTP(rr, req)

			// //Confirm the response has the right status code
			if status := rr.Code; status != http.StatusCreated {
				t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusCreated, status)
			}

			//The assert package checks if both JSON string are equal and for a plus, it actually confirms if our manually built JSON string is valid
			var result map[string]interface{}
			json.Unmarshal([]byte(rr.Body.String()), &result)
			actual := result["Id"]


			assert.NotEmpty(t, actual, "Response id differs")

			//GET Stuff
			url :=  "/api?id="+ actual.(string)
			req1, err1 := http.NewRequest("GET", url, nil)


			checkError(err1, t)

			rr1 := httptest.NewRecorder()

			//Make the handler function satisfy http.Handler
			//https://lanreadelowo.com/blog/2017/04/03/http-in-go/
			http.HandlerFunc(restHandler).
				ServeHTTP(rr1, req1)

			// //Confirm the response has the right status code
			if status := rr1.Code; status != http.StatusOK {
				t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusOK, status)
			}

			//Confirm the returned json is what we expected
			//Manually build up the expected json string
			expected := string(testContent)
			expectedEventDate, _ := time.Parse(time.RFC3339,testTime)

			//The assert package checks if both JSON string are equal and for a plus, it actually confirms if our manually built JSON string is valid
			var result1 map[string]interface{}
			json.Unmarshal([]byte(rr1.Body.String()), &result1)
			actualContent := result1["Content"]
			actualEventDate, _ := time.Parse(time.RFC3339, fmt.Sprintf("%v", result1["EventDate"]))

			assert.Equal(t, expected, actualContent, "Response content differs")
			assert.True(t, expectedEventDate.UnixNano() == actualEventDate.UnixNano(), "Response event date differs")


			//update stuff


			timeToUpdate := "2020-11-05T13:15:30Z"
			contentToUpdate := "someMoreContent"
			nameToUpdate := "testput"

			var updateData = []byte(`{"name":"`+nameToUpdate+`","content": "`+contentToUpdate+`", "eventDate": "`+timeToUpdate+`"}`)


			putUrl :=  "/api?id="+ actual.(string)
			putRequest, err := http.NewRequest("PUT", putUrl,  bytes.NewBuffer(updateData))

			checkError(err, t)

			putRecorder := httptest.NewRecorder()


			http.HandlerFunc(restHandler).
			ServeHTTP(putRecorder, putRequest)

			// //Confirm the response has the right status code
			if status := putRecorder.Code; status != http.StatusAccepted {
				t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusCreated, status)
			}


			deleteUrl :=  "/api?id="+ actual.(string)
			deleteRequest, err := http.NewRequest("DELETE", deleteUrl,nil)

			checkError(err, t)

			deleteRecorder := httptest.NewRecorder()


			http.HandlerFunc(restHandler).
			ServeHTTP(deleteRecorder, deleteRequest)

			// //Confirm the response has the right status code
			if status := deleteRecorder.Code; status != http.StatusNoContent {
				t.Errorf("Status code differs. Expected %d .\n Got %d instead", http.StatusCreated, status)
			}
		}


