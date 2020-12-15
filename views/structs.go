package views

type Response struct {
	Code int         `json:"code"`
	Body interface{} `json:"body"`
}

type PostRequest struct {
	ID int `json:"ID"`
	Name string `json:"name"`
	Task string `json:"task"`
}
