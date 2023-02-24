package action

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/ytlike/goctl-apix/generate"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

func Generator(ctx *cli.Context) error {
	proName := ctx.String("proName")

	if proName == "" {
		return fmt.Errorf("请配置proName")
	}

	p, err := plugin.NewPlugin()
	if err != nil {
		return err
	}
	return generate.Do(p, proName)
}
