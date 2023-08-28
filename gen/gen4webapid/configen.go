package gen4webapid

import "github.com/starter-go/application"

//starter:configen(version="4")

// ExportComForWebAPIDesc 导出组件配置
func ExportComForWebAPIDesc(cr application.ComponentRegistry) error {
	return registerComponents(cr)
	// return nil
}
