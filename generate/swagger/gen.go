package swgger

import (
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"path/filepath"

	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
)

type DocContext struct {
	// Plugin is api date
	Plugin *plugin.Plugin
	// ApiName is the api file name.
	ApiName string
	// SwaggerFile is the swagger file path.
	SwaggerFile string
	// Host is api request address
	Host string
	// BasePath is url request prefix
	BasePath string
	// Output is the output directory of the generated files.
	Output string
}

func (g *Generator) Generate(dctx *DocContext) error {
	abs, err := filepath.Abs(dctx.Output)
	if err != nil {
		return err
	}

	err = pathx.MkdirIfNotExist(abs)
	if err != nil {
		return err
	}

	err = g.GenDoc(dctx)
	if err != nil {
		return err
	}

	return err
}
