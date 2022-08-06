package f

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func ListFiles(w http.ResponseWriter, r *http.Request) {
	log.Println(time.Now())

	currentDir, _ := os.Getwd()
	log.Println("currentDir", currentDir)

	err := filepath.Walk("/", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 特定のディレクトリを無視したい場合は `filepath.SkipDir` を返す
		// 例えば `AAA` という名前のディレクトリを無視する場合は以下のようにする
		// if info.IsDir() && info.Name() == "AAA" {
		// 	return filepath.SkipDir
		// }

		if info.IsDir() && info.Name() == "etc" {
			return filepath.SkipDir
		}

		fmt.Printf("path: %#v\n", path)
		return nil
	})

	if err != nil {
		panic(err)
	}
}
