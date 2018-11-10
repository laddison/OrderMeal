package model

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
)

type SubData struct {
	Utf8  string
	Token string
}

func (this *SubData) GetDoc() string {
	doc, err := goquery.NewDocument("https://jinshuju.net/f/5WhTlk")
	if err != nil {
		log.Fatal(err)
	}

	var token string

	doc.Find("#new_entry").Each(func(i int, s *goquery.Selection) {
		//utf8, _ = s.Find("input[name='utf8']").Attr("value")
		token, _ = s.Find("input[name='authenticity_token']").Attr("value")
	})

	return token
}

func (this *SubData) HttpPost() {
	cookieJar, _ := cookiejar.New(nil)
	client := &http.Client{
		Jar: cookieJar,
	}

	token := this.GetDoc()

	v := url.Values{}
	v.Set("utf8", "✓")
	v.Set("authenticity_token", token)
	v.Set("entry[field_1]", "EIz0")
	v.Set("entry[field_2]", "7")
	v.Set("embedded", "")
	v.Set("banner", "")
	v.Set("header", "")
	v.Set("background", "")

	body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码

	reqest, err := http.NewRequest("POST", "https://jinshuju.net/f/5WhTlk", body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
	reqest.Header.Set("Referer", "https://jinshuju.net/f/5WhTlk")

	resp, err := client.Do(reqest) //发送请求
	defer resp.Body.Close()        //一定要关闭resp.Body
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
	}

	fmt.Println(string(content))
}
