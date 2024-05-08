package tools

type HeaderParam struct {
	AppName   string `header:"app-name"`
	RequestID string `header:"request-id"`
}

type AuthList struct {
	Name       string
	Sign       string
	HttpMethod string
	HttpPath   string
}
