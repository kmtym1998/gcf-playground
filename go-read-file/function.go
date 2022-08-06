package f

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func ListFiles(w http.ResponseWriter, r *http.Request) {
	func() {
		currentDir, _ := os.Getwd()
		log.Println("currentDir", currentDir)
	}()

	func() {
		err := filepath.Walk("/workspace", func(path string, info os.FileInfo, err error) error {
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

	func() {
		b, err := os.ReadFile("/workspace/main.go")
		if err != nil {
			panic(err)
		}

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
