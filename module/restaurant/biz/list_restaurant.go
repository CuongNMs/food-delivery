package bizrestaurant

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

// tạo interface store
type ListRestaurantStorage interface {
	ListDataWithCondition(ctx context.Context,
		filter *restaurantmodel.Filter,
		paging *common.Paging,
		moreKeys ...string) ([]restaurantmodel.Restaurant, error)
}

// tạo struct biz có properties là interface storage
type listRestaurantBiz struct {
	storage ListRestaurantStorage
}

// Contructor của biz. Cho phép inject thông qua contructor
func NewListRestaurantBiz(storage ListRestaurantStorage) *listRestaurantBiz {
	return &listRestaurantBiz{storage: storage}
}

// Method của struct biz.
func (biz *listRestaurantBiz) ListRestaurant(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging) ([]restaurantmodel.Restaurant, error) {

	// Logic business here...
	result, err := biz.storage.ListDataWithCondition(ctx, filter, paging)
	if err != nil {
		return nil, err
	}
	return result, nil
}
