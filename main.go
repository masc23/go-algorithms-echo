package main

import (
	"github.com/masc23/go-algorithms-echo/core"

	// only imported to run the init function, which sets up all endpoints
	_ "github.com/masc23/go-algorithms-echo/base64"
	_ "github.com/masc23/go-algorithms-echo/mathematics"
	_ "github.com/masc23/go-algorithms-echo/rot13"
)

func main2() {
	core.Start()
}
