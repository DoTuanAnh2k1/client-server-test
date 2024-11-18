package server

import "github.com/gin-gonic/gin"

func NewGin() *gin.Engine {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "ok")
	})
	r.GET("/info", getInfo)

	r.GET("/trigger/on/:name", triggerOn)
	r.GET("/trigger/off/:name", triggerOff)
	r.GET("/trigger/problem/:name", triggerProblem)
	r.GET("/trigger/measure", getMeasure)

	return r
}
