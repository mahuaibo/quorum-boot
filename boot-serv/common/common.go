package common

import "strings"

func GetConsortiumName(name string) string {
	return DB_PREFIX + name
}

func GetNodeName(publicKey string) string {
	return DOC_PREFIX + publicKey
}

func HandlePublicKey(publicKey string) string {
	publicKey = strings.ToLower(publicKey)
	if strings.HasPrefix(publicKey, "0x04") {
		publicKey = strings.ReplaceAll(publicKey, "0x04", "")
	} else if strings.HasPrefix(publicKey, "0x") {
		publicKey = strings.ReplaceAll(publicKey, "0x", "")
	}
	return publicKey
}
