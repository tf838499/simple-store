package customer

import (
	"context"
	"log"
	"simple-store/internal/adapter/repository/RedisCache"
)

func (c *CustomerService) SetOuathCode(ctx context.Context, userParam string, codeParam string) error {

	err := c.authRepo.SetUserInfo(ctx, RedisCache.OauthInfo{Name: userParam, Code: codeParam})
	if err != nil {
		log.Println(err.Error())
		c.logger(ctx).Error().Err(err).Msg("failed to insert user code")
		return err
	}

	return err
}
func (c *CustomerService) GetOuathCode(ctx context.Context, userParam string) (string, error) {

	code, err := c.authRepo.GetUserInfo(ctx, RedisCache.OauthInfo{Name: userParam})
	if err != nil {
		log.Println(err.Error())
		c.logger(ctx).Error().Err(err).Msg("failed to insert user code")
		return code, err
	}

	return code, err
}
