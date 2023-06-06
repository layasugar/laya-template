package edb

import (
	"log"
	"net/http"

	"github.com/olivere/elastic/v7"
)

// dbConfig Cluser Base Config
type dbConfig struct {
	name string
	dsn  string
	user string
	pwd  string
}

// Open 开启连接
func (c *dbConfig) Open() *elastic.Client {
	tr := NewTransport()
	httpClient := &http.Client{
		Transport: tr,
	}

	// Create a client
	client, err := elastic.NewClient(
		elastic.SetURL(c.dsn),
		elastic.SetBasicAuth(c.user, c.pwd),
		elastic.SetHttpClient(httpClient),
		elastic.SetHealthcheck(false),
		elastic.SetSniff(false),
	)
	if err != nil {
		panic(err)
	}
	log.Printf("[app.edb] es success, name: %s", c.name)
	return client
}

func InitConn(m []map[string]interface{}) {
	for _, item := range m {
		var dbc = dbConfig{}

		if name, ok := item["name"]; ok {
			if nameStr, okInterface := name.(string); okInterface {
				if nameStr == "" {
					dbc.name = defaultEdbName
				} else {
					dbc.name = nameStr
				}
			}
		} else {
			dbc.name = defaultEdbName
		}

		if dsn, ok := item["dsn"]; ok {
			if dsnStr, okInterface := dsn.(string); okInterface {
				dbc.dsn = dsnStr
			}
		}

		if user, ok := item["user"]; ok {
			if userStr, okInterface := user.(string); okInterface {
				dbc.user = userStr
			}
		}

		if pwd, ok := item["pwd"]; ok {
			if pwdStr, okInterface := pwd.(string); okInterface {
				dbc.pwd = pwdStr
			}
		}

		setEdb(dbc.name, dbc.Open())
	}
}

func GetClient(name ...string) *elastic.Client {
	if len(name) > 0 {
		return getEdb(name[0])
	} else {
		return getEdb(defaultEdbName)
	}
}
