package main

import (
    "fmt"
    "net/http"
    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()
    router.GET("/sensors", getsensors)
    router.GET("/sensors/:id", getsensorByID)
    router.POST("/sensors", postsensor)

    fmt.Println("Server will run on http://localhost:8000")
    router.Run(":8000")
}

type Sensor struct {
    Sensor string `json:"sensor"`
    Value  string `json:"value"`
}

// getsensors responds with the list of all sensors as JSON.
func getsensors(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, GetSensors())
}

// postsensors adds an sensor from JSON received in the request body.
func postsensor(c *gin.Context) {
    var data Sensor

    if err := c.BindJSON(&data); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    c.IndentedJSON(http.StatusCreated, PostSensor(data.Sensor, data.Value))
}

func getsensorByID(c *gin.Context) {
    id := c.Param("id")

    sensors := GetSensors()

    // Loop through the list of sensors, looking for
    // an sensor whose id value matches the parameter.
    for _, a := range sensors {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "sensor not found"})
}