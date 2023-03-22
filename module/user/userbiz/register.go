package userbiz

import (
	"context"
	"food-delivery/common"
	"food-delivery/module/user/usermodel"
)

type RegisterStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
	CreateUser(ctx context.Context, data *usermodel.UserCreate) error
}

type Hasher interface {
	Hash(data string) string
}

type registerBusiness struct {
	registerStorage RegisterStorage
	hasher          Hasher
}

func NewRegisterBusiness(storage RegisterStorage, hasher Hasher) *registerBusiness {
	return &registerBusiness{
		registerStorage: storage,
		hasher:          hasher,
	}
}

func (r *registerBusiness) Register(ctx context.Context, data *usermodel.UserCreate) error {
	user, _ := r.registerStorage.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if user != nil {
		return usermodel.ErrEmailExisted
	}
	salt := common.GenSalt(50)
	data.Password = r.hasher.Hash(data.Password + salt)
	data.Salt = salt
	data.Role = "user"
	//data.Status = 1
	if err := r.registerStorage.CreateUser(ctx, data); err != nil {
		return common.ErrCannotCreateEntity(err)
	}
	return nil
}
