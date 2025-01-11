package requests

type MethodAPI struct {
	POST    string            `json:"POST"`
	PUT     string            `json:"PUT"`
	GET     string            `json:"GET"`
	DELETE  string            `json:"DELETE"`
	OPTIONS string            `json:"OPTIONS"`
	Headers map[string]string `json:"headers"`
}