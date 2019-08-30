package fetcher

import (
	"bufio"
	"bytes"
	"io"
	"log"
	"os"
	"strings"
)

// 本地文件抓取
var excludes = [...]string{".cpp", ".txt"}

func FetchLocal(url string) ([]byte, error) {
	if !strings.HasSuffix(url, ".h") {
		return nil, nil
	}
	for _, m := range excludes {
		if strings.Contains(url, m) {
			return nil, nil
		}
	}
	//读文件
	//log.Printf("=============read %s",url)
	file, err := os.OpenFile(url, os.O_RDONLY, 0666)
	defer file.Close()
	if err != nil {
		if os.IsPermission(err) {
			log.Printf("error no permission")
		} else {
			log.Printf("error open file %v", err)
		}
		return nil, err
	}
	reader := bufio.NewReader(file)
	var buffer bytes.Buffer
	for {
		buf, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Printf("error read file %s ,%v", url, err)
		}
		//log.Printf(" buf = %s\n", string(buf))
		buffer.Write(buf)
	}
	return buffer.Bytes(), nil
}
