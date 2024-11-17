package main

import (
	"encoding/json"
	"fmt"
	"github.com/bigkevmcd/go-configparser"
	"github.com/siteworxpro/img-proxy-url-generator/aws"
	"github.com/siteworxpro/img-proxy-url-generator/generator"
	proto "github.com/siteworxpro/img-proxy-url-generator/grpc"
	"github.com/siteworxpro/img-proxy-url-generator/printer"
	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"html/template"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

var keyBin, saltBin []byte

var imgGenerator *generator.Generator

var Version = "v0.0.0"

var awsConfig aws.Config

type jsonRequest struct {
	Image  string   `json:"image"`
	Params []string `json:"params"`
	Format string   `json:"format"`
}

func main() {

	pr := printer.NewPrinter()

	var commands []*cli.Command

	commands = append(commands, &cli.Command{
		Name:  "generate",
		Usage: "Generate an image from a URL",
		Action: func(c *cli.Context) error {
			return run(c, pr)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "image",
				Aliases:  []string{"i"},
				Required: true,
			},
			&cli.StringFlag{
				Name:    "format",
				Aliases: []string{"f"},
				Usage:   "Convert the image to the specified format",
			},
			&cli.StringSliceFlag{
				Name:    "params",
				Aliases: []string{"p"},
				Usage:   "Processing options to be passed to the generator ref: https://docs.imgproxy.net/usage/processing",
			},
		},
	})

	commands = append(commands, &cli.Command{
		Name:  "server",
		Usage: "Start a webserver for s3 file browsing and the web service",
		Action: func(c *cli.Context) error {
			return startServer(c, pr)
		},
	})

	commands = append(commands, &cli.Command{
		Name:  "grpc",
		Usage: "Start a grpc service",
		Flags: []cli.Flag{
			&cli.IntFlag{
				Name:     "port",
				Aliases:  []string{"p"},
				Usage:    "Port to listen on",
				Required: false,
				Value:    9000,
			},
		},
		Action: func(c *cli.Context) error {
			err := initGenerator(c.String("config"))

			if err != nil {
				return err
			}

			s := grpc.NewServer()
			addr := fmt.Sprintf(":%d", c.Int("port"))
			println("listening on", addr)
			lis, err := net.Listen("tcp", addr)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			proto.RegisterGeneratorServer(s, proto.NewService(imgGenerator))
			err = s.Serve(lis)
			if err != nil {
				log.Fatalf("failed to serve: %v", err)
			}

			return nil
		},
	})

	commands = append(commands, &cli.Command{
		Name:  "decrypt",
		Usage: "decrypt an image url contents",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:     "url",
				Aliases:  []string{"u"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			err := initGenerator(c.String("config"))
			if err != nil {
				return err
			}

			plain, err := imgGenerator.Decrypt(c.String("url"))
			if err != nil {
				return err
			}

			pr.LogSuccess(plain)

			return nil
		},
	})

	app := &cli.App{
		Name:           "img-proxy-url-generator",
		Usage:          "URL Generator for the img proxy service",
		DefaultCommand: "generate",
		Version:        Version,
		Commands:       commands,
		Action: func(c *cli.Context) error {
			return run(c, pr)
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "config",
				Aliases:     []string{"c"},
				Usage:       "Config file to load from",
				DefaultText: "imgproxy.cfg",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		pr.LogError(err.Error())

		os.Exit(1)
	}
}

func startServer(c *cli.Context, p *printer.Printer) error {
	err := initGenerator(c.String("config"))
	if err != nil {
		return err
	}

	_, err = os.Stat("./templates")
	if !os.IsNotExist(err) {
		awsClient := aws.NewClient(&awsConfig)
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			contToken := r.URL.Query().Get("next")

			var next *string
			if contToken == "" {
				next = nil
			} else {
				next = &contToken
			}

			contents, err := awsClient.ListBucketContents(next)
			if err != nil {
				return
			}

			for i, content := range contents.Images {
				contents.Images[i].Url, _ = signURL("s3://"+awsConfig.Bucket+"/"+content.Name, []string{"pr:sq"}, "")
				contents.Images[i].Download, _ = signURL("s3://"+awsConfig.Bucket+"/"+content.Name, []string{""}, "")
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

func run(c *cli.Context, p *printer.Printer) error {
	err := initGenerator(c.String("config"))

	if err != nil {
		return err
	}

	url, err := signURL(c.String("image"), c.StringSlice("params"), c.String("format"))

	if err != nil {
		return err
	}

	p.LogInfo("Url Generated...")

	println(url)

	return nil
}

func initGenerator(config string) error {
	var err error

	if config == "" {
		config = "imgproxy.cfg"
	}

	p, err := configparser.NewConfigParserFromFile(config)
	if err != nil {
		return err
	}

	if !p.HasSection("img-proxy") {
		return fmt.Errorf("config error - [img-proxy] config required")
	}

	config, err = p.Get("img-proxy", "key")
	if config != "" {
		keyBin = []byte(config)
	}

	config, err = p.Get("img-proxy", "salt")
	saltBin = []byte(config)

	hostConf, err := p.Get("img-proxy", "host")
	if err != nil {
		return err
	}

	plainConfig, err := p.Get("img-proxy", "plain-url")

	encKey, err := p.Get("img-proxy", "encryption-key")

	generatorConfig := generator.Config{
		Salt:          saltBin,
		Key:           keyBin,
		Host:          hostConf,
		EncryptionKey: &encKey,
		PlainUrl:      plainConfig != "",
	}

	imgGenerator, err = generator.NewGenerator(generatorConfig)
	if err != nil {
		return err
	}

	if p.HasSection("aws") {
		awsConfig.AwsSecret, _ = p.Get("aws", "secret")
		awsConfig.AwsKey, _ = p.Get("aws", "key")
		awsConfig.AwsRole, _ = p.Get("aws", "role")
		awsConfig.Bucket, _ = p.Get("aws", "bucket")
	}

	return nil
}

func signURL(file string, params []string, formatS string) (string, error) {
	format, err := imgGenerator.StringToFormat(formatS)
	if err != nil {
		return "", err
	}

	url, err := imgGenerator.GenerateUrl(file, params, format)

	if err != nil {
		return "", err
	}

	return url, nil
}
