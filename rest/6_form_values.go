package main

import (
	"fmt"
	"strings"
	"time"
)

import (
	"github.com/kataras/iris" // go get -u github.com/kataras/iris
)

// Company's structure
type Company struct {
	Name          string `json:"name"`
	Size          int    `json:"size"`
	Location      string `json:"location"`
	EstablishedOn string `json:"establishedOn"`
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
			EstablishedOn: strings.Split(fmt.Sprintf("%s", time.Now()), " ")[0],
		}

		// Setting Headers
		ctx.Header("Authorizaion", "^abc#Rishi67^Token$")

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

	// Getting POST data (FormValue(), FormValues())
	app.Handle("POST", "/createUser", func(ctx iris.Context) {
		/*
			* POST
			* Body
				- form-data

			Do not set Header's Content-Type as application/json
			otherwise map[string][]string will be empty
		*/

		// FormValue()
		fmt.Println(ctx.FormValue("fname"))
		fmt.Println(ctx.FormValue("lname"))
		fmt.Println(ctx.FormValue("age"))
		fmt.Println(ctx.FormValue("langs"), "\n")

		// FormValues()
		formValuesMap := ctx.FormValues()
		fmt.Println(formValuesMap)

		// Iterating over POST data
		for key, vArr := range formValuesMap {
			fmt.Println(key)
			for index, value := range vArr {
				fmt.Println("\t", index, value, "\n")
			}
		}

		// Response
		ctx.JSON(User{})
	})

	app.Run(iris.Addr(":8080"))
}
