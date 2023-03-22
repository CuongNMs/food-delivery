package restaurantmodel

import (
	"food-delivery/common"
	"strings"
)

type RestaurantType string

const EntityName = "restaurant"
const TypeNormal RestaurantType = "normal"
const TypePremium RestaurantType = "premium"

type Restaurant struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name;"`
	Addr            string         `json:"addr" gorm:"column:addr;"`
	Type            RestaurantType `json:"type" gorm:"column:type;"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (r *Restaurant) Mask() {
	r.GenUID(common.DbTypeRestaurant)
}

type RestaurantUpdate struct {
	Name  *string        `json:"name" gorm:"column:name"`
	Addr  *string        `json:"addr" gorm:"column:addr"`
	Logo  *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover *common.Images `json:"cover" gorm:"column:cover;"`
}

func (Restaurant) TableName() string {
	return "restaurants"
}

type RestaurantCreate struct {
	common.SQLModel `json:",inline"`
	Name            string         `json:"name" gorm:"column:name"`
	Addr            string         `json:"addr" gorm:"column:addr"`
	Logo            *common.Image  `json:"logo" gorm:"column:logo;"`
	Cover           *common.Images `json:"cover" gorm:"column:cover;"`
}

func (r *RestaurantCreate) Mask() {
	r.GenUID(common.DbTypeRestaurant)
}

func (data *RestaurantCreate) Validate() error {
	data.Name = strings.TrimSpace(data.Name)
	if data.Name == "" {
		return common.ErrNameIsEmpty
	}
	return nil
}

func (RestaurantCreate) TableName() string {
	return Restaurant{}.TableName()
}

func (RestaurantUpdate) TableName() string {
	return Restaurant{}.TableName()
}
