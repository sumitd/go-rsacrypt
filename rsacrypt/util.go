package rsacrypt

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type PublicKey struct {
	Id  int    `json:"id"`
	Key string `json:"key"`
}
type KeysResponse struct {
	Collection []PublicKey
}

// Fetch the public keys for username
func RSAPublicKey(username string) ([]PublicKey, error) {
	config := ReadConfig()
	content, err := getContent(fmt.Sprintf(config.PublicKeyURL, username))
	if err != nil {
		return nil, err
	}

	keys := make([]PublicKey, 0)

	if err := json.Unmarshal(content, &keys); err != nil {
		return nil, err
	}

	return keys, err
}

// Fetch the content of a URL will return it as an
// array of bytes if retrieved successfully.
func getContent(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	config := ReadConfig()
	req.SetBasicAuth("", config.ApiKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
