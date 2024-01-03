package template

type Response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
}


type ResponsePagination struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Page    int         `json:"page"`
	Perpage int         `json:"perpage"`
	Total   int64       `json:"total"`
}
