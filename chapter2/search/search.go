package search

import (
	"encoding/json"
	"fmt"
	"os"
)

const dataFile = "data/data.json"

//Feed 种子结构体 名称、URI和类型
type Feed struct {
	Name string
	URI  string
	Type string
}

// Run test
func Run() {
	fmt.Println("search")
	feeds, _ := RetrieveFeeds()
	for index, feed := range feeds {
		fmt.Println(index, feed.Name, feed.URI, feed.Type)
	}
}

//RetrieveFeeds 获取从文件中获取种子
func RetrieveFeeds() ([]*Feed, error) {
	//读取data.joson文件，组成Feed结构体
	file, err := os.Open(dataFile)
	//
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var feeds []*Feed
	err = json.NewDecoder(file).Decode(&feeds)
	return feeds, err
}
