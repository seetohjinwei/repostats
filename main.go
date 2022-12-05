package main

import (
	"flag"
	"fmt"

	"github.com/seetohjinwei/repostats/cli"
	"github.com/seetohjinwei/repostats/messages"
	"github.com/seetohjinwei/repostats/web"
)

var modes = map[string]func(){
	"cli": cli.Start,
	"web": web.Start,
}

func main() {
	maybeMode := flag.String("mode", "cli", "either cli or web")
	flag.Parse()

	found := false
	for mode, app := range modes {
		if *maybeMode == mode {
			app()
			return
		}
	}
	if !found {
		fmt.Printf(messages.INVALID_MODE, *maybeMode)
	}
}
