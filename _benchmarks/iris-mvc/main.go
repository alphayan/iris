package main

/// TODO: remove this on the "master" branch, or even replace it
// with the "iris-mvc" (the new implementatioin is even faster, close to handlers version,
// with bindings or without).

import (
	"github.com/alphayan/iris/_benchmarks/iris-mvc/controllers"

	"github.com/alphayan/iris"
	"github.com/alphayan/iris/mvc"
)

func main() {
	app := iris.New()
	mvc.New(app.Party("/api/values/{id}")).
		Handle(new(controllers.ValuesController))

	app.Run(iris.Addr(":5000"), iris.WithoutVersionChecker)
}

// +2MB/s faster than the previous implementation, 0.4MB/s difference from the raw handlers.
