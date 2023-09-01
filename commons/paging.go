package commons

type Paging struct {
		Page  int   `bson:"page" json:"page" form:"page"`
		Limit int   `bson:"limit" json:"limit" form:"limit"`
		Total int64 `bson:"total" json:"total" form:"-"`
}

func (p *Paging) Process() {
	if p.Page <= 0 {
			p.Page = 1
	}

	if p.Limit <= 0 || p.Limit >= 10 {
		p.Limit = 10
	}
}













