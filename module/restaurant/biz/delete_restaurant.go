package bizrestaurant

import (
	"context"
	"errors"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

type DeleteRestaurantStorage interface {
	FindDataWithCondition(ctx context.Context,
		condition map[string]interface{},
		moreKeys ...string) (*restaurantmodel.Restaurant, error)
	Delete(context context.Context, id int) error
}

type deleteRestaurantBiz struct {
	storage DeleteRestaurantStorage
}

func NewDeleteRestaurantBiz(storage DeleteRestaurantStorage) *deleteRestaurantBiz {
	return &deleteRestaurantBiz{storage: storage}
}

func (biz *deleteRestaurantBiz) DeleteRestaurant(context context.Context, id int) error {
	oldData, err := biz.storage.FindDataWithCondition(context, map[string]interface{}{"id": id})
	if err != common.ErrRecordNotFound {
		return err
	}
	if oldData.Status == 0 {
		return errors.New("Data has deleted")
	}
	if err := biz.storage.Delete(context, id); err != nil {
		return err
	}
	return nil
}
