package common

import (
	"github.com/gin-gonic/gin"
	"github.com/mitchellh/mapstructure"
)

func Convert(c *gin.Context, source interface{}, target interface{}) error {
	if err := c.ShouldBind(source); err != nil {
		return err
	}
	if err := mapstructure.Decode(source, target); err != nil {
		return err
	}
	return nil
}
