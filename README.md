# GoTral [![Go Report Card](https://goreportcard.com/badge/github.com/codenoid/GoTral)](https://goreportcard.com/report/github.com/codenoid/GoTral)

Encrypted Golang Centralized Configuration management

## Feature

- E2E Encryption (via http/s)
- Basic auth support
- Multiple config file in single server
- Easy to use, simple API and return as map[string]string
- ....and more (/next update)

## Setup & Install

Make sure you already have [GoTral server](https://github.com/codenoid/GoTral-Server) up

For the library just like usual way

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
	secret := "somehardpw" // or just put string in there

	config, err := gotral.DirectLoad("http://localhost:6969/config?id=ecommerce.json", secret)
	if err != nil { fmt.Println(err) }
	if val, err := config.Get("mysql_username"); !err {
		fmt.Println(val)
	}

	// with basic auth support
	withOpt := gotral.GoTral{
		Url: "http://localhost:6969/config?id=ecommerce.json",
		Passphrase: "somehardpw",
		BasicAuth: true,
		Username: "guest",
		Password: "guest",
	}

	config, err = withOpt.LoadConfig()
	if err != nil { fmt.Println(err) }
	if val, err := config.Get("mysql_username"); !err {
		fmt.Println(val)
	}
}
```
