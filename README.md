Negroni-authz [![GoDoc](https://godoc.org/github.com/casbin/negroni-authz?status.svg)](https://godoc.org/github.com/casbin/negroni-authz)
======

Negroni-authz is an authorization middleware for [Negroni](https://github.com/urfave/negroni), it's based on [https://github.com/casbin/casbin](https://github.com/casbin/casbin).

## Installation

    go get github.com/casbin/negroni-authz

## Simple Example

```Go
package main

import (
	"fmt"
	"net/http"

	"github.com/casbin/casbin"
	"github.com/casbin/negroni-authz"
	"github.com/urfave/negroni"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "Welcome to the home page!")
	})

	n := negroni.Classic()

	// load the casbin model and policy from files, database is also supported.
	e := casbin.NewEnforcer("authz_model.conf", "authz_policy.csv")
	n.Use(authz.Authorizer(e))

	http.ListenAndServe(":3000", n)
}
```

## How to control the access

The authorization determines a request based on ``{subject, object, action}``, which means what ``subject`` can perform what ``action`` on what ``object``. In this plugin, the meanings are:

1. ``subject``: the logged-on user name
2. ``object``: the URL path for the web resource like "dataset1/item1"
3. ``action``: HTTP method like GET, POST, PUT, DELETE, or the high-level actions you defined like "read-file", "write-blog"

For how to write authorization policy and other details, please refer to [the Casbin's documentation](https://github.com/casbin/casbin).

## Getting Help

- [Casbin](https://github.com/casbin/casbin)

## License

This project is under MIT License. See the [LICENSE](LICENSE) file for the full license text.
