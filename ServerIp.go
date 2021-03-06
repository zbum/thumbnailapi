package thumbnailapi

import (
	"encoding/json"
	"net/http"
	"os"
)

type Header struct {
	ResultCode int`json:"resultCode"`
	ResultMessage string`json:"resultMessage"`
	IsSuccessful bool`json:"isSuccessful"`
}

type Content struct {
	Ip string`json:"ip"`
}

type Result struct {
	Content Content `json:"content"`
}

type Response struct {
	Header Header `json:"header"`
	Result Result `json:"result"`
}

func ServerIp(res http.ResponseWriter, req *http.Request)string {
	data := Response{
		Header: Header{ResultCode: 0, ResultMessage: "", IsSuccessful: true},
		Result: Result{Content: Content{Ip: os.Getenv("HOST_IP")}},
	}
	doc, _ := json.Marshal(data)
	res.Header().Set("Content-Type", "application/json")
	return string(doc)
}
