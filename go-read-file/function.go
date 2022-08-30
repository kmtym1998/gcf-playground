package f

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/GoogleCloudPlatform/functions-framework-go/functions"
	"github.com/kmtym1998/gcf-playground/postgres"
)

func init() {
	functions.HTTP("Ping", Ping)
}

type user struct {
	id        int
	name      string
	email     string
	createdAt string
	updatedAt string
}

func Ping(w http.ResponseWriter, r *http.Request) {
	func() {
		uri := "user=kmtym1998 password=kmtym1998 database=app host=10.0.0.2 port=5432"
		pg := postgres.NewPGService(uri)

		if err := pg.Open(); err != nil {
			panic(err)
		}

		if err := pg.DB.Ping(); err != nil {
			panic(err)
		}

		rows, err := pg.DB.Query("SELECT * FROM users;")
		if err != nil {
			panic(err)
		}
		defer rows.Close()

		ul := []user{}
		u := user{}
		for rows.Next() {
			if err := rows.Scan(&u.id, &u.name, &u.email, &u.createdAt, &u.updatedAt); err != nil {
				panic(err)
			}

			ul = append(ul, u)
		}

		log.Printf("%+v", ul)
	}()

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
