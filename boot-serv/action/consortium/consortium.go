package consortium

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mahuaibo/quorum-boot/boot-serv/common"
	"github.com/mahuaibo/quorum-boot/boot-serv/database"
	"net/http"
	"strings"
)

type Alloc map[string]struct {
	Balance string `json:"balance"`
}

type Node struct {
	PublicKey string `json:"publicKey"`
	Host      string `json:"host"`
	Port      uint64 `json:"port"`
	RaftPort  uint64 `json:"raftport"`
}

func Boot(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
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
	consortium := struct {
		Id         string `json:"id"`
		Detail     string `json:"detail"`
		ChainId    uint64 `json:"chainId"`
		Consensus  string `json:"consensus"`
		Difficulty string `json:"difficulty"`
		GasLimit   string `json:"gasLimit"`
		Alloc      Alloc  `json:"alloc"`
		Nodes      []Node `json:"nodes"`
	}{}
	err = database.GetConsortium(client, id, &consortium)
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
	for _, row := range result.Rows {
		if row.Id != common.DOC_GENESIS {
			node := Node{}
			publicKey := strings.ReplaceAll(row.Id, common.DOC_PREFIX, "")
			err = database.GetNode(client, id, publicKey, &node)
			if err != nil {
				common.Response(&w, common.Error, err.Error(), nil)
				return
			}
			consortium.Nodes = append(consortium.Nodes, node)
		}
	}
	common.Response(&w, common.Normal, "Success.", consortium)
}

func List(w http.ResponseWriter, r *http.Request) {
	client, err := database.CouchClient()
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	names, err := client.AllDBs()
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	common.Response(&w, common.Normal, "Success.", names)
	return
}

func New(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	param := struct {
		Id         string `json:"id"`
		Detail     string `json:"detail"`
		ChainId    uint64 `json:"chainId,omitempty"`
		Consensus  string `json:"consensus,omitempty"`
		Difficulty string `json:"difficulty,omitempty"`
		GasLimit   string `json:"gasLimit,omitempty"`
		Alloc      Alloc  `json:"alloc"`
	}{}
	decoder.Decode(&param)

	if strings.ToLower(param.Consensus) == "istanbul" {
		common.Response(&w, common.Error, "Temporarily not supported.", nil)
		return
	}
	if strings.ToLower(param.Consensus) != "raft" && strings.ToLower(param.Consensus) != "istanbul" {
		common.Response(&w, common.Error, "Consensus is invalid.", nil)
		return
	}

	client, err := database.CouchClient()
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	has, err := database.HasConsortium(client, param.Id)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	if has {
		common.Response(&w, common.Error, "consortium "+param.Id+" already exists.", nil)
		return
	}

	_, err = database.AddConsortium(client, param.Id, param)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	common.Response(&w, common.Normal, "Success.", param)
	return
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
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
	if has {
		err = database.DeleteConsortium(client, id)
		if err != nil {
			common.Response(&w, common.Error, err.Error(), nil)
			return
		}
	}
	data := struct {
		Id string `json:"id"`
	}{
		Id: id,
	}
	common.Response(&w, common.Normal, "Success.", data)
	return
}

func Modify(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	decoder := json.NewDecoder(r.Body)
	param := struct {
		Id         string `json:"id"`
		Detail     string `json:"detail"`
		ChainId    uint64 `json:"chainId,omitempty"`
		Consensus  string `json:"consensus,omitempty"`
		Difficulty string `json:"difficulty,omitempty"`
		GasLimit   string `json:"gasLimit,omitempty"`
		Alloc      Alloc  `json:"alloc"`
	}{}
	decoder.Decode(&param)

	client, err := database.CouchClient()
	has, err := database.HasConsortium(client, id)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	if !has {
		common.Response(&w, common.Error, "consortium "+id+" not exist", nil)
	}
	data := struct {
		Rev        string `json:"_rev"`
		Id         string `json:"id"`
		Detail     string `json:"detail"`
		ChainId    uint64 `json:"chainId"`
		Consensus  string `json:"consensus"`
		Difficulty string `json:"difficulty"`
		GasLimit   string `json:"gasLimit"`
		Alloc      Alloc  `json:"alloc"`
	}{}
	err = database.GetConsortium(client, id, &data)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	param.Id = id
	_, err = database.ModifyConsortium(client, id, data.Rev, param)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	common.Response(&w, common.Normal, "Success.", param)
	return
}

func Get(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
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
	param := struct {
		Id         string `json:"id"`
		Detail     string `json:"detail"`
		ChainId    uint64 `json:"chainId"`
		Consensus  string `json:"consensus"`
		Difficulty string `json:"difficulty"`
		GasLimit   string `json:"gasLimit"`
		Alloc      Alloc  `json:"alloc"`
	}{}
	err = database.GetConsortium(client, id, &param)
	if err != nil {
		common.Response(&w, common.Error, err.Error(), nil)
		return
	}
	common.Response(&w, common.Normal, "Success.", param)
}

func Exist(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
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
	data := struct {
		Id    string `json:"id"`
		Exist bool   `json:"exist"`
	}{
		Id:    id,
		Exist: has,
	}
	common.Response(&w, common.Normal, "Success.", data)
}
