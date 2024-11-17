package view

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *userView) Index(c *gin.Context) {
    // session := sessions.Default(c)

    c.HTML(http.StatusOK, "index.html", gin.H{
        "title":       "Index",
    })
}