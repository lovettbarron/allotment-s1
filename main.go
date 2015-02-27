package main

import (
	"fmt"
	"net/http"
	_ "os"
	"time"
	"sync"
	"log"
	"bytes"
	"github.com/gorilla/mux"
	"github.com/mitchellh/goamz/aws"
	"github.com/mitchellh/goamz/s3"
)

const (
	Port string = ":8000"
	webcamUrl string = "http://webmarin.com/images/wc/Camera.jpg"
	UpdateCycle int = 60 // in seconds

	Key string = "123"
	Secret string = "123"
)

type Image struct {
	name string // date
	path string // s3
	data []byte // for image comparison
	mutex *sync.Mutex	
}

var r = mux.NewRouter()

func main() {
	// auth, err := aws.EnvAuth()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// client := s3.New(auth, aws.USEast)

	r.HandleFunc("/",GetIndex)
	r.HandleFunc("/{date}", GetDates)

	http.Handle("/", r)

	fmt.Println("Server running at port",Port)
	http.ListenAndServe(Port, nil)
}

func CheckAtInterval() <-chan bool {
	ticker := time.NewTicker(time.Duration(UpdateCycle) * time.Second)
	quit := make(chan bool)

	go func() {
		for {
			select {
				case <- ticker.C:
				    fetchImage();
				case <- quit:
				    ticker.Stop()
			}
		}
	}()

	return quit
}

func fetchImage() *Image {

	time := time.Now().Unix()
	// filename,err := os.Create(string(time))

	check := http.Client{
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
		    r.URL.Opaque = r.URL.Path
		    return nil
		},
	}
	resp, err := check.Get(webcamUrl) // add a filter to check redirect
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	buf.ReadFrom(resp.Body)


	img := &Image{
		string(time),
		"",
		buf.Bytes(),
		new(sync.Mutex),
	}

	return img
}

func writeImage() {
	// resp, err := client.ListBuckets()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// log.Print(fmt.Sprintf("%T %+v", resp.Buckets[0], resp.Buckets[0]))
}

func getBucket()  []s3.Key {
    auth := aws.Auth{
        AccessKey: Key,
        SecretKey: Secret,
    }
    euwest := aws.EUWest

    connection := s3.New(auth, euwest)
    mybucket := connection.Bucket("allotment")
    res, err := mybucket.List("", "", "", 1000)
    if err != nil {
        log.Fatal(err)
    }
    for _, v := range res.Contents {
        fmt.Println(v.Key)
    }
    return res.Contents
}

func getDay() {

}
