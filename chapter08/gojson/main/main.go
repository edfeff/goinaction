package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

//json格式
type Contact struct {
	Name    string
	Title   string
	Contact struct {
		Home string
		Cell string
	}
}

var str = `{
	"name":"wpp",
	"title":"ttt",
	"contact":{
		"home":"4.55.555",
		"cell":"333-444"
	}
}`

func strToJson() {
	var c Contact
	//编码成json对象
	err := json.Unmarshal([]byte(str), &c)
	if err != nil {
		fmt.Println("error", err)
		os.Exit(1)
	}
	fmt.Println(c)
	fmt.Println(c.Name)

}

type (
	gResult struct {
		Gs    string `json:"Gs"`
		Un    string `json:"un"`
		URL   string `json:"url"`
		Title string `json:"title"`
	}
	gResponse struct {
		ResponseData struct {
			Results []gResult `json:"results"`
		} `json:"responseData"`
	}
)

//
type jsonReader struct {
	name string
}

func (jr *jsonReader) Read(p []byte) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

func jsonDecode() {
	var gr gResponse

	jr := jsonReader{"wpp"}
	err := json.NewDecoder(&jr).Decode(&gr)
	if err != nil {
		log.Println("err", err)
	}
	fmt.Println(gr)
}
func main() {
	//strToJson()
	//jsonDecode();
	strToJsonMap()
}
func strToJsonMap() {
	//json to map
	//键 string 值 任意类型
	var c map[string]interface{}
	err := json.Unmarshal([]byte(str), &c)
	if err != nil {
		fmt.Println("err", err)
		return
	}
	fmt.Println(c)
	fmt.Println(c["name"])
	fmt.Println(c["title"])
	fmt.Println(c["contact"])
	//内部的json格式无法转换,必须强制类型转换
	home := c["contact"].(map[string]interface{})["home"]
	cell := c["contact"].(map[string]interface{})["cell"]
	fmt.Println(home)
	fmt.Println(cell)
}
