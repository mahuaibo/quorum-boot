package database

import (
	"errors"
	"github.com/fjl/go-couchdb"
	"github.com/mahuaibo/quorum-boot/boot-serv/common"
	"github.com/mahuaibo/quorum-boot/boot-serv/config"
)

func CouchClient() (*couchdb.Client, error) {
	client, err := couchdb.NewClient(*config.COUCHDB_ENDPOINT, nil)
	client.SetAuth(couchdb.BasicAuth(*config.COUCHDB_USERNAME, *config.COUCHDB_PASSWORD))
	return client, err
}

func HasNode(client *couchdb.Client, id, publicKey string) (bool, string, error) {
	has, err := HasConsortium(client, id)
	if err != nil {
		return false, "", err
	}
	if !has {
		return false, "", errors.New("consortium " + id + " not exist")
	}
	result := struct {
		Rows []struct {
			Id    string `json:"id"`
			Value struct {
				Rev string `json:"rev"`
			} `json:"value"`
		} `json:"rows"`
	}{}
	client.DB(common.GetConsortiumName(id)).AllDocs(&result, nil)
	for _, row := range result.Rows {
		if row.Id == common.GetNodeName(publicKey) {
			return true, row.Value.Rev, nil
		}
	}
	return false, "", nil
}

func AddNode(client *couchdb.Client, id, publicKey string, data interface{}) (string, error) {
	rev, err := client.DB(common.GetConsortiumName(id)).Put(common.GetNodeName(publicKey), data, "")
	return rev, err
}

func ModifyNode(client *couchdb.Client, id, publicKey, rev string, data interface{}) (string, error) {
	rev, err := client.DB(common.GetConsortiumName(id)).Put(common.GetNodeName(publicKey), data, rev)
	return rev, err
}

func DeleteNode(client *couchdb.Client, id, publicKey, rev string) (string, error) {
	_rev, err := client.DB(common.GetConsortiumName(id)).Delete(common.GetNodeName(publicKey), rev)
	return _rev, err
}

func GetNode(client *couchdb.Client, id, publicKey string, data interface{}) error {
	err := client.DB(common.GetConsortiumName(id)).Get(common.GetNodeName(publicKey), data, nil)
	return err
}

func HasConsortium(client *couchdb.Client, id string) (bool, error) {
	names, err := client.AllDBs()
	if err != nil {
		return false, err
	}
	for _, _name := range names {
		if common.GetConsortiumName(id) == _name {
			return true, nil
		}
	}
	return false, nil
}

func ModifyConsortium(client *couchdb.Client, id, rev string, data interface{}) (string, error) {
	rev, err := client.DB(common.GetConsortiumName(id)).Put(common.DOC_GENESIS, data, rev)
	return rev, err
}

func DeleteConsortium(client *couchdb.Client, id string) error {
	err := client.DeleteDB(common.GetConsortiumName(id))
	return err
}

func AddConsortium(client *couchdb.Client, id string, data interface{}) (string, error) {
	db, err := client.EnsureDB(common.GetConsortiumName(id))
	if err != nil {
		return "", err
	}
	rev, err := db.Put(common.DOC_GENESIS, data, "")
	if err != nil {
		return "", err
	}
	return rev, err
}

func GetConsortium(client *couchdb.Client, id string, data interface{}) error {
	err := client.DB(common.GetConsortiumName(id)).Get(common.DOC_GENESIS, data, nil)
	return err
}
