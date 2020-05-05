package config

import "flag"

var (
	HTTP_PORT        *string
	COUCHDB_ENDPOINT *string
	COUCHDB_USERNAME *string
	COUCHDB_PASSWORD *string
)

func init() {
	HTTP_PORT = flag.String("port", "8080", "http server port.")
	COUCHDB_ENDPOINT = flag.String("couchdb.endpoint", "http://127.0.0.1:5984", "couch-db server endpoint.")
	COUCHDB_USERNAME = flag.String("couchdb.username", "", "couch-db server username.")
	COUCHDB_PASSWORD = flag.String("couchdb.password", "", "couch-db server password.")
	flag.Parse()
}
