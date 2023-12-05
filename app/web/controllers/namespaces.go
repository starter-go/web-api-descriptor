package controllers

import (
	"github.com/bitwormhole/web-api-descriptor/app/web/dto"
	"github.com/bitwormhole/web-api-descriptor/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

// NamespaceController ...
type NamespaceController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

}

func (inst *NamespaceController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *NamespaceController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *NamespaceController) route(g libgin.RouterProxy) error {

	g = g.For("namespaces")

	// g.POST("", inst.handle)
	// g.DELETE(":id", inst.handle)
	// g.PUT(":id", inst.handle)
	// g.GET(":id", inst.handle)

	g.GET("", inst.handleGetList)
	g.GET(":name", inst.handleGetOne)

	return nil
}

func (inst *NamespaceController) handle(c *gin.Context) {
	task := &myNamespaceTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doAction)
}

func (inst *NamespaceController) handleGetList(c *gin.Context) {
	task := &myNamespaceTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doGetList)
}

func (inst *NamespaceController) handleGetOne(c *gin.Context) {
	task := &myNamespaceTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doGetOne)
}

////////////////////////////////////////////////////////////////////////////////

type myNamespaceTask struct {
	context *gin.Context
	parent  *NamespaceController

	wantRequestID   bool
	wantRequestBody bool

	body1 vo.Namespaces
	body2 vo.Namespaces
}

func (inst *myNamespaceTask) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myNamespaceTask) send(err error) {
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

func (inst *myNamespaceTask) exec(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myNamespaceTask) doAction() error {
	return nil
}

func (inst *myNamespaceTask) doGetList() error {

	list := inst.body2.Namespaces

	ns1 := &dto.Namespace{}
	list = append(list, ns1)

	ns2 := &dto.Namespace{}
	list = append(list, ns2)

	ns3 := &dto.Namespace{}
	list = append(list, ns3)

	inst.body2.Namespaces = list
	return nil
}

func (inst *myNamespaceTask) doGetOne() error {

	ns1 := &dto.Namespace{}
	api1 := &dto.API{}

	ns1.APIs = make([]*dto.API, 0)
	ns1.Types = []string{"a", "b", "c", "d"}
	ns1.APIs = append(ns1.APIs, api1)

	api1.Path = "/a/b/c"

	list := inst.body2.Namespaces
	list = append(list, ns1)
	inst.body2.Namespaces = list
	return nil
}

////////////////////////////////////////////////////////////////////////////////
