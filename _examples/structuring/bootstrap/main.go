package main

import (
	"github.com/alphayan/iris/_examples/structuring/bootstrap/bootstrap"
	"github.com/alphayan/iris/_examples/structuring/bootstrap/middleware/identity"
	"github.com/alphayan/iris/_examples/structuring/bootstrap/routes"
)

func newApp() *bootstrap.Bootstrapper {
	app := bootstrap.New("Awesome App", "kataras2006@hotmail.com")
	app.Bootstrap()
	app.Configure(identity.Configure, routes.Configure)
	return app
}

func main() {
	app := newApp()
	app.Listen(":8080")
}
