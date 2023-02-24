package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/ytlike/goctl-apix/action"
	"github.com/ytlike/goctl-apix/generate"
	"github.com/zeromicro/go-zero/tools/goctl/api/parser"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"log"
	"os"
	"runtime"
)

var (
	version  = "20220621"
	commands = []*cli.Command{
		{
			Name:   "apix",
			Usage:  "from api file to zrpc project and swagger doc file",
			Action: action.Generator,
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "proName",
					Usage: "project name",
				},
			},
		},
	}
)

func main1() {
	app := cli.NewApp()
	app.Usage = "a plugin of goctl to generate zrpc project and swagger doc file"
	app.Version = fmt.Sprintf("%s %s/%s", version, runtime.GOOS, runtime.GOARCH)
	app.Commands = commands
	if err := app.Run(os.Args); err != nil {
		fmt.Printf("goctl-protobuf: %+v\n", err)
	}
}

func main() {
	api, err := parser.Parse("./test.api")
	if err != nil {
		log.Fatalf("api parse err:%v", err)
	}

	p := &plugin.Plugin{
		Api:         api,
		ApiFilePath: "./",
		Style:       "gozero",
		Dir:         ".",
	}

	err = generate.Do(p, "test")
	if err != nil {
		log.Fatalf("generate err:%v", err)
	}
}
