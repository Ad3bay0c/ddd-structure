package handler

import "github.com/gin-gonic/gin"

func (app *Application) Create() gin.HandlerFunc {
	return func(c *gin.Context) {
        c.JSON(200, gin.H{
            "message": "pong",
        })
    }
}
