package main

import (
	"time"
)

import (
	"github.com/kataras/iris" // go get -u github.com/kataras/iris
)

type Company struct {
	Name          string    `json:"name"`
	Size          int       `json: "size"`
	Location      string    `json: "location"`
	EstablishedOn time.Time `json:"establishedOn"`
}

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
			"location":   "Bangalore",
			"company": Company{
				Name:          "Relevance Lab Pvt Ltd",
				Size:          250,
				Location:      "Bangalore",
				EstablishedOn: time.Now(),
			},
		})
	})

	app.Run(iris.Addr(":8080"))
}
