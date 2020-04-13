package main

import (
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

//context (TimeErrorの設定)
func longProcess(ctx context.Context, ch chan string) {
	fmt.Println("run")
	time.Sleep(1 * time.Second)
	fmt.Println("fin")
	ch <- "result"
}

//json
//`json:"name"`送信時のencode名の指定
//omitempty -> からの場合は省略
type Person struct {
	Name      string   `json:"name",omitempty`
	Age       int      `json:"age"`
	Nicknames []string `json:"nicknames"`
}

//MarshalJSON()とすることで、json.Marshalを呼んだ際にこっちが呼ばれる
func (p Person) MarshalJSON() ([]byte, error) {
	v, err := json.Marshal(&struct {
		Name string
	}{
		Name: "Mr." + p.Name,
	})
	return v, err
}

//UnmarshalJSON()
func (p *Person) UnmarshalJSON(b []byte) error {
	type Person2 struct {
		Name      string
		Age       int
		Nicknames []string
	}
	var p2 Person2
	err := json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Println(err)
	}
	p.Name = p2.Name + "!"
	p.Age = p2.Age
	p.Nicknames = p2.Nicknames
	return err
}

func main() {
	//context (TimeErrorの設定)
	ch := make(chan string)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()
	go longProcess(ctx, ch)
	// CTXLOOP:
	// 	for {
	// 		select {
	// 		case <-ctx.Done():
	// 			fmt.Println("Error")
	// 			break CTXLOOP
	// 		case <-ch:
	// 			fmt.Println("success")
	// 			break CTXLOOP
	// 		}
	// 	}
	// 	fmt.Println("########come########")

	//ioutil
	//読み込み
	content, err := ioutil.ReadFile("lib-prac.go")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(string(content))
	//書き込み（ファイル作成）
	// if err := ioutil.WriteFile("ioutil_tmp.go", content, 0666); err != nil {
	// 	log.Fatalln(err)
	// }
	//読み込み（byte）
	r := bytes.NewBuffer([]byte("abc"))
	contents, _ := ioutil.ReadAll(r)
	fmt.Println(contents, string(contents))

	//http
	base, _ := url.Parse("http://example.com/")
	reference, _ := url.Parse("/test?a=1&b=2")
	endpoint := base.ResolveReference(reference).String()
	fmt.Println("endpoint : ", endpoint)
	req, _ := http.NewRequest("GET", endpoint, nil)
	req.Header.Add("If-None-Match", `W/"test"`)
	q := req.URL.Query()
	q.Add("c", "3&%")
	fmt.Println("q : ", q)
	fmt.Println("q.Encode() : ", q.Encode())
	req.URL.RawQuery = q.Encode()

	var client *http.Client = &http.Client{}
	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	//json
	b := []byte(`{"name":"mike","age":20,"nicknames":["me","you","her"]}`)
	var p Person
	//jsonデータの取得
	if err := json.Unmarshal(b, &p); err != nil {
		fmt.Println(err)
	}
	fmt.Println(p.Name, p.Age, p.Nicknames)
	//データをJsonに変換
	v, _ := json.Marshal(p)
	fmt.Println(string(v))

	//API
	const apiKey = "User1Key"
	const apiSecret = "User1Secret"
	data := ([]byte("data"))
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	sign := hex.EncodeToString(h.Sum(nil))
	fmt.Println("---------API---------")
	fmt.Println(sign)

	Server(apiKey, sign, data)
	fmt.Println("---------------------")
}

//API
var DB = map[string]string{
	"User1Key": "User1Secret",
	"User2Key": "User2Secret",
}
//サーバ側と、クライアント側で一致するか確認を行う
func Server(apiKey, sign string, data []byte) {
	apiSecret := DB[apiKey]
	h := hmac.New(sha256.New, []byte(apiSecret))
	h.Write(data)
	expectedHMAC := hex.EncodeToString(h.Sum(nil))
	fmt.Println(sign == expectedHMAC)
}