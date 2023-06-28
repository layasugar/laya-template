package constants

// ServerType 服务类型
type ServerType int

// RunMode 运行模式
type RunMode string

type RequestLog struct {
	URL     string              `json:"url"`
	Method  string              `json:"method"`
	IP      []string            `json:"ip"`
	Path    string              `json:"path"`
	Headers map[string][]string `json:"headers"`
	Query   interface{}         `json:"query"`
	Body    interface{}         `json:"body"`
}
