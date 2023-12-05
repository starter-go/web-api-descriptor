package controllers

import (
	"github.com/bitwormhole/web-api-descriptor/app/web/dto"
	"github.com/bitwormhole/web-api-descriptor/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

// ServiceController ...
type ServiceController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

}

func (inst *ServiceController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *ServiceController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *ServiceController) route(g libgin.RouterProxy) error {

	g = g.For("services")

	// g.POST("", inst.handle)
	// g.DELETE(":id", inst.handle)
	// g.PUT(":id", inst.handle)
	// g.GET(":id", inst.handle)
	// g.GET("", inst.handle)

	g.GET("", inst.handleGetList)

	g.GET("/index.json", inst.handleGetList)

	return nil
}

func (inst *ServiceController) handle(c *gin.Context) {
	task := &myServiceTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doAction)
}

func (inst *ServiceController) handleGetList(c *gin.Context) {
	task := &myServiceTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doGetList)
}

////////////////////////////////////////////////////////////////////////////////

type myServiceTask struct {
	context *gin.Context
	parent  *ServiceController

	wantRequestID   bool
	wantRequestBody bool

	body1 vo.Services
	body2 vo.Services
}

func (inst *myServiceTask) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myServiceTask) send(err error) {
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

func (inst *myServiceTask) exec(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myServiceTask) doAction() error {
	return nil
}

func (inst *myServiceTask) doGetList() error {

	list := inst.body2.Services

	s1 := &dto.Service{}
	list = append(list, s1)

	s2 := &dto.Service{}
	list = append(list, s2)

	s3 := &dto.Service{}
	list = append(list, s3)

	inst.body2.Services = list
	return nil
}

////////////////////////////////////////////////////////////////////////////////
