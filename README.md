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
	config, err := gotral.LoadConfig
	if err != nil { fmt.Println(err) }
	fmt.Println(config)
}
```

