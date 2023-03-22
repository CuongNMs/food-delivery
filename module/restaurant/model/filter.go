package restaurantmodel

type Filter struct {
	OwnerId int    `json:"owner_id,omitempty" form:"owner_id"`
	Status  string `json:"status,omitempty" form:"status"`
}
