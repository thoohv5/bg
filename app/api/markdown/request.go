package markdown

import "html/template"

// markdown 请求
type MdReq struct {
	// label
	Label string `form:"label" json:"label" param:"label"`
}

// markdown 返回值
type MdResp struct {
	Name    string
	Content template.HTML
}

// list 请求
type ListReq struct {
}

// list 返回值
type ListResp struct {
	List map[string][]string
}
