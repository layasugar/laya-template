package hall

import (
	"laya-go/controllers"
)

// ControllerSingleton one object of controller
var ControllerSingleton = &controller{}

type controller struct {
	*controllers.BaseController
}
