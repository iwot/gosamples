// Json形式のデータをItem型に読み込む。
package main

import (
	"encoding/json"
	"fmt"
)

type Item struct {
	Code  string
	Year  int64 `json:",string"`
	Month int64 `json:"m,string"`
}

func main() {
	j := `{"code":"A-01", "year":"2013", "m":"7"}`
	item := new(Item)
	err := json.Unmarshal([]byte(j), item)
	if err == nil {
		fmt.Printf("%+v\n", item)
	} else {
		fmt.Println(err)
	}
}
