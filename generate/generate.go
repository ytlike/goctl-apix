package generate

import (
	"fmt"
	"github.com/urfave/cli/v2"
	swgger "github.com/ytlike/goctl-apix/generate/swagger"
	"github.com/ytlike/goctl-apix/generate/zrpc"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
	"github.com/zeromicro/go-zero/tools/goctl/util/pathx"
	"os"
	"path/filepath"
	"strings"
)

func Do(in *plugin.Plugin, ctx *cli.Context) error {
	err := docGen(in)
	if err != nil {
		return fmt.Errorf("docGen err:%v", err)
	}

	err = zRpcGen(in)
	if err != nil {
		return fmt.Errorf("zRpcGen err:%v", err)
	}

	return nil
}

func docGen(in *plugin.Plugin) error {
	g := swgger.NewGenerator()

	apiDir, apiName := paresApiFilePath(in.ApiFilePath)
	swaggerFile := filepath.Join(apiDir, "doc", fmt.Sprintf("%s.json", apiName))
	outDir := filepath.Join(apiDir, "doc")
	dCtx := &swgger.DocContext{
		Plugin:      in,
		SwaggerFile: swaggerFile,
		ApiName:     apiName,
		Host:        "",
		BasePath:    "",
		Output:      outDir,
	}

	return g.Generate(dCtx)
}

func zRpcGen(in *plugin.Plugin) error {
	g := zrpc.NewGenerator("gozero", true)

	apiDir, apiName := paresApiFilePath(in.ApiFilePath)
	pbOutDir := filepath.Join(apiDir, "pb") //跟api同级目录
	protoFile := filepath.Join(pbOutDir, fmt.Sprintf("%s.proto", apiName))
	zRpcOutDir := filepath.Join(in.Dir, apiName)

	ctx := &zrpc.ZRpcContext{
		Plugin:    in,
		ApiName:   apiName,
		ProtoFile: protoFile,
		ProtocCmd: fmt.Sprintf("protoc -I=%s %s --go_out=%s --go-grpc_out=%s",
			pbOutDir, apiName+".proto", pbOutDir, pbOutDir),
		IsGooglePlugin: true,
		GoOutput:       pbOutDir,
		GrpcOutput:     pbOutDir,
		Output:         zRpcOutDir,
	}

	err := g.Generate(ctx)
	if err != nil {
		return err
	}

	//移动pb生成文件
	srcPb := filepath.Join(pbOutDir, apiName, fmt.Sprintf("%s.pb.go", apiName))
	distPb := filepath.Join(pbOutDir, fmt.Sprintf("%s.pb.go", apiName))
	pathx.Copy(srcPb, distPb)
	srcGrpcPb := filepath.Join(pbOutDir, apiName, fmt.Sprintf("%s_grpc.pb.go", apiName))
	distGrpcPb := filepath.Join(pbOutDir, fmt.Sprintf("%s_grpc.pb.go", apiName))
	pathx.Copy(srcGrpcPb, distGrpcPb)
	os.RemoveAll(filepath.Join(pbOutDir, apiName))

	return nil
}

func paresApiFilePath(apiFilePath string) (string, string) {
	var (
		apiDir  string
		apiName string
	)

	apiFilePath, _ = filepath.Abs(apiFilePath)
	apiDir = filepath.Dir(apiFilePath)
	apiName = strings.Replace(filepath.Base(apiFilePath), ".api", "", 1)
	return apiDir, apiName
}
