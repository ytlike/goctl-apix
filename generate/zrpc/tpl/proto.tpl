syntax = "proto3";

package {{.package}};
option go_package="{{.goPackage}}";

{{- range $import := .imports}}
import "$import";
{{- end}}

{{- range $message := .messages}}
message $message.Request.Name {
    {{- range $param := $message.Request.Params}}
    $param.Name
    {{- end}}
}

message $message.Response.Name {
   {{- range $param := $message.Params}}
   {{- end}}
}
{{- end}}

{{- range $service := .services}}
service $service.Name}} {
  {{- range $rpcFunc := $service.RpcFuncList}}
  rpc $rpcFunc.Name ($rpcFunc.Request) returns($rpcFunc.Response);
  {{- end}}
}
{{- end}}