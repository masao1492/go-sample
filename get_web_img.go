package main

import (
	"flag"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// GetImg URLページに含まれているimgのURLを取得しダウンロードします。
func GetImg(url string) {
	doc, _ := goquery.NewDocument(url)
	var i = 1
	doc.Find(".description img").Each(func(_ int, s *goquery.Selection) {
		url, _ := s.Attr("src")
		response, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer response.Body.Close()

		filename := strings.Join([]string{strconv.Itoa(i), ".jpg"}, "")
		file, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		io.Copy(file, response.Body)
		i++
	})
}

func checkDir(outputDir string) {
	if f, err := os.Stat(outputDir); os.IsNotExist(err) || !f.IsDir() {
		os.Mkdir(outputDir, 0777)
	}
	os.Chdir(outputDir)
}

func main() {
	flag.Parse()
	/* flag.Arg(0): 取得したい画像のあるURL,
	   flag.Arg(1): 移動したい、もしくは作りたいディレクトリ名,
	   flag.Arg(2): 取得したい画像のurlが含まれるセレクタ名
	*/
	url := flag.Arg(0)
	if url == "" {
		panic("url is not input")
	}
	outputDir := flag.Arg(1)
	if outputDir == "" {
		outputDir = "."
	}
	checkDir(outputDir)
	// fmt.Println(url)
	GetImg(url)
}
