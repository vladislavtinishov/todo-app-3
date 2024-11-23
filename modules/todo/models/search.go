package todomodels

type Search struct {
	Ids         []uint `json:"ids" form:"ids"`
	Title       string `json:"title" form:"title"`
	Description string `json:"description" form:"description"`
	StatusID    uint   `json:"status_id" form:"status_id"`
}
