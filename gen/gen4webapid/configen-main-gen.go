package gen4webapid

import "github.com/starter-go/application"

func nop(a ... any) {    
}

func registerComponents(cr application.ComponentRegistry) error {
    ac:=&autoRegistrar{}
    ac.init(cr)
    return ac.addAll()
}

type comFactory interface {
    register(cr application.ComponentRegistry) error
}

type autoRegistrar struct {
    cr application.ComponentRegistry
}

func (inst *autoRegistrar) init(cr application.ComponentRegistry) {
	inst.cr = cr
}

func (inst *autoRegistrar) register(factory comFactory) error {
	return factory.register(inst.cr)
}

func (inst*autoRegistrar) addAll() error {

    
    inst.register(&pe8d8413902_controllers_ExampleController{})
    inst.register(&pe8d8413902_controllers_NamespaceController{})
    inst.register(&pe8d8413902_controllers_SchemaController{})
    inst.register(&pe8d8413902_controllers_ServiceController{})


    return nil
}
