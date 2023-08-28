package controllers

import (
	"github.com/bitwormhole/web-api-descriptor/app/web/vo"
	"github.com/gin-gonic/gin"
	"github.com/starter-go/libgin"
)

////////////////////////////////////////////////////////////////////////////////

// ExampleController ...
type ExampleController struct {

	//starter:component
	_as func(libgin.Controller) //starter:as(".")

	Responder libgin.Responder //starter:inject("#")

}

func (inst *ExampleController) _impl() libgin.Controller {
	return inst
}

// Registration ...
func (inst *ExampleController) Registration() *libgin.ControllerRegistration {
	return &libgin.ControllerRegistration{
		Route: inst.route,
	}
}

func (inst *ExampleController) route(g *gin.RouterGroup) error {

	g = g.Group("example-ppser")

	g.POST("", inst.handle)
	g.DELETE(":id", inst.handle)
	g.PUT(":id", inst.handle)
	g.GET(":id", inst.handle)
	g.GET("", inst.handle)

	return nil
}

func (inst *ExampleController) handle(c *gin.Context) {
	task := &myExampleTask{
		context:         c,
		parent:          inst,
		wantRequestID:   false,
		wantRequestBody: false,
	}
	task.exec(task.doAction)
}

////////////////////////////////////////////////////////////////////////////////

type myExampleTask struct {
	context *gin.Context
	parent  *ExampleController

	wantRequestID   bool
	wantRequestBody bool

	body1 vo.Example
	body2 vo.Example
}

func (inst *myExampleTask) open() error {

	c := inst.context

	if inst.wantRequestBody {
		err := c.BindJSON(&inst.body1)
		if err != nil {
			return err
		}
	}

	return nil
}

func (inst *myExampleTask) send(err error) {
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

func (inst *myExampleTask) exec(fn func() error) {
	err := inst.open()
	if err == nil {
		err = fn()
	}
	inst.send(err)
}

func (inst *myExampleTask) doAction() error {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
