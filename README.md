# go-ebrest 
[![Go Report Card](https://goreportcard.com/badge/github.com/Bolzer/easybillRest)](https://goreportcard.com/report/github.com/bolZer/easybillRest)
[![GoDoc Reference](https://godoc.org/github.com/BolZer/easybillRest?status.svg)](https://godoc.org/github.com/BolZer/easybillRest)
[![Build Status](https://travis-ci.org/BolZer/easybillRest.svg?branch=master)](https://travis-ci.org/BolZer/easybillRest)
[![Generic badge](https://img.shields.io/badge/Version-0.1.0-important.svg)]()

`go-ebrest` is a library to work with the easybill REST API (https://www.easybill.de/api/) written in GO.

All Resources except `attachment` are available.
The missing one will be available shortly.

The library supports only the `Bearer` Authentication and calls the API only
through `HTTPS`.

Documentation will be extended in due time.

```bash
go get -u github.com/BolZer/go-ebrest
```


## Usage

```Go
package main

import (
	"log"
	"github.com/davecgh/go-spew/spew"
	"github.com/BolZer/go-ebrest"
)

func main() {
	client := ebrest.New("API KEY")

	customer, err := client.Customers().CreateCustomer(map[string]string{
		"last_name":    "Test Customer",
		"company_name": "Test Company",
	})
	
	if err != nil {
		log.Fatalf(err.Error())
	}

	customerID := int(customer["id"].(float64))
	spew.Dump(customerID)
}

```