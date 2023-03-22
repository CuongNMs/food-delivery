package userbiz

import (
	"context"
	"food-delivery/component/tokenprovider"
	"food-delivery/module/user/usermodel"
	"time"
)

type LoginStorage interface {
	FindUser(ctx context.Context, conditions map[string]interface{}, moreInfo ...string) (*usermodel.User, error)
}
type loginBusiness struct {
	//appCtx        appctx.AppContext
	storageUser   LoginStorage
	tokenProvider tokenprovider.Token
	hasher        Hasher
	expiry        int
}

func NewLoginBusiness(storageUser LoginStorage,
	tokenProvider tokenprovider.Token,
	hasher Hasher,
	expiry int) *loginBusiness {
	return &loginBusiness{
		storageUser:   storageUser,
		tokenProvider: tokenProvider,
		hasher:        hasher,
		expiry:        expiry,
	}
}

func (business *loginBusiness) Login(ctx context.Context, data *usermodel.UserLogin) (*tokenprovider.Token, error) {
	user, err := business.storageUser.FindUser(ctx, map[string]interface{}{"email": data.Email})
	if err != nil {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	passHashed := business.hasher.Hash(data.Password + user.Salt)
	if user.Password != passHashed {
		return nil, usermodel.ErrUsernameOrPasswordInvalid
	}
	payload := tokenprovider.TokenPayload{
		UserId: user.Id,
		Role:   user.Role,
	}
	// Fake return value
	return &tokenprovider.Token{Token: payload.Role, Created: time.Now(), Expiry: 0}, nil
}
