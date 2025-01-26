package aws

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type Image struct {
	Url      string
	Download string
	Name     string
	S3Path   string
}

type BucketList struct {
	Images     []Image
	StartAfter string
}

func (s *Service) ListBucketContents(continuationToken *string) (*BucketList, error) {
	v2, err := s.s3.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:            aws.String(s.bucket),
		MaxKeys:           aws.Int64(20),
		ContinuationToken: continuationToken,
	})

	if err != nil {
		return nil, err
	}

	bucketList := BucketList{
		Images: make([]Image, 0),
	}

	if v2.NextContinuationToken != nil {
		bucketList.StartAfter = *v2.NextContinuationToken
	}

	for _, item := range v2.Contents {
		if *item.Size == 0 {
			continue
		}

		image := Image{
			Name:   *item.Key,
			S3Path: "s3://" + s.bucket + "/" + *item.Key,
		}
		bucketList.Images = append(bucketList.Images, image)
	}

	return &bucketList, nil
}
