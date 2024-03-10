package main

import (
	"encoding/json"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/add", func(writer http.ResponseWriter, request *http.Request) {
		// 解析参数
		_ = request.ParseForm()
		a, _ := strconv.Atoi(request.Form["a"][0])
		b, _ := strconv.Atoi(request.Form["b"][0])
		writer.Header().Set("Content-Type", "application/json")
		data, _ := json.Marshal(map[string]int{
			"data": a + b,
		})
		_, _ = writer.Write(data)
	})
	_ = http.ListenAndServe(":8088", nil)
}
