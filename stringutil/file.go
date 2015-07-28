package file

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.Open("/Users/june/Code/go/src/github.com/csk6124/juneGoHome/web/server.go")
	if err != nil {
		return 
	}

	defer file.Close()

	// 파일 크기 구함 
	stat, err := file.Stat()
	if err != nil {
		return
	}

	// 파일 읽음
	bs := make([]byte, stat.Size())
	_, err = file.Read(bs)
	if err != nil {
		return
	}

	str := string(bs)
	fmt.Println(str)
}