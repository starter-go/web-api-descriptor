package gen4webapid
import (
    pe8d841390 "github.com/bitwormhole/web-api-descriptor/app/web/controllers"
    pd1a916a20 "github.com/starter-go/libgin"
     "github.com/starter-go/application"
)

// type pe8d841390.ExampleController in package:github.com/bitwormhole/web-api-descriptor/app/web/controllers
//
// id:com-e8d84139021d4b21-controllers-ExampleController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pe8d8413902_controllers_ExampleController struct {
}

func (inst* pe8d8413902_controllers_ExampleController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e8d84139021d4b21-controllers-ExampleController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe8d8413902_controllers_ExampleController) new() any {
    return &pe8d841390.ExampleController{}
}

func (inst* pe8d8413902_controllers_ExampleController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe8d841390.ExampleController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)


    return nil
}


func (inst*pe8d8413902_controllers_ExampleController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}



// type pe8d841390.NamespaceController in package:github.com/bitwormhole/web-api-descriptor/app/web/controllers
//
// id:com-e8d84139021d4b21-controllers-NamespaceController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pe8d8413902_controllers_NamespaceController struct {
}

func (inst* pe8d8413902_controllers_NamespaceController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e8d84139021d4b21-controllers-NamespaceController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe8d8413902_controllers_NamespaceController) new() any {
    return &pe8d841390.NamespaceController{}
}

func (inst* pe8d8413902_controllers_NamespaceController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe8d841390.NamespaceController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)


    return nil
}


func (inst*pe8d8413902_controllers_NamespaceController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}



// type pe8d841390.ServiceController in package:github.com/bitwormhole/web-api-descriptor/app/web/controllers
//
// id:com-e8d84139021d4b21-controllers-ServiceController
// class:class-d1a916a203352fd5d33eabc36896b42e-Controller
// alias:
// scope:singleton
//
type pe8d8413902_controllers_ServiceController struct {
}

func (inst* pe8d8413902_controllers_ServiceController) register(cr application.ComponentRegistry) error {
	r := cr.NewRegistration()
	r.ID = "com-e8d84139021d4b21-controllers-ServiceController"
	r.Classes = "class-d1a916a203352fd5d33eabc36896b42e-Controller"
	r.Aliases = ""
	r.Scope = "singleton"
	r.NewFunc = inst.new
	r.InjectFunc = inst.inject
	return r.Commit()
}

func (inst* pe8d8413902_controllers_ServiceController) new() any {
    return &pe8d841390.ServiceController{}
}

func (inst* pe8d8413902_controllers_ServiceController) inject(injext application.InjectionExt, instance any) error {
	ie := injext
	com := instance.(*pe8d841390.ServiceController)
	nop(ie, com)

	
    com.Responder = inst.getResponder(ie)


    return nil
}


func (inst*pe8d8413902_controllers_ServiceController) getResponder(ie application.InjectionExt)pd1a916a20.Responder{
    return ie.GetComponent("#alias-d1a916a203352fd5d33eabc36896b42e-Responder").(pd1a916a20.Responder)
}


