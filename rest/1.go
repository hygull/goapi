package main

import (
	"github.com/kataras/iris" // go get -u github.com/kataras/iris
)

// main() function starts here
func main() {
	app := iris.Default()

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.JSON(iris.Map{
			"fname":      "Rishikesh",
			"lname":      "Agrawani",
			"age":        20,
			"langs":      []string{"Python", "Golang", "TypeScript", "JavaScript"},
			"profession": "Software development",
			"workingAs":  "Python/Django developer",
			"company":    "Relevance Lab Pvt Ltd",
			"location":   "Bangalore",
		})
	})

	app.Run(iris.Addr(":8080"))
}
