package markdown

import (
	"github.com/gin-gonic/gin"

	"github.com/thoohv5/bg/library/request"
	"github.com/thoohv5/bg/library/response"
)

type controller struct {
	markdownDir string
}

func New() *controller {
	return &controller{
		markdownDir: "markdown",
	}
}

func (s *controller) Markdown(gtx *gin.Context) {

	req := new(MdReq)
	if err := request.Bind(gtx, req); nil != err {
		response.Error(gtx, err)
		return
	}

	ret, err := NewService().Markdown(req)
	if nil != err {
		response.Error(gtx, err)
		return
	}

	response.Success(gtx, ret, "index.html")
}

func (s *controller) List(gtx *gin.Context) {
	req := new(ListReq)
	if err := request.Bind(gtx, req); nil != err {
		response.Error(gtx, err)
		return
	}

	ret, err := NewService().List(req)
	if nil != err {
		response.Error(gtx, err)
		return
	}

	response.Success(gtx, ret, "success")
}
