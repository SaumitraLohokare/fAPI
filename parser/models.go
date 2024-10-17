package parser

type Context struct {
	Port   uint16  `json:"port"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Url  string  `json:"url"`
	Get  Request `json:"GET"`
	Post Request `json:"POST"`
	// TODO: Add other types of requests
}

type Request struct {
	Status      int    `json:"status"`
	ContentType string `json:"Content-Type"`
	Response    string `json:"response"`
}
