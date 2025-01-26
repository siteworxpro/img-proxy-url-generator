package commands

import (
	"encoding/json"
	"fmt"
	"github.com/siteworxpro/img-proxy-url-generator/aws"
	"github.com/siteworxpro/img-proxy-url-generator/config"
	"github.com/siteworxpro/img-proxy-url-generator/generator"
	"github.com/siteworxpro/img-proxy-url-generator/printer"
	"github.com/urfave/cli/v2"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type jsonRequest struct {
	Image  string   `json:"image"`
	Params []string `json:"params"`
	Format string   `json:"format"`
}

func ServerCommand() *cli.Command {
	return &cli.Command{
		Name:  "server",
		Usage: "Start a webserver for s3 file browsing and the web service",
		Action: func(c *cli.Context) error {
			p := printer.NewPrinter()
			return startServer(c, p)
		},
	}
}

func startServer(c *cli.Context, p *printer.Printer) error {
	cfg, err := config.NewConfig(c.String("config"))
	if err != nil {
		return err
	}

	ig, err := generator.NewGenerator(cfg)
	if err != nil {
		return err
	}

	_, err = os.Stat("./templates")
	if !os.IsNotExist(err) {
		awsClient := aws.NewClient(cfg)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(err.Error()))

				return
			}
			contToken := r.URL.Query().Get("next")

			var next *string
			if contToken == "" {
				next = nil
			} else {
				next = &contToken
			}

			contents, err := awsClient.ListBucketContents(next)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte(err.Error()))

				return
			}

			for i, content := range contents.Images {
				contents.Images[i].Url, _ = ig.GenerateUrl("s3://"+cfg.Aws.AwsBucket+"/"+content.Name, []string{"pr:sq"}, "")
				contents.Images[i].Download, _ = ig.GenerateUrl("s3://"+cfg.Aws.AwsBucket+"/"+content.Name, []string{""}, "")
			}

			file, _ := os.ReadFile("./templates/index.gohtml")

			tmpl := template.Must(template.New("index").Parse(string(file)))

			err = tmpl.Execute(w, contents)

			if err != nil {
				println(err.Error())
			}
		})
	}

	http.HandleFunc("/generate", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" {
			w.WriteHeader(404)

			return
		}

		bodyContents := make([]byte, r.ContentLength)
		_, _ = r.Body.Read(bodyContents)

		jr := jsonRequest{}
		err = json.Unmarshal(bodyContents, &jr)
		if err != nil {
			println(err.Error())
			w.WriteHeader(500)
			return
		}

		url, err := signURL(jr.Image, jr.Params, jr.Format)
		if err != nil {
			println(err.Error())
			w.WriteHeader(500)
			return
		}

		log.Println(fmt.Sprintf("%s - [%s] - (%s)", jr.Image, strings.Join(jr.Params, ","), url))

		_, _ = w.Write([]byte(url))
	})

	p.LogSuccess("Starting http server on port 8080. http://localhost:8080")
	log.Fatal(http.ListenAndServe("0.0.0.0:8080", nil))

	return nil
}
