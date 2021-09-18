package structs

type PageStruct struct {
	Total    int64
	Page     int64
	Size     int64
	LastPage int64
	Data     interface{}
}
