package common

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
}

func New() *controller {
	return &controller{}
}

// NoRoute
func (c *controller) NoRoute(gtx *gin.Context) {
	gtx.HTML(http.StatusNotFound, "404.html", nil)
}

// NoMethod
func (c *controller) NoMethod(gtx *gin.Context) {
	gtx.HTML(http.StatusNotFound, "404.html", nil)
}
