package structs

type PageStruct struct {
	Total    int64       `json:"total"`
	Page     int64       `json:"page"`
	Size     int64       `json:"size"`
	LastPage int64       `json:"last_page"`
	Data     interface{} `json:"data"`
}
