package controllers

import (
	"github.com/bitwormhole/web-api-descriptor/app/web/dto"
	"github.com/bitwormhole/web-api-descriptor/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

// SchemaController ...
type SchemaController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

}

func (inst *SchemaController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *SchemaController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *SchemaController) route(g libgin.RouterProxy) error {

	g = g.For("schemata")

	// g.POST("", inst.handle)
	// g.DELETE(":id", inst.handle)
	// g.PUT(":id", inst.handle)
	// g.GET(":id", inst.handle)

	g.GET("", inst.handleGetList)
	g.GET(":name", inst.handleGetOne)

	return nil
}

func (inst *SchemaController) handle(c *gin.Context) {
	task := &mySchemaTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doAction)
}

func (inst *SchemaController) handleGetList(c *gin.Context) {
	task := &mySchemaTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doGetList)
}

func (inst *SchemaController) handleGetOne(c *gin.Context) {
	task := &mySchemaTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doGetOne)
}

////////////////////////////////////////////////////////////////////////////////

type mySchemaTask struct {
	context *gin.Context
	parent  *SchemaController

	wantRequestID   bool
	wantRequestBody bool

	body1 vo.Schema
	body2 vo.Schema
}

func (inst *mySchemaTask) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *mySchemaTask) send(err error) {
	c := inst.context
	d := &inst.body2
	s := inst.body2.Status
	r := &libgin.Response{
		Context: c,
		Data:    d,
		Error:   err,
		Status:  s,
	}
	inst.parent.Responder.Send(r)
}

func (inst *mySchemaTask) exec(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *mySchemaTask) doAction() error {
	return nil
}

func (inst *mySchemaTask) doGetList() error {

	schema := &inst.body2.Schema
	nslist := inst.body2.Namespaces
	ns := &dto.Namespace{}
	api := &dto.API{}

	schema.URL = "(url of this doc)"

	ns.Types = append(ns.Types, "xxx")
	ns.Types = append(ns.Types, "yyy")
	ns.Types = append(ns.Types, "zzz")

	ns.APIs = append(ns.APIs, api)
	ns.APIs = append(ns.APIs, api)

	api.Params = append(api.Params, "a")
	api.Params = append(api.Params, "b")
	api.Params = append(api.Params, "c")

	nslist = append(nslist, ns)
	inst.body2.Namespaces = nslist
	return nil
}

func (inst *mySchemaTask) doGetOne() error {

	// 	ns1 := &dto.Schema{}
	// 	api1 := &dto.API{}

	// 	ns1.APIs = make([]*dto.API, 0)
	// 	ns1.Types = []string{"a", "b", "c", "d"}
	// 	ns1.APIs = append(ns1.APIs, api1)

	// 	api1.Path = "/a/b/c"

	// 	list := inst.body2.Schema
	// 	list = append(list, ns1)
	// 	inst.body2.Schema = list

	return nil
}

////////////////////////////////////////////////////////////////////////////////
