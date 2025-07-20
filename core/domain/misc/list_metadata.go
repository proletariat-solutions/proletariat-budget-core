package misc

type ListMetadata struct {
	Total  uint `json:"total"`
	Limit  uint `json:"limit"`
	Offset uint `json:"offset"`
}

type ListParams struct {
	Limit  uint `json:"limit"`
	Offset uint `json:"offset"`
}
