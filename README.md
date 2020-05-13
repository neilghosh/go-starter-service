# go-starter-service
This is a sample go app for go 1.12+. This example simply takes a request at the home URL and writes some sample data to Google Cloud Datastore (of a pre-configured GCP project).
This uses go modules for dependency to avoid all the confusion around GOPATH and relative directory of external packages.

## Install Go and Google Cloud SDK 
Note that if the min version is not met , go module features won't work.
https://golang.org/doc/install
https://cloud.google.com/sdk/docs#deb

```
go version 

gcloud init //to setup a project 
gcloud auth application-default login //to get the default credential to talk to cloud services locally 
```

## build and run test
This should download all the dependencies and install in a directory 

```
go test
```

## Run
```
go run main.go
```

## Invoke Endpoint
Check the portn in which the server is running
```
POST localhost:8080/api
{Name:test Content:someContent EventDate:2019-11-05 13:15:30 +0000 UTC}

GET localhost:8080/api?id=004161a0-b42b-4ccb-ac2c-dd7c6a81e049

PUT localhost:8080/api?id=004161a0-b42b-4ccb-ac2c-dd7c6a81e049
{Name:test Content:someContent EventDate:2019-11-05 13:15:30 +0000 UTC}

DELETE localhost:8080/api?id=004161a0-b42b-4ccb-ac2c-dd7c6a81e049
```

## Debug 
Install ```dlv```

```
go get -u github.com/go-delve/delve/cmd/dlv
```
Restart VSCode

Then you can debug from VSCode with the following run debug configuration 

```
{
        "name": "Launch file",
        "type": "go",
        "request": "launch",
        "mode": "auto",
        "program": "${file}"
}
   
```

## References 
* https://golang.org/pkg/testing/
* https://medium.com/rungo/unit-testing-made-easy-in-go-25077669318
* https://blog.golang.org/json-and-go
* https://www.sohamkamani.com/blog/2017/10/18/parsing-json-in-golang/
* https://golang.org/pkg/net/http/httptest/#NewRequest


##Change
