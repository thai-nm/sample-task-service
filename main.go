package main

import (
	"github.com/gin-gonic/gin"
)

type task struct {
	Name        string
	Description string
}

var tasks = []task{
	{
		Name:        "akaPMS",
		Description: "HLS product",
	},
	{
		Name:        "IaC",
		Description: "Internal DevOps project",
	},
}

func getAllTasks(c *gin.Context) {
	c.JSON(200, tasks)
}

func createTask(c *gin.Context) {
	var newTask task

	if err := c.BindJSON(&newTask); err != nil {
		return
	}

	tasks = append(tasks, newTask)
	c.IndentedJSON(200, tasks)
}

func main() {

	server := gin.Default()

	server.GET("/tasks", getAllTasks)
	server.POST("/tasks", createTask)

	server.Run("0.0.0.0:8080")
}
