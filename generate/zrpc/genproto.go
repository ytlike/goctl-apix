package zrpc

import (
	_ "embed"
	"fmt"
	"github.com/zeromicro/go-zero/tools/goctl/api/spec"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"os"
	"path/filepath"
	"text/template"
)

//go:embed tpl/proto.tpl
var protoTemplate string

type (
	// message struct
	field struct {
		Type    string
		Name    string
		Index   int
		Comment string
	}

	message struct {
		Name     string
		Fields   []field
		Comments []string
	}

	// service struct
	rpcCall struct {
		Name     string
		Request  string
		Response string
		Comments []string
	}

	service struct {
		Name     string
		RpcCalls []rpcCall
		Comment  string
	}
)

func (g *Generator) GenProto(p *plugin.Plugin, protoFile, apiName string) error {
	abs := filepath.Dir(protoFile)
	pathx.MkdirIfNotExist(abs)

	protoTpl := template.Must(template.New("protoTpl").Parse(protoTemplate))
	fp, err := os.OpenFile(protoFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return fmt.Errorf("GenProto not found proto file err:%v", err)
	}

	api := p.Api
	protoTpl.Execute(fp, map[string]interface{}{
		"package": func() interface{} {
			return "pb"
		}(),
		"goPackage": func() interface{} {
			return filepath.Join(abs, "pb")
		}(),
		"imports": func() interface{} {
			return nil // todo
		}(),
		"messages": func() interface{} {
			messages := make([]message, 0)
			for _, apiType := range api.Types {
				m := message{
					Name:     apiType.Name(),
					Comments: apiType.Documents(),
				}
				ds, ok := apiType.(spec.DefineStruct)
				if ok {
					m.Name = ds.Name()
					for i, member := range ds.Members {
						f := field{
							Name:    member.Name,
							Index:   i + 1,
							Comment: member.Comment,
						}

						switch mt := member.Type.(type) {
						case spec.DefineStruct, spec.PrimitiveType:
							f.Type = toPbType(mt.Name())
						case spec.MapType:
							f.Type = fmt.Sprintf("map<%s,%s>", toPbType(mt.Key), toPbType(mt.Value.Name()))
						case spec.ArrayType:
							f.Type = fmt.Sprintf("repeated %s", toPbType(mt.Value.Name()))
						case spec.InterfaceType:
							panic(fmt.Sprintf("interface not deal %+v", mt))
						case spec.PointerType:
							f.Type = fmt.Sprintf("optional %s", toPbType(mt.Type.Name()))
						}

						m.Fields = append(m.Fields, f)
					}
				}
				messages = append(messages, m)
			}
			return messages
		}(),
		"services": func() interface{} {
			as := api.Service
			services := make([]service, 0)
			s := service{
				Name: as.Name,
			}

			for _, route := range as.Routes() {
				rpc := rpcCall{
					Name: route.Handler,
					Request: func() string {
						if route.RequestType != nil {
							return route.RequestType.Name()
						}

						return "Empty"
					}(),
					Response: func() string {
						if route.ResponseType != nil {
							return route.ResponseType.Name()
						}

						return "Empty"
					}(),
					Comments: route.HandlerDoc,
				}
				s.RpcCalls = append(s.RpcCalls, rpc)
			}

			services = append(services, s)
			return services
		}(),
	})
	fp.Close()

	return nil
}

func toPbType(at string) (pt string) {
	switch at {
	case "int":
		pt = "int32"
	default:
		pt = at
	}
	return
}
