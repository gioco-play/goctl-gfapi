package {{.PkgName}}

import (
	"net/http"
	{{.ImportPackages}}

	{{if .HasRequest}}"github.com/zeromicro/go-zero/rest/httpx"{{end}}
)

func {{.HandlerName}}(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		{{if .HasRequest}}var req types.{{.RequestType}}
		if err := httpx.Parse(r, &req); err != nil {
            respx.Response(w, nil, err)
            return
        }

        if err := util.MyValidator.Struct(req); err != nil {
            respx.Response(w, nil, err)
            return
        }{{end}}

		l := {{.LogicName}}.New{{.LogicType}}(r.Context(), svcCtx)
		{{if .HasResp}}resp, {{end}}err := l.{{.Call}}({{if .HasRequest}}&req{{end}})
        {{if .HasResp}}respx.Response(w, resp, err){{else}}respx.Response(w, nil, err){{end}}
	}
}
