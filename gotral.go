/**
* github.com/codenoid - Developer
* code source : - https://github.com/codenoid/GoTral
*               - https://golang.org/pkg/net/http/
*/
package gotral

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

import "github.com/mervick/aes-everywhere/go/aes256"

// GoTral : enabling uses of basic auth
type GoTral struct {
	Url        string
	Passphrase string
	BasicAuth  bool
	Username   string
	Password   string
}

// config : a data structure that come from GoTral server
// usually just {"key": "what", "value": "somethingsecret"}
type config struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// confret : return config data as string map
// easy to use/get/understand
type confret map[string]string

// Get : returns a value or an error for a key
func (config confret) Get(key string) (string, error) {
	if len(config) == 0 {
		return "", fmt.Errorf("The config is empty")
	}
	if val, ok := config[key]; ok {
		return val, nil
	}
	return "", fmt.Errorf("The key doesn't exist")
}

// DirectLoad : directly load config from given url and decrypt with given passphrase
// usually this function only called on app boot time
func DirectLoad(url string, passphrase string) (confret, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401
	if resp.StatusCode == 401 {
		return nil, fmt.Errorf("Unauthorized, username & password for basic auth needed")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	decrypt := aes256.Decrypt(string(body), passphrase)

	// initialize slices of config to receive data from api
	var decoded []config
	err = json.Unmarshal([]byte(decrypt), &decoded)
	if err != nil {
		return nil, err
	}

	// initialize empty map[string]string
	result := make(map[string]string)

	// iterate and put all data into result
	for _, value := range decoded {
		// break and return error if there is duplicate key
		// that come from api
		if _, ok := result[value.Key]; ok {
			return nil, fmt.Errorf("error: duplicate config key : %v", value.Key)
		}
		result[value.Key] = value.Value
	}

	return result, nil
}

// LoadConfig : basic auth version support
func (r GoTral) LoadConfig() (confret, error) {

	if r.BasicAuth == false {
		val, err := DirectLoad(r.Url, r.Passphrase)
		if err != nil {
			return nil, err
		}
		return val, nil
	}

	// initialize http client with/out option
	client := &http.Client{}

	// create new request by given url
	req, err := http.NewRequest("GET", r.Url, nil)
	if err != nil {
		return nil, err
	}

	// pass auth for BasicAuth
	req.SetBasicAuth(r.Username, r.Password)

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	// https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/401
	if resp.StatusCode == 401 {
		return nil, fmt.Errorf("Unauthorized, wrong username/password")
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	decrypt := aes256.Decrypt(string(body), r.Passphrase)

	// initialize slices of config to receive data from api
	var decoded []config
	err = json.Unmarshal([]byte(decrypt), &decoded)
	if err != nil {
		return nil, err
	}

	// initialize empty map[string]string
	result := make(map[string]string)

	// iterate and put all data into result
	for _, value := range decoded {
		// break and return error if there is duplicate key
		// that come from api
		if _, ok := result[value.Key]; ok {
			return nil, fmt.Errorf("error: duplicate config key : %v", value.Key)
		}
		result[value.Key] = value.Value
	}

	return result, nil
}

