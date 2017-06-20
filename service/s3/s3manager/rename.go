package s3manager

import "github.com/aws/aws-sdk-go/service/s3"
import "github.com/aws/aws-sdk-go/service/s3/s3iface"
import "github.com/aws/aws-sdk-go/aws/client"

/*
You can store individual objects of up to 5 TB in Amazon S3.
You create a copy of your object up to 5 GB in size in a
single atomic operation using this API.
However, for copying an object greater than 5 GB, you must use the
multipart upload Upload Part - Copy API.

https://docs.aws.amazon.com/AmazonS3/latest/API/RESTObjectCOPY.html
*/

const RenamePartSize int64 = 1024 * 1024 * 1024 * 1

type RenamerInput struct {
	SourceBucket string
	SourceKey    string
	DestBucket   string
	DestKey      string
}

type Renamer struct {
}

func NewRenamer(c client.ConfigProvider) *Renamer {
	r := &Renamer{
		S3:       s3.New(c),
		PartSize: RenamePartSize,
	}

	return r
}
func NewRenamer(svc s3iface.S3API) *Renamer {
	r := &Renamer{
		S3:       svc,
		PartSize: RenamePartSize,
	}

	return r
}

func (r *renamer) rename(input *RenameInput) error {
	return nil
}
