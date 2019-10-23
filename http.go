/*
* github.com/codenoid - Developer
* code source : - https://github.com/codenoid/GoTral
* 				- https://golang.org/pkg/net/http/
*
 */
package gotral

import (
	"fmt"
	"net/http"

	"github.com/codenoid/gotral"
)

// config : a data structure that come from GoTral server
// usually just {"key": "what", "value": "somethingsecret"}
type config struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

// confret : return config data as string map
// easy to use/get/understand
type confret map[string]string

// LoadConfig : load config from given url and decrypt with given
// passphrase, usually this function only called on app boot time
func LoadConfig(url string, passphrase string) (confret, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	decrypt, err := gotral.Decrypt(body, passphrase)
	if err != nil {
		return nil, fmt.Errorf("failed to decrypt data, got : %v", err.Error())
	}

	// initialize slices of config to receive data from api
	decoded := new([]config)
	err = json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		return nil, err
	}

	// initialize empty map[string]string
	result := make(map[string]string)

	// iterate and put all data into result
	for _, value := range decoded {
		// break and return error if there is duplicate key
		// that come from api
		if val, ok := dict[val.Key]; ok {
			return nil, fmt.Errorf("error: duplicate config key : %v", val.Key)
		}
		result[value.Key] = value.Value
	}

	return result, nil
}
