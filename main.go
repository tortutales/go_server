package main
import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)
type Day struct {
    Id int `json:"id"`
    Date string `json:"date"`
    Color string `json:"color"`
}
func dayHandler(c *gin.Context) {
    days := []Day {
        Day {0, "2022-06-03", "0xFF00FF00"},
        Day {0, "2022-06-04", "0xFF00FF00"},
    }
    c.JSON(200, gin.H{
        "days": days,
    })
}
func main() {
    r := gin.Default()
    r.Use(cors.Default())

    r.GET("/days", dayHandler)

    r.Run(":5000")
}
