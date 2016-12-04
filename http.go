package main // import "github.com/aphexddb/lightshow"

import (
	"encoding/json"
	"net/http"

	log "github.com/Sirupsen/logrus"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// InterfaceToPrettyJSON serializes an interface into pretty formatted JSON
func InterfaceToPrettyJSON(i interface{}) ([]byte, error) {
	json, marshalErr := json.Marshal(i)
	if marshalErr != nil {
		return nil, marshalErr
	}
	sJSON, jsonErr := simplejson.NewJson(json)
	if jsonErr != nil {
		return nil, jsonErr
	}
	respJSON, prettyErr := sJSON.EncodePretty()
	if prettyErr != nil {
		return nil, prettyErr
	}
	return respJSON, nil
}

// ServeHTTP starts serving HTTP
func ServeHTTP() {
	r := mux.NewRouter().StrictSlash(false)

	// api endpoints
	r.HandleFunc("/foo", FooHandler).Methods("GET")

	// Serve HTML as catchall
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./html/")))

	// Log HTTP calls through logrus
	w := log.StandardLogger().Writer()
	defer w.Close()
	loggedRouter := handlers.CombinedLoggingHandler(w, r)

	// Serve HTTP
	httpServer := ":8000"
	log.Infof("HTTP server starting %s", httpServer)
	log.Fatal(http.ListenAndServe(httpServer, loggedRouter))
}
