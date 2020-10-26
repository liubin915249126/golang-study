package get

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func NameQuery(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // 是 c.Request.URL.Query().Get("lastname") 的简写
	c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
}