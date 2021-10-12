package handlers

// HTTPError defines the field which should be the response.
type HTTPError struct {
	Status string `json:"status" example:"error"`          // The status error of string.
	Error  string `json:"error" example:"a message error"` // The message error description of string.
}

// HTTPSuccess defines the field which should be the response.
type HTTPSuccess struct {
	Status string      `json:"status" example:"success"` // The status success of string.
	Data   interface{} `json:"data"`                     // The result of json.
}
