package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

type ResponseData struct {
	Data int `json:"data"`
}

func Add(a int, b int) int {
	req := HttpRequest.NewRequest()
	res, _ := req.Get(fmt.Sprintf(`http://localhost:8088/add?a=%d&b=%d`, a, b))
	body, _ := res.Body()
	rspData := ResponseData{}
	_ = json.Unmarshal(body, &rspData)
	return rspData.Data
}

func main() {
	res := Add(1, 2)
	fmt.Println(res)
}
