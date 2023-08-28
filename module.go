package webapidescriptor

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modulegin"
	modulesecuritygingorm "github.com/starter-go/module-security-gin-gorm"

	"github.com/bitwormhole/web-api-descriptor/gen/gen4webapid"
)

const (
	theModuleName     = "github.com/bitwormhole/web-api-descriptor"
	theModuleVersion  = "v0.0.0"
	theModuleRevision = 0
	theModuleResPath  = "src/main/resources"
)

//go:embed "src/main/resources"
var theModuleResFS embed.FS

// Module 导出模块 "github.com/bitwormhole/passport-server"
func Module() application.Module {
	mb := application.ModuleBuilder{}
	mb.Name(theModuleName)
	mb.Version(theModuleVersion)
	mb.Revision(theModuleRevision)

	mb.EmbedResources(theModuleResFS, theModuleResPath)
	mb.Components(gen4webapid.ExportComForWebAPIDesc)

	mb.Depend(modulegin.Module())
	mb.Depend(modulesecuritygingorm.Module())
	return mb.Create()
}
