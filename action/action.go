package action

import (
	"github.com/urfave/cli/v2"
	"github.com/ytlike/goctl-apix/generate"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

func Generator(ctx *cli.Context) error {
	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	return generate.Do(p, ctx)
}
