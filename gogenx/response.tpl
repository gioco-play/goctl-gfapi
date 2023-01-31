package respx

import (
    "net/http"

    "github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
    Code string         `json:"code"`
    Msg  string      `json:"msg"`
    Data interface{} `json:"data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
    var body Body
    if err != nil {
        body.Code = "9999"
        body.Msg = err.Error()
    } else {
        body.Code = "0"
        body.Msg = "OK"
        body.Data = resp
    }
    httpx.OkJson(w, body)
}