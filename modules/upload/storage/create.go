package uploadstorage

import (
	"context"
	"golang_01/common"
)

func (store *sqlStore) CreateImage(context context.Context, data *common.Image) error {
	if err := store.db.Table(data.TableName()).Create(data).Error; err != nil {
		common.ErrDB(err)
	}

	return nil
}
