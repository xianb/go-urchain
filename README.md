# go-urchain

A Go client for integrating with the btc.urchain.com API.

## Usage

To use the package, import it in your code:

```

import "github.com/xianb/go-urchain"

```

## Example:

```
package main

import "github.com/xianb/go-urchain"

func main() {
    var opt = urchain.Options{
        ApiKey: "1234567890",
    }
    var cli = urchain.NewClient(opt)
    
    _, err := cli.CheckUrchainHealth()
    if err != nil {
        panic(err)
    }
}

```


