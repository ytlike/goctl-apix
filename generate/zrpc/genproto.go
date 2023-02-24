package zrpc

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"os"
	"text/template"
)

//go:embed tpl/proto.tpl
var protoTemplate string

func (g *Generator) GenProto(p *plugin.Plugin, proName string) error {
	protoTpl := template.Must(template.New("protoTpl").Parse(protoTemplate))
	fp, err := os.OpenFile(p.ApiFilePath+"/"+proName+".proto", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return fmt.Errorf("GenProto not found proto file err:%v", err)
	}

	protoTpl.Execute(fp, map[string]interface{}{
		"package": func() interface{} {
			return proName
		}(),
		"goPackage": func() interface{} {
			return "./" + proName
		}(),
		"imports": func() interface{} {
			return nil
		}(),
		"messages": func() interface{} {
			type (
				param struct {
					Name string
					Type string
				}

				request struct {
					Name   string
					Params []param
				}

				response struct {
					Name   string
					Params []param
				}

				message struct {
					Request  request
					Response response
				}
			)

			messages := make([]message, 0)
			messages = append(messages, message{
				Request: request{
					Name:   "1",
					Params: []param{{Name: "1", Type: "1"}},
				},
				Response: response{
					Name:   "1",
					Params: []param{{Name: "1", Type: "1"}},
				},
			})

			return messages
		}(),
		"services": func() interface{} {
			return nil
		}(),
	})
	fp.Close()

	return nil
}
