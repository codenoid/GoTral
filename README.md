# WIP: GoTral

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

	config, err := gotral.LoadConfig("http://someurlto.gotral.api", secret)
	if err != nil { fmt.Println(err) }
	fmt.Println(config)
}
```

