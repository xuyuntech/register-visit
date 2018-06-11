package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (a *Api) test(c *gin.Context) {
	key := c.Query("key")
	if key == "" {
		RespErr(c, fmt.Errorf("need param key"))
		return
	}

	Resp(c, map[string]string{
		"results": "123123",
	})
}
