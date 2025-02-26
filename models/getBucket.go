package models

import (
	"encoding/xml"
	"time"
)

//GetBucketResponse get bucket response
type GetBucketResponse struct {
	XMLName     xml.Name `xml:"ListBucketResult"`
	Name        string
	Prefix      string
	Marker      string
	MaxKeys     int
	IsTruncated bool
	Contents    []GetBucketResponseContent
	NextMarker  string
}

//GetBucketResponseContent Get Bucket Response Content
type GetBucketResponseContent struct {
	Key          string
	LastModified time.Time
	Tag          string `xml:"ETag"`
	Size         int64
	StorageClass string
	Owner        Owner
}
