package main

import (
	"time"
)

import (
	"github.com/kataras/iris" // go get -u github.com/kataras/iris
)

// Company's structure
type Company struct {
	Name          string    `json:"name"`
	Size          int       `json:"size"`
	Location      string    `json:"location"`
	EstablishedOn time.Time `json:"establishedOn"`
}

// User's structure
type User struct {
	Fname      string   `json:"fname"`
	Lname      string   `json:"lname"`
	Age        int      `json:"age"`
	Langs      []string `json:"langs"`
	Profession string   `json:"profession"`
	WorkingAs  string   `json:"workingAs"`
	Location   string   `json:"location"`
	Company    Company  `json:"company"`
}

// main() function starts here
func main() {
	app := iris.Default()

	app.Handle("GET", "/", func(ctx iris.Context) {
		company := Company{
			Name:          "Relevance Lab Pvt Ltd",
			Size:          250,
			Location:      "Sarjapura, Bangalore",
			EstablishedOn: time.Now(),
		}

		ctx.JSON(User{
			Fname:      "Rishikesh",
			Lname:      "Agrawani",
			Age:        25,
			Langs:      []string{"Python", "Golang", "Typescript", "JavaScript"},
			Profession: "Software Development",
			WorkingAs:  "Software Developer",
			Location:   "Bangalore",
			Company:    company,
		})
	})

	app.Run(iris.Addr(":8080"))
}
