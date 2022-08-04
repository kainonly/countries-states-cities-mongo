package main

import (
	"github.com/kainonly/countries-states-cities-mongo/bootstrap"
	"net/http"
)

func main() {
	api, err := bootstrap.NewAPI()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/event-invoke", api.EventInvoke)
	http.ListenAndServe(api.Values.Address, nil)
}
