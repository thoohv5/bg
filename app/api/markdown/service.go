package markdown

import (
	"fmt"
	"html/template"
	"strings"

	"github.com/thoohv5/bg/app/service/file"
	"github.com/thoohv5/bg/app/service/markdown"
)

type (
	service struct {
		markdownDir string
	}
	IService interface {
		Markdown(req *MdReq) (*MdResp, error)
		List(req *ListReq) (*ListResp, error)
	}
)

func NewService() IService {
	return &service{
		markdownDir: "markdown",
	}
}

func (s *service) Markdown(req *MdReq) (*MdResp, error) {
	resp := new(MdResp)
	name := req.Label

	fileName := name
	if params := strings.Split(name, "."); len(params) > 0 {
		fileName = params[0]
	}
	if fileNames := strings.Split(fileName, "-"); len(fileNames) > 0 {
		fileName = strings.Join(fileNames, "/")
	}

	// content
	content, err := markdown.New().Parse(fmt.Sprintf("./%s/%s.md", s.markdownDir, fileName))
	if nil != err {
		return resp, err
	}

	resp.Name = fileName
	resp.Content = template.HTML(content)
	return resp, nil
}

func (s *service) List(req *ListReq) (*ListResp, error) {
	resp := new(ListResp)

	// list
	files, err := file.New().Walk(fmt.Sprintf("./%s", s.markdownDir))
	if nil != err {
		return resp, err
	}
	list := make(map[string][]string, 0)
	for _, f := range files {
		fs := strings.Split(strings.Split(f.Name, s.markdownDir)[1], "/")[1]
		if strings.HasSuffix(f.Name, ".md") {
			list[fs] = append(list[fs], strings.Split(strings.Trim(strings.Split(f.Name, s.markdownDir)[1], "/.md"), "/")[1])
		}
	}

	resp.List = list
	return resp, nil
}
