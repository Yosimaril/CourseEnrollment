package helpers

func ParseUintParam(c *gin.Context, name string) (uint, error) {
	value := c.Param(name)

	id, err := strconv.ParseUint(value, 10, 64)

	return uint(id), err
}
