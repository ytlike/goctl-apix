package generate

import (
	"fmt"
	swgger "github.com/ytlike/goctl-apix/generate/swagger"
	"github.com/ytlike/goctl-apix/generate/zrpc"
	"github.com/zeromicro/go-zero/tools/goctl/plugin"
)

func Do(in *plugin.Plugin, pronName string) error {
	err := docGen(in, pronName)
	if err != nil {
		return fmt.Errorf("docGen err:%v", err)
	}

	err = zRpcGen(in, pronName)
	if err != nil {
		return fmt.Errorf("zRpcGen err:%v", err)
	}

	return nil
}

func docGen(in *plugin.Plugin, pronName string) error {
	g := swgger.NewGenerator()

	dCtx := &swgger.DocContext{
		Plugin:   in,
		Host:     "",
		BasePath: "",
		Output:   in.Dir,
		ProName:  pronName,
	}

	return g.Generate(dCtx)
}

func zRpcGen(in *plugin.Plugin, pronName string) error {
	g := zrpc.NewGenerator("gozero", true)

	pbOutDir := in.Dir + "/" + pronName
	src := pronName + ".proto"
	ctx := &zrpc.ZRpcContext{
		Plugin:  in,
		Src:     src,
		ProName: pronName,
		ProtocCmd: fmt.Sprintf("protoc -I=%s %s --go_out=%s --go_opt=Mbase/common.proto=./base --go-grpc_out=%s",
			in.Dir, src, pbOutDir, pbOutDir),
		IsGooglePlugin: true,
		GoOutput:       pbOutDir,
		GrpcOutput:     pbOutDir,
		Output:         pbOutDir,
	}

	return g.Generate(ctx)
}
