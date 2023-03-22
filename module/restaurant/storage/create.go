package storagerestaurant

import (
	"context"
	"food-delivery/module/restaurant/model"
)

// implement interface của biz
func (s *sqlStorage) CreateRestaurant(ctx context.Context, data *restaurantmodel.RestaurantCreate) error {
	if err := s.db.Create(&data).Error; err != nil {
		return err
	}
	return nil
}
