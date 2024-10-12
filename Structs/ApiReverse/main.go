package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Object struct {
	Header Header `json:"header"`
	Data   []Data `json:"data"`
}

type Header struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

type Data struct {
	Type string     `json:"type"`
	Id   int        `json:"id"`
	Attr Attributes `json:"attributes"`
}

type Attributes struct {
	Email string `json:"email"`
	A_Ids []int  `json:"article_ids"`
}

func ReadResponse(rawResp string) (Object, error) {
	var resp Object
	err := json.Unmarshal([]byte(rawResp), &resp)
	return resp, err
}

func main() {
	//objects := []Object{}
	rawResp := `{
		"header": {
			"code": 0,
			"message": ""
		},
		"data": [{
			"type": "user",
			"id": 100,
			"attributes": {
				"email": "bob@yandex.ru",
				"article_ids": [10, 11, 12]
			}
		}]
	}`
	ans, err := ReadResponse(rawResp)
	if err != nil {
		log.Fatal()
	}
	fmt.Println(ans, ans.Data[0].Type)
}
