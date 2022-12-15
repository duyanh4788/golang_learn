package uploadprovider

import (
	"context"
	"golang_01/common"
)

type UploadProvider interface {
	SaveFileUpload(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
