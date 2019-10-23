# WIP: GoTral [![Go Report Card](https://goreportcard.com/badge/github.com/codenoid/GoTral)](https://goreportcard.com/report/github.com/codenoid/GoTral)

Encrypted Golang Centralized Configuration management

## Install

Just like usual way

```
go get -u github.com/codenoid/gotral
```

## Example Usage

```
package main

import (
	"fmt"

	"github.com/codenoid/gotral"
)

func main() {

	// super secret key
	secret := string([]byte{97, 121, 97, 109}) // or just put string in there

	config, err := gotral.DirectLoad("http://someurlto.gotral.api", secret)
	if err != nil { fmt.Println(err) }
	fmt.Println(config)

	// with basic auth support
	withOpt := gotral.GoTral{
		Url: "https://jigsaw.w3.org/HTTP/Basic",
		Passphrase: "DecryptPassword",
		BasicAuth: true,
		Username: "guest",
		Password: "guest",
	}

	config, err = withOpt.LoadConfig()
	if err != nil { fmt.Println(err) }
	fmt.Println(config)
}

```

