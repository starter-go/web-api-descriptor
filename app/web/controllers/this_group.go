package controllers

import "github.com/gin-gonic/gin"

func makeGroupFor(g *gin.RouterGroup, typeName string) *gin.RouterGroup {
	return g.Group("/api/demo/" + typeName)
}
