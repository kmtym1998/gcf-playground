package f

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
)

func init() {
	functions.HTTP("ListFiles", ListFiles)
}

func ListFiles(w http.ResponseWriter, r *http.Request) {
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
