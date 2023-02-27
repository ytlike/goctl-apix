syntax = "proto3";

package {{.package}};
option go_package="{{.goPackage}}";

{{- range $import := .imports}}
import "$import";
{{end}}

message Empty{}

{{- range $message := .messages}}
{{- range $comment := $message.Comments}}
{{$comment}}
{{- end}}
message {{$message.Name}} {
    {{- range $field := $message.Fields}}
    {{$field.Type}} {{$field.Name}} = {{$field.Index}}; {{$field.Comment}}
    {{- end}}
}
{{end}}

{{- range $service := .services}}
{{$service.Comment}}
service {{$service.Name}} {
  {{- range $rpcCall := $service.RpcCalls}}
  {{- range $comment := $rpcCall.Comments}}
  {{$comment}}
  {{- end}}
  rpc {{$rpcCall.Name}}({{$rpcCall.Request}}) returns({{$rpcCall.Response}});
  {{- end}}
}
{{end}}