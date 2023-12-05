package webapidescriptor

import (
	"embed"

	"github.com/starter-go/application"
	"github.com/starter-go/libgin/modules/libgin"
	"github.com/starter-go/security-gin-gorm/modules/securitygingorm"

	"github.com/bitwormhole/web-api-descriptor/gen/gen4webapid"
)

const (
	theModuleName     = "github.com/bitwormhole/web-api-descriptor"
	theModuleVersion  = "v0.0.1"
	theModuleRevision = 1
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

	mb.Depend(libgin.Module())
	mb.Depend(securitygingorm.Module())
	return mb.Create()
}
