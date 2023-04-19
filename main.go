package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type Person struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Lastname string `json:"lastname"`
}

var persons []Person

func main() {

	r := gin.Default()
	defer r.Run(":8080")

	r.GET("/person", getPersons)
	r.GET("/person/:id", getPerson)
	r.POST("person", createPerson)
	r.PUT("person/id", updatePerson)
	r.DELETE("person/id", deletePerson)

}

func getPersons(c *gin.Context) {
	c.JSON(http.StatusOK, persons)
}

func getPerson(c *gin.Context) {
	id := c.Param("id")
	for _, u := range persons {
		if strconv.Itoa(u.ID) == id {
			c.JSON(http.StatusOK, u)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "that person doesn't exist"})
}

func createPerson(c *gin.Context) {
	var person Person
	err := c.BindJSON(&person)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	person.ID = len(persons) + 1
	persons = append(persons, person)
	c.JSON(http.StatusCreated, person)
}
func updatePerson(c *gin.Context) {
	id := c.Param("id")
	var NewPerson Person
	err := c.BindJSON(&NewPerson)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	for i, u := range persons {
		if strconv.Itoa(u.ID) == id {
			NewPerson.ID = u.ID
			persons[i] = NewPerson
			c.JSON(http.StatusOK, NewPerson)
		}
	}
}

func deletePerson(c *gin.Context) {
	id := c.Param("id")
	for i, u := range persons {
		if strconv.Itoa(u.ID) == id {
			persons = append(persons[:i], persons[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"message": "User Delete"})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "Person Not Found"})
}
