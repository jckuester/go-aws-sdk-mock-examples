package main

import (
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/aws/session"
	"fmt"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
)

func main() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	SomeMethod(s3.New(sess))
}

func SomeMethod(client s3iface.S3API) {
	resp, err := client.ListBuckets(&s3.ListBucketsInput{})

	if err == nil {
		for _, b :=  range resp.Buckets {
			fmt.Printf("Bucket: %s\n", *b.Name)
		}
	} else {
		fmt.Println("No buckets found")
	}
}
