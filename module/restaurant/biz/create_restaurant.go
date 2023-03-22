package bizrestaurant

import (
	"context"
	"food-delivery/common"
	"food-delivery/module/restaurant/model"
)

// tạo interface store
type CreateRestaurantStorage interface {
	CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error // được implement tại storage
}

// tạo struct biz có properties là interface storage
type createRestaurantBiz struct {
	storage CreateRestaurantStorage
}

// Contructor của biz. Cho phép inject thông qua contructor
func NewCreateRestaurantBiz(storage CreateRestaurantStorage) *createRestaurantBiz {
	return &createRestaurantBiz{storage: storage}
}

// Method của struct biz.
func (biz *createRestaurantBiz) CreateRestaurant(context context.Context, data *restaurantmodel.RestaurantCreate) error {

	// Logic business here...
	if err := data.Validate(); err != nil {
		return err
	}
	if err := biz.storage.CreateRestaurant(context, data); err != nil {
		return common.ErrCannotCreateEntity(err)
	}
	return nil
}
