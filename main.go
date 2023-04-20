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

type PersonController struct {
	persons []Person
}

func main() {

	r := gin.Default()
	defer r.Run(":8080")

	pc := &PersonController{
		persons: []Person{},
	}

	r.GET("/person", pc.getPersons)
	r.GET("/person/:id", pc.getPerson)
	r.POST("person", pc.createPerson)
	r.PUT("/person/:id", pc.updatePerson)
	r.DELETE("/person/:id", pc.deletePerson)

}

func (pc *PersonController) getPersons(c *gin.Context) {
	c.JSON(http.StatusOK, pc.persons)
}

func (pc *PersonController) getPerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for _, person := range pc.persons {
		if person.ID == id {
			c.JSON(http.StatusOK, person)
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "personn not found"})
}

func (pc *PersonController) createPerson(c *gin.Context) {
	var person Person
	err := c.BindJSON(&person)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	person.ID = len(pc.persons) + 1
	pc.persons = append(pc.persons, person)

	c.JSON(http.StatusCreated, person)
}

func (pc *PersonController) updatePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	var newPerson Person
	if err := c.BindJSON(&newPerson); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, person := range pc.persons {
		if person.ID == id {
			newPerson.ID = id
			pc.persons[i] = newPerson

			c.JSON(http.StatusOK, newPerson)
			return
		}
	}
	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "person not found"})
}

func (pc *PersonController) deletePerson(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, person := range pc.persons {
		if person.ID == id {
			pc.persons = append(pc.persons[:i], pc.persons[i+1:]...)

			c.Status(http.StatusNoContent)
			return
		}
	}

	c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"error": "item not found"})

}
