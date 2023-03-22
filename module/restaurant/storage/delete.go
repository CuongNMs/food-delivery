package storagerestaurant

import (
	"context"
	"food-delivery/module/restaurant/model"
)

func (s *sqlStorage) Delete(context context.Context, id int) error {

	//Soft delete
	if err := s.db.Table(restaurantmodel.Restaurant{}.TableName()).
		Where("id=?", id).
		Updates(map[string]interface{}{"status": 0}).Error; err != nil {
		return err
	}
	return nil
}
