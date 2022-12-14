package uploadprovider

import (
	"context"
	"golang_01/common"
)

type UploadProvide interface {
	SaveFileUpload(ctx context.Context, data []byte, dst string) (*common.Image, error)
}
