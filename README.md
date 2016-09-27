# iris
package main

import "github.com/kataras/iris"

func main() {
    iris.Get("/hello", func(ctx *iris.Context) {
        ctx.Write("Hello %s", "iris")
    })
    iris.Listen(":8080")
}
