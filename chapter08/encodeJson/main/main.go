package main

import (
	"encoding/json"
	"fmt"
)

//编码json
func main() {
	mapToJson()
}

func mapToJson() {
	c := make(map[string]interface{})
	c["name"] = "GG"
	c["title"] = "TT"
	c["contact"] = map[string]interface{}{
		"home": "2222",
		"cell": "3333",
	}
	//编码JSON,带格式化的,指定前缀和缩进
	data, err := json.MarshalIndent(c, "+", "--")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(data)
	fmt.Println(string(data))
	//没有格式化的
	data2, err2 := json.Marshal(c)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	fmt.Println(data2)
	fmt.Println(string(data2))
}
