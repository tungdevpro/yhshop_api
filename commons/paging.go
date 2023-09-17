package commons

import "strings"

type Paging struct {
	Page       int    `json:"page" form:"page"`
	Limit      int    `json:"limit" form:"limit"`
	Total      int64  `json:"total" form:"-"`
	FakeCursor string `json:"cursor" form:"cursor"`
	NextCursor string `json:"next_cursor"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
		p.Page = 1
	}

	if p.Limit <= 0 || p.Limit >= 50 {
		p.Limit = 10
	}

	p.FakeCursor = strings.TrimSpace(p.FakeCursor)
}
