package helpers

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseUintParam(c *gin.Context, name string) (uint, error) {
	value := c.Param(name)

	id, err := strconv.ParseUint(value, 10, 64)

	return uint(id), err
}

func ParseUintQuery(c *gin.Context, name string) (uint, error) {
	value := c.Query(name)

	id, err := strconv.ParseUint(value, 10, 64)

	return uint(id), err
}
