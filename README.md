# Atomicassets API Client

[![Test](https://github.com/eosswedenorg-go/atomicasset/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/eosswedenorg-go/atomicasset/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/eosswedenorg-go/atomicasset.svg)](https://pkg.go.dev/github.com/eosswedenorg-go/atomicasset)
[![Go Report Card](https://goreportcard.com/badge/github.com/eosswedenorg-go/atomicasset)](https://goreportcard.com/report/github.com/eosswedenorg-go/atomicasset)

This package aims to implement a client for Pinknetwork's [atomicassets API](https://github.com/pinknetworkx/eosio-contract-api) in go.

### Install package

```bash
go get -u github.com/eosswedenorg-go/atomicasset@latest
```

### Example

```go

package main

import (
	"fmt"
	"github.com/eosswedenorg-go/atomicasset"
)

func main() {

	client := atomicasset.New("https://wax.api.atomicassets.io")

	health, err := client.GetHealth()
	if err != nil {
		panic(err)
	}

	fmt.Println("Head block at:", health.Data.Chain.HeadBlock)
}

```

### TODO

* implement `stats` resource

### Author

Henrik Hautakoski - [Sw/eden](https://eossweden.org/) - [henrik@eossweden.org](mailto:henrik@eossweden.org)
