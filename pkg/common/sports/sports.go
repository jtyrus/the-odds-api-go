package sports

type Sport struct {
	Key          string `json:"key"`
	Group        string `json:"group"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	HasOutrights bool   `json:"has_outrights"`
}