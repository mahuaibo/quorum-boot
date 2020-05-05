package node

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mahuaibo/quorum-boot/boot-serv/common"
	"github.com/mahuaibo/quorum-boot/boot-serv/database"
	"net/http"
	"strings"
)

func List(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	client, err := database.CouchClient()
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
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
	data := []string{}
	for _, row := range result.Rows {
		if row.Id != common.DOC_GENESIS {
			data = append(data, strings.ReplaceAll(row.Id, common.DOC_PREFIX, ""))
		}
	}
	common.Response(&w, common.Normal, "Success.", data)
	return
}

func New(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	decoder := json.NewDecoder(r.Body)
	param := struct {
		PublicKey string `json:"publicKey"`
		Host      string `json:"host"`
		Port      uint64 `json:"port"`
		RaftPort  uint64 `json:"raftport"`
	}{}
	decoder.Decode(&param)
	param.PublicKey = common.HandlePublicKey(param.PublicKey)

	client, err := database.CouchClient()
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	has, _, err := database.HasNode(client, id, param.PublicKey)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	if has {
		common.Response(&w, common.Error, "node "+param.PublicKey+" already existed.", nil)
		return
	}
	_, err = database.AddNode(client, id, param.PublicKey, param)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	common.Response(&w, common.Normal, "Success.", param)
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	publicKey := mux.Vars(r)["publicKey"]
	publicKey = common.HandlePublicKey(publicKey)
	client, err := database.CouchClient()
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	has, rev, err := database.HasNode(client, id, publicKey)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	if has {
		_, err = database.DeleteNode(client, id, publicKey, rev)
		if err != nil {
			common.Response(&w, common.Error, err.Error(), nil)
			return
		}
	}
	data := struct {
		Id        string `json:"id"`
		PublicKey string `json:"publicKey"`
	}{
		Id:        id,
		PublicKey: publicKey,
	}
	common.Response(&w, common.Normal, "Success.", data)
	return
}

func Modify(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	publicKey := mux.Vars(r)["publicKey"]
	publicKey = common.HandlePublicKey(publicKey)
	decoder := json.NewDecoder(r.Body)
	param := struct {
		PublicKey string `json:"publicKey"`
		Host      string `json:"host"`
		Port      uint64 `json:"port"`
		RaftPort  uint64 `json:"raftport"`
	}{}
	decoder.Decode(&param)
	param.PublicKey = publicKey

	client, err := database.CouchClient()
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}

	has, rev, err := database.HasNode(client, id, param.PublicKey)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	if !has {
		common.Response(&w, common.Error, "node "+param.PublicKey+" not exist", nil)
	}
	_, err = database.ModifyNode(client, id, param.PublicKey, rev, param)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	common.Response(&w, common.Normal, "Success.", param)
	return
}

func Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	publicKey := mux.Vars(r)["publicKey"]
	publicKey = common.HandlePublicKey(publicKey)

	client, err := database.CouchClient()
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}

	has, err := database.HasConsortium(client, id)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	if !has {
		common.Response(&w, common.Error, "consortium "+id+" not exist", nil)
		return
	}
	one := struct {
		PublicKey string `json:"publicKey"`
		Host      string `json:"host"`
		Port      uint64 `json:"port"`
		RaftPort  uint64 `json:"raftport"`
	}{}

	err = database.GetNode(client, id, publicKey, &one)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	common.Response(&w, common.Normal, "Success.", one)
	return
}

func Exist(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	publicKey := mux.Vars(r)["publicKey"]
	publicKey = common.HandlePublicKey(publicKey)

	client, err := database.CouchClient()
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	has, _, err := database.HasNode(client, id, publicKey)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	data := struct {
		PublicKey string `json:"publicKey"`
		Exist     bool   `json:"exist"`
	}{
		PublicKey: publicKey,
		Exist:     has,
	}
	common.Response(&w, common.Normal, "Success.", data)
	return
}
