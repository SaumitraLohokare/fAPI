package parser

type Context struct {
	Port   uint16  `json:"port"`
	Routes []Route `json:"routes"`
}

type Route struct {
	Url    string  `json:"url"`
	Get    Request `json:"GET"`
	Post   Request `json:"POST"`
	Put    Request `json:"PUT"`
	Delete Request `json:"DELETE"`
}

type Request struct {
	Status      int    `json:"status"`
	ContentType string `json:"Content-Type"`
	Response    string `json:"response"`
}
