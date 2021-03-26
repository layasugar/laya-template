package utils

type HeaderParam struct {
	AppName    string `header:"app-name"`
	RequestID  string `header:"request-id"`
	XTEmployee string `header:"xt-employee"`
}

type XTEmployee struct {
	RealName string `json:"real_name"`
}

type AuthList struct {
	Name       string
	Sign       string
	HttpMethod string
	HttpPath   string
}
