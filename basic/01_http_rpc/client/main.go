package main

import (
	"encoding/json"
	"fmt"
	"github.com/kirinlabs/HttpRequest"
)

func Add(a int, b int) int {
	req := HttpRequest.NewRequest()
	res, _ := req.Get(fmt.Sprintf("localhost:8088/add?a=%d&b=%d", a, b))
	body, _ := res.Body()
	type Res struct {
		Data int `json:"data"`
	}
	var resData Res
	_ = json.Unmarshal(body, &resData)
	return resData.Data
}

func main() {
	fmt.Println(Add(1, 3))
	return
}
