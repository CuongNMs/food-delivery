package uploadbusiness

import (
	"bytes"
	"context"
	"food-delivery/common"
	"food-delivery/component/uploadprovider"
	"food-delivery/module/upload/uploadmodel"
	"strings"
)

type CreateImageStorage interface {
	CreateImage(ctx context.Context, data *common.Image) error
}

type uploadBiz struct {
	provider   uploadprovider.UploadProvider
	imgStorage CreateImageStorage
}

func NewUploadBiz(provider uploadprovider.UploadProvider, imgStorage CreateImageStorage) *uploadBiz {
	return &uploadBiz{provider: provider, imgStorage: imgStorage}
}

func (b *uploadBiz) Upload(ctx context.Context, data []byte, folder, filename string) (*common.Image, error) {
	fileBytes := bytes.NewBuffer(data)
	_, _, err := getImageDimension(fileBytes)

	if err != nil {
		return nil, uploadmodel.ErrFileIsNotImage
	}
	if strings.TrimSpace(folder) == "" {
		folder = "img"
	}
	return nil, nil
}

func getImageDimension(fileBytes *bytes.Buffer) (interface{}, interface{}, interface{}) {
	return 1, 1, nil
}
