package middleware

import (
	"github.com/BurntSushi/toml"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"laya-go/ship"
	"laya-go/ship/response"
	"net/http"
	"strings"
)

func Response() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		if c.Writer.Written() {
			return
		}

		params := c.Keys
		if len(params) == 0 {
			return
		}

		lang := GetLang(c.GetHeader("Accept-Language"))
		resp := GetResponse(params, lang)
		c.JSON(http.StatusOK, resp)
	}
}

func GetLang(lang string) string {
	if lang == "" {
		lang = ship.DefaultLang
	}
	return lang
}

func GetResponse(params map[string]interface{}, lang string) interface{} {
	var resp response.Response
	for name, value := range params {
		if !strings.HasPrefix(name, "$.") {
			continue
		}
		lastOne := strings.Split(name, ".")[len(strings.Split(name, "."))-1]

		switch lastOne {
		case "code":
			resp.Code = value.(int)
		case "response":
			op, _ := value.(response.Response)
			resp = op
		}
		resp.Location = name
		resp.Msg = GetMessage(lang)
	}
	return resp
}

func GetMessage(lang string) string {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)
	bundle.LoadMessageFile("/ship/lang/zh.toml")
	bundle.LoadMessageFile("/ship/lang/en.toml")


	loc := i18n.NewLocalizer(bundle, lang)

	return loc.MustLocalize(&i18n.LocalizeConfig{
		DefaultMessage: &i18n.Message{
			ID:    "PersonCats",
			One:   "{{.Name}} has {{.Count}} cat.",
			Other: "{{.Name}} has {{.Count}} cats.",
		},
		TemplateData: map[string]interface{}{
			"Name":  "Nick",
			"Count": 2,
		},
		PluralCount: 1,
	})
}
