# go-starter-service
This is a sample go app for go 1.12+. This example simply takes a request at the home URL and writes some sample data to Google Cloud Datastore (of a pre-configured GCP project).
This uses go modules for dependency to avoid all the confusion around GOPATH and relative directory of external packages.


## Test
```
go test
```

## Run
```
go run main.go
```

## Invoke Endpoint
```
curl localhost:8080?key=test
```

## References 
* https://golang.org/pkg/testing/
* https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
* https://blog.golang.org/json-and-go
* https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/
* https://golang.org/pkg/net/http/httptest/#NewRequest
*
