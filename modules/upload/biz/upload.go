package uploadbiz

import (
	"bytes"
	"context"
	"fmt"
	"golang_01/common"
	"golang_01/component/uploadprovider"
	"image"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"log"
	"path/filepath"
	"strings"
	"time"
)

type CreateImagesStore interface {
	CreateImage(context context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider uploadprovider.UploadProvide
	imgStore CreateImagesStore
}

func NewUploadBiz(provider uploadprovider.UploadProvide, imgStore CreateImagesStore) *uploadBiz {
	return &uploadBiz{provider: provider, imgStore: imgStore}
}

func (biz *uploadBiz) Upload(ctx context.Context, data []byte, folder, fileName string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)

	w, h, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, common.ErrInvalidRequest(err)
	}

	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}

	fileExt := filepath.Ext(fileName)
	fileName = fmt.Sprintf("%d%s", time.Now().Nanosecond(), fileExt)

	img, err := biz.provider.SaveFileUpload(ctx, data, fmt.Sprintf("%s%s", folder, fileName))

	if err != nil {
		return nil, common.ErrIntenval(err)
	}

	img.Width = w
	img.Height = h
	img.CloudName = "s3"
	img.Extension = fileExt
	return img, nil
}

func getImageDimension(reader io.Reader) (int, int, error) {
	img, _, err := image.DecodeConfig(reader)

	if err != nil {
		log.Println("err:", err)
		return 0, 0, err
	}

	return img.Width, img.Height, nil
}
