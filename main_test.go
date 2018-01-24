package main

import (
"testing"

"github.com/aws/aws-sdk-go/service/s3"
	"github.com/golang/mock/gomock"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/stretchr/testify/mock"
	"github.com/jckuester/go-mock-examples/mocks_mockery"
	"github.com/jckuester/go-mock-examples/mocks_gomock"
)

func TestS3ListBuckets_mockery(t *testing.T) {
	mockS3 := &mocks_mockery.S3API{}

	mockResultFn := func(input *s3.ListBucketsInput) *s3.ListBucketsOutput {
		output := &s3.ListBucketsOutput{}
		output.SetBuckets([]*s3.Bucket{
			{
				Name: aws.String("test-bucket-mockery"),
			},
		})
		return output
	}

	mockS3.On("ListBuckets", mock.MatchedBy(func(input *s3.ListBucketsInput) bool {
		return true
	})).Return(mockResultFn, nil)

	SomeMethod(mockS3)
}

func TestS3ListBuckets_gomock(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockS3 := mock_s3iface.NewMockS3API(mockCtrl)
	mockS3.EXPECT().ListBuckets(&s3.ListBucketsInput{}).Return(&s3.ListBucketsOutput{
		Buckets: []*s3.Bucket{
			{
				Name: aws.String("test-bucket-gomock"),
			},
		},
	}, nil)

	SomeMethod(mockS3)
}
