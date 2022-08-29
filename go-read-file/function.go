package f

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("ListFiles", ListFiles)
}

func ListFiles(w http.ResponseWriter, r *http.Request) {
	// IP 見る
	func() {
		url := "https://rakko.tools/tools/2/"
		byteResp, httpResp, err := Request(http.MethodGet, url, nil, nil)
		if err != nil {
			log.Fatalln("err httpwrap.Request", err)
		}

		htmlResp := string(byteResp)
		startTag := `<textarea id="resultTextArea">`
		endTag := `</textarea>`
		startTagPosition := strings.Index(htmlResp, startTag)
		endTagPosition := strings.LastIndex(htmlResp, endTag) // </textarea> が 2つあるのでその後の方

		println(htmlResp[startTagPosition+len(startTag) : endTagPosition])

		println("httpResp.Status", httpResp.Status)
	}()

	// 静的ファイルを読み取りたい
	func() {
		b, err := os.ReadFile("static/hoge.json")
		if err != nil {
			log.Println(err)
		}
		log.Println("static json", string(b))

		b, err = os.ReadFile(os.Getenv("SOURCE_DIR") + "static/hoge.json")
		if err != nil {
			log.Println(err)
		}

		log.Println("static json", string(b))
	}()

	// カレントディレクトリ
	func() {
		currentDir, _ := os.Getwd()
		log.Println("currentDir", currentDir)
	}()

	// /workspace 配下のファイルを読む
	func() {
		err := filepath.Walk("/workspace", func(path string, _ os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			fmt.Printf("path: %#v\n", path)
			return nil
		})
		if err != nil {
			panic(err)
		}
	}()

	// /workspace/main.go 配下のファイルを読む
	func() {
		b, err := os.ReadFile("/workspace/main.go")
		if err != nil {
			panic(err)
		}
		log.Println("== main.go ==========")
		log.Println(string(b))
		log.Println("============")

		b, err = os.ReadFile("/workspace/go.mod")
		if err != nil {
			panic(err)
		}
		log.Println("== go.mod ==========")
		log.Println(string(b))
		log.Println("============")

		dirs, err := os.ReadDir("/")
		if err != nil {
			panic(err)
		}

		for _, dir := range dirs {
			log.Println("dir.Name()", dir.Name())
		}
		logEntry, err := json.Marshal(map[string]string{
			"severity": "DEBUG",
			"message":  string(b),
		})
		if err != nil {
			panic(err)
		}

		fmt.Print(string(logEntry))
	}()
}

type reqHeaders struct {
	Key   string
	Value string
}

func Request(method string, url string, body *[]byte, headers []reqHeaders) ([]byte, *http.Response, error) {
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, nil, err
	}

	for _, rh := range headers {
		req.Header.Set(rh.Key, rh.Value)
	}

	client := new(http.Client)

	resp, err := client.Do(req)
	if err != nil {
		return nil, resp, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Printf("err resp.Body.Close(). err: %+v, resp: %+v", err, resp)
		}
	}()

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, resp, err
	}

	return b, resp, err
}
