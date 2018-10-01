package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"os"
)

func main() {

  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "The Go server index page")
  })

  http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
      response, err := http.Get("http://api.airvisual.com/v2/nearest_city?key=qocJKMPPufJsCaEAd")

        if err != nil {
           fmt.Print(err.Error())
           os.Exit(1)
        }

        responseData, err := ioutil.ReadAll(response.Body)
        if err != nil {
                log.Fatal(err)
        }
	fmt.Fprintf(w, "The output of the curl request is:  %s\n", responseData)

	})

	http.ListenAndServe(":5000", nil)
}
