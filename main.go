package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/runabol/tork-demo-codexec/handler"
	"github.com/runabol/tork/cli"
	"github.com/runabol/tork/conf"
	"github.com/runabol/tork/engine"
)

func main() {
	if err := conf.LoadConfig(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	app := cli.New()

	app.CustomizeEngine(func(eng *engine.Engine) error {
		eng.RegisterEndpoint(http.MethodPost, "/exec", handler.Handler)
		return nil
	})

	if err := app.Run(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
