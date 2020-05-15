package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
    // "io/ioutil"

	"cloud.google.com/go/datastore"
	"github.com/google/uuid"
)

const PROJECT_ID = "demoneil"
const FEED_ENTITY = "FeedItem"

type FeedItem struct {
	FeedName string
	FeedItemId string
	Active   bool
	Content  string
	Created  time.Time
	EventDate time.Time
}

type FeedItemRequest struct {
	Name string 
	Content string `json:"content"`
	EventDate time.Time
}

type FeedItemResponse struct {
	Id string 
}

func echoHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprint(w, "Hello Master")
}

func restHandler(w http.ResponseWriter, r *http.Request) {

	//TODO Do we really needs this if handler routing works
	// if r.URL.Path != "/api" {
	// 	http.NotFound(w, r)
	// 	return
	// }

	w.Header().Set("Content-Type", "application/json")
    switch r.Method {

    case "GET":
		w.WriteHeader(http.StatusOK)
		ids, ok := r.URL.Query()["id"]
	    var id string

		if !ok || len(ids[0]) < 1 {
			log.Println("Url Param 'id' is missing")			
		} else {
			id = ids[0]
		}

		var response = getFeed(id)
		json.NewEncoder(w).Encode(response)
        //w.Write([]byte(`{"message": "get called"}`))
	case "POST":
		w.WriteHeader(http.StatusCreated)
		var request FeedItemRequest

		// This is another way to get the body 
		// body, er := ioutil.ReadAll(r.Body)
		// if er != nil {
		// 	panic(er)
		// }
		// log.Println(string(body))
		// err = json.Unmarshal(body, &request)

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		var response = postFeed(request)
		json.NewEncoder(w).Encode(response)
    case "PUT":

		ids, ok := r.URL.Query()["id"]
	    var id string

		if !ok || len(ids[0]) < 1 {
			log.Println("Url Param 'id' is missing")			
		} else {
			id = ids[0]
		}


        w.WriteHeader(http.StatusAccepted)
		var request FeedItemRequest

		err := json.NewDecoder(r.Body).Decode(&request)

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	
		var response = updateFeed(id, request)
		json.NewEncoder(w).Encode(response)
	case "DELETE":
		
		ids, ok := r.URL.Query()["id"]
	    var id string

		if !ok || len(ids[0]) < 1 {
			log.Println("Url Param 'id' is missing")			
		} else {
			id = ids[0]
		}
		deleteFeed(id)
        w.WriteHeader(http.StatusNoContent)
        w.Write([]byte(`{"message": "delete called"}`))
    default:
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"message": "not found"}`))
	}
	
}

	func getDataStoreClient(ctx context.Context) *datastore.Client {
		// Creates a datastore client.
		datastoreClient, err := datastore.NewClient(ctx, PROJECT_ID)
		if err != nil {
			log.Fatal(err)	
		}
		return datastoreClient

	}

	func writeToDatabase(key string,feedItem *FeedItem){
		ctx := context.Background()
		datastoreClient := getDataStoreClient(ctx)
		dskey := datastore.NameKey(FEED_ENTITY, key, nil)
		dskey, _ = datastoreClient.Put(ctx, dskey, feedItem)
		log.Printf(fmt.Sprintf("Written entry to database %v", dskey))
	}

	func deleteFeed(id string) {
		ctx := context.Background()
		datastoreClient :=  getDataStoreClient(ctx);

		dskey := datastore.NameKey("FeedItem", id, nil)
		if err := datastoreClient.Delete(ctx, dskey); err != nil {
			log.Printf(fmt.Sprintf("deleted entry to database %v", dskey))
		}
	}

	func getFeed(id string) FeedItem {
		ctx := context.Background()
		datastoreClient :=  getDataStoreClient(ctx);

		dskey := datastore.NameKey("FeedItem", id, nil)
		feedItem := FeedItem{}
		if err := datastoreClient.Get(ctx, dskey, &feedItem); err != nil {
			log.Printf(fmt.Sprintf("Written entry to database %v", dskey))
		}
	
		log.Printf(fmt.Sprintf("Written entry to database %v", dskey))
		return feedItem
	} 

	func  postFeed(request FeedItemRequest) FeedItemResponse {
		log.Printf("Request: %+v", request)

		feedItem := &FeedItem{
			FeedName: request.Name,
			Active:  true,
			Content : request.Content, 
			Created: time.Now(),
			EventDate : request.EventDate,
		}
	
		key := uuid.New().String()		
		writeToDatabase(key, feedItem)
	
		//taskStr, _ := json.Marshal(feedItem)

		response := FeedItemResponse{
			Id:     key,
		}
	
		data, _ := json.Marshal(response)
		log.Println("Respionse : "+string(data))
		return response
	}

	func  updateFeed(id string, request FeedItemRequest) FeedItemResponse {
		log.Printf("Request: %+v", request)

		feedItem := &FeedItem{
			FeedName: request.Name,
			Active:  true,
			Content : request.Content, 
			Created: time.Now(),
			EventDate : request.EventDate,
		}
	
		key := id		
		writeToDatabase(key, feedItem)
	
		//taskStr, _ := json.Marshal(feedItem)

		response := FeedItemResponse{
			Id:     key,
		}
	
		data, _ := json.Marshal(response)
		log.Println("Respionse : "+string(data))
		return response
	}

func main() {
	http.HandleFunc("/api/", restHandler)
	http.HandleFunc("/", echoHandler)

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
