package parser

type Root struct {
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
	Status   uint     `json:"status"`
	Response Response `json:"response"`
}

type Response struct {
	ContentType string `json:"Content-Type"`
	Value       string `json:"value"`
}
