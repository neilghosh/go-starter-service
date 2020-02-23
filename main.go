package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
)

type Employee struct {
	Name    string
	Active  bool
	Age     int
	Salary  float64
	Created time.Time
}

// indexHandler responds to requests.
func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	ctx := context.Background()

	// Sets your Google Cloud Platform project ID.
	//TODO not sure if it can be defaulted while running within appengine
	projectID := "demoneil"

	// Creates a datastore client.
	datastoreClient, _ := datastore.NewClient(ctx, projectID)

	//NamesKey for setting the key yourself.
	keys, ok := r.URL.Query()["key"]
	var key string

	task := &Employee{
		Name:    "Neil Ghosh",
		Active:  true,
		Age:     99,
		Salary:  10.02,
		Created: time.Now(),
	}

	if !ok || len(keys[0]) < 1 {
		log.Println("Url Param 'key' is missing")
		//Generate Random ID otherwise
		key = uuid.New().String()
	} else {
		key = keys[0]
	}
	dskey := datastore.NameKey("__namespace__", key, nil)
	dskey, _ = datastoreClient.Put(ctx, dskey, task)

	log.Printf(fmt.Sprintf("Written entry to database %v", dskey))
	taskStr, _ := json.Marshal(task)
	fmt.Fprint(w, string(taskStr))
}

func main() {
	http.HandleFunc("/", indexHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("Defaulting to port %s", port)
	}

	log.Printf("Listening on port %s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
