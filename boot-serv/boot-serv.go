package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mahuaibo/quorum-boot/boot-serv/action/consortium"
	"github.com/mahuaibo/quorum-boot/boot-serv/action/node"
	"github.com/mahuaibo/quorum-boot/boot-serv/config"
	"net/http"
)

var (
	router *mux.Router
)

func init() {
	router = mux.NewRouter()
	funcs := []struct {
		Name   string
		Body   func(w http.ResponseWriter, r *http.Request)
		Method string
	}{
		// consortium
		{
			Name:   "/consortiums/{id}/boot",
			Body:   consortium.Boot,
			Method: "GET",
		},
		{
			Name:   "/consortiums",
			Body:   consortium.List,
			Method: "GET",
		}, {
			Name:   "/consortiums",
			Body:   consortium.New,
			Method: "POST",
		}, {
			Name:   "/consortiums/{id}",
			Body:   consortium.Delete,
			Method: "DELETE",
		}, {
			Name:   "/consortiums/{id}",
			Body:   consortium.Modify,
			Method: "PUT",
		}, {
			Name:   "/consortiums/{id}",
			Body:   consortium.Get,
			Method: "GET",
		}, {
			Name:   "/consortiums/{id}/exist",
			Body:   consortium.Exist,
			Method: "GET",
		},
		// node
		{
			Name:   "/consortiums/{id}/nodes",
			Body:   node.List,
			Method: "GET",
		}, {
			Name:   "/consortiums/{id}/nodes",
			Body:   node.New,
			Method: "POST",
		}, {
			Name:   "/consortiums/{id}/nodes/{publicKey}",
			Body:   node.Delete,
			Method: "DELETE",
		}, {
			Name:   "/consortiums/{id}/nodes/{publicKey}",
			Body:   node.Modify,
			Method: "PUT",
		}, {
			Name:   "/consortiums/{id}/nodes/{publicKey}",
			Body:   node.Get,
			Method: "GET",
		}, {
			Name:   "/consortiums/{id}/nodes/{publicKey}/exist",
			Body:   node.Exist,
			Method: "GET",
		},
	}

	for _, fun := range funcs {
		router.HandleFunc(fun.Name, fun.Body).Methods(fun.Method)
	}
}

func main() {
	println("Running the boot-service...")
	err := http.ListenAndServe(":"+(*config.HTTP_PORT), router)
	if err != nil {
		fmt.Errorf("Error: %s\n", err)
	}
}
