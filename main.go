package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// struct
type person struct {
	Id      int    `json: "id"`
	Name    string `json: "name"`
	Age     int    `json: "age"`
	Country string `json: "country"`
}

// slice (temporal database)
var people = []person{
	{Id: 1, Name: "Raisa Ramírez", Age: 30, Country: "El Salvador"},
	{Id: 2, Name: "Habby Lovo", Age: 32, Country: "Colombia"},
	{Id: 3, Name: "Marcela Calero", Age: 22, Country: "España"},
	{Id: 4, Name: "Edith Herrera", Age: 29, Country: "Croacia"},
	{Id: 5, Name: "Paola Pineda", Age: 22, Country: "Guatemala"},
}

// endpoint GET
func getPeople(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{"data": people, "message": "Showing records", "total": len(people)})
}

// endpoint GET by id
func getPerson(c *gin.Context) {
	id := c.Param("id")

	for _, person := range people {
		if fmt.Sprint(person.Id) == id {
			c.IndentedJSON(http.StatusOK, gin.H{"data": person, "message": "Record found it"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

// endpoint delete by id
func deletePerson(c *gin.Context) {
	id := c.Param("id")
	newList := []person{}
	for _, person := range people {
		if fmt.Sprint(person.Id) != id {
			newList = append(newList, person)
		}
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message": "Person " + id + " deleted", "data": newList})
}

// endpoint update by id
func updatePerson(c *gin.Context) {
	id := c.Param("id")
	var personEdited person
	c.BindJSON(&personEdited)

	for i, person := range people {
		if fmt.Sprint(person.Id) == id {
			people[i] = personEdited
			c.IndentedJSON(http.StatusOK, gin.H{"data": people, "message": "Record updated"})
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Person not found"})
}

// endpoint insert by id
func insertPerson(c *gin.Context) {
	var person = person{}
	c.BindJSON(&person)
	people = append(people, person)
	c.IndentedJSON(http.StatusOK, gin.H{"data": people, "message": "Person created"})
}

func main() {
	router := gin.Default()
	// routes
	router.GET("/people", getPeople)
	router.GET("/person/:id", getPerson)
	router.DELETE("/person/:id", deletePerson)
	router.PUT("/person/:id", updatePerson)
	router.POST("/person", insertPerson)

	// local deployment
	router.Run("localhost:8080")

}
