package report

import (
	"context"
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"github.com/charmbracelet/lipgloss/table"
	"github.com/siteworxpro/img-proxy-url-generator/aws"
	"github.com/siteworxpro/img-proxy-url-generator/config"
	"github.com/siteworxpro/img-proxy-url-generator/generator"
	"github.com/siteworxpro/img-proxy-url-generator/redis"
	"sort"
	"strconv"
	"time"
)

const lastAccessKey = "imgproxy:%s:last_access"
const requestsKey = "imgproxy:%s:requests"

var rowStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF"))
var headerStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true)

func Handle(config *config.Config) error {
	a := aws.NewClient(config)
	r, err := redis.New(config)
	ig, err := generator.NewGenerator(config)
	if err != nil {
		return err
	}

	var continuationToken *string
	list, err := a.ListBucketContents(continuationToken)
	if err != nil {
		return err
	}

	var rows [][]string
	for list.StartAfter != "" {
		for _, image := range list.Images {
			dlUrl, err := ig.GenerateUrl(image.S3Path, []string{}, generator.DEF)
			if err != nil {
				return err
			}
			lastAccessedS, err := r.GetClient().Get(context.Background(), fmt.Sprintf(lastAccessKey, image.S3Path)).Result()
			lastAccessedI, _ := strconv.ParseInt(lastAccessedS, 10, 64)
			lastAccessed := time.Unix(lastAccessedI, 0)
			requestsCount, err := r.GetClient().Get(context.Background(), fmt.Sprintf(requestsKey, image.S3Path)).Result()

			rows = append(rows, []string{image.S3Path, requestsCount, lastAccessed.Format(time.DateTime), dlUrl})
		}

		continuationToken = &list.StartAfter
		list, err = a.ListBucketContents(continuationToken)
	}

	// sort by last accessed
	sort.Slice(rows, func(i, j int) bool {
		return rows[i][2] > rows[j][2]
	})

	t := table.New().StyleFunc(func(row int, col int) lipgloss.Style {
		switch {
		case row == 0:
			return headerStyle
		default:
			return rowStyle
		}
	}).
		Headers("Image", "Times Accessed", "Last Accessed", "URL").
		Rows(rows...)

	fmt.Println(t)

	return nil
}
