package fetcher

import (
	"bufio"
	"fmt"
	"golang.org/x/net/html/charset"
	"golang.org/x/text/encoding"
	"golang.org/x/text/transform"
	"io/ioutil"
	"log"
	"net/http"
)

func Fetch(url string) ([]byte,error){
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalln(err)
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/73.0.3683.86 Safari/537.36")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	//以下代码单纯请求容易403
	//resp, err := http.Get(url)
	//if err != nil {
	//	panic(err)
	//}
	//defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Println("err status code ", resp.StatusCode)
		return nil,fmt.Errorf("err status code %d",resp.StatusCode)
	}

	//这里可能出现中文乱码，放到其他语言类似，需要做自动编码检测
	bodyReader := bufio.NewReader(resp.Body)
	encoding := determineEncoding(bodyReader)
	//fmt.Println("determine encoding ", encoding)
	utf8Reader := transform.NewReader(bodyReader, encoding.NewDecoder())
	contents, err := ioutil.ReadAll(utf8Reader)
	if err != nil {
		panic(err)
	}
	return contents,nil
}

func determineEncoding(r *bufio.Reader) encoding.Encoding {
	content, err := r.Peek(1024)
	if err != nil {
		panic(err)
	}
	e, _, _ := charset.DetermineEncoding(content, "")
	return e
}