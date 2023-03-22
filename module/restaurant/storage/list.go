package storagerestaurant

import (
	"context"
	"food-delivery/common"
	restaurantmodel "food-delivery/module/restaurant/model"
)

func (s *sqlStorage) ListDataWithCondition(ctx context.Context,
	filter *restaurantmodel.Filter,
	paging *common.Paging,
	moreKeys ...string) ([]restaurantmodel.Restaurant, error) {
	var result []restaurantmodel.Restaurant
	db := s.db.Table(restaurantmodel.Restaurant{}.TableName()).Where("status in (1)")
	if f := filter; f != nil {
		if f.OwnerId > 0 {
			db = db.Where("owner_id = ?", f.OwnerId)
		}
		if len(f.Status) > 0 {
			db = db.Where("status in (?)", f.Status)
		}
	}
	if err := db.Count(&paging.Total).Error; err != nil {
		return nil, err
	}

	if v := paging.FakeCursor; v != "" {
		uid, err := common.FromBase58(v)
		if err != nil {
			common.ErrInternal(err)
		}
		db = db.Where("id < ?" + uid.GetLocalID())
	} else {
		offset := (paging.Page - 1) * paging.Limit
		db = db.Offset(offset)
	}

	if err := s.db.
		Limit(paging.Limit).
		Order("id desc").
		Find(&result).Error; err != nil {
		return nil, err
	}

	if len(result) > 0 {
		last := result[len(result)-1]
		last.Mask()
		paging.NextCursor = last.FakeId.String()
	}
	return result, nil
}
