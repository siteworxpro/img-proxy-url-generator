package grpc

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/siteworxpro/img-proxy-url-generator/generator"
	"log"
	"strings"
)

type GeneratorService struct {
	UnimplementedGeneratorServer
	imgGenerator *generator.Generator
}

func NewService(imgGenerator *generator.Generator) *GeneratorService {
	return &GeneratorService{imgGenerator: imgGenerator}
}

func (s *GeneratorService) Generate(c context.Context, r *UrlRequest) (*UrlResponse, error) {

	var err error
	format := generator.DEF

	if r.Format != nil {
		format, err = s.imgGenerator.StringToFormat(r.Format.String())
		if err != nil {
			return nil, err
		}
	}

	url, err := s.imgGenerator.GenerateUrl(*r.Image, r.Params, format)
	if err != nil {
		return nil, err
	}

	log.Println(fmt.Sprintf("%s - [%s] - (%s)", *r.Image, strings.Join(r.Params, ","), url))

	return &UrlResponse{Url: aws.String(url)}, nil
}
