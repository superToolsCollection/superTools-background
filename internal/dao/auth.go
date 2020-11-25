package dao

import "superTools-background/internal/model"

/**
* @Author: super
* @Date: 2020-09-23 20:09
* @Description: 在auth表内获取appKey以及appSecret
**/

func (d *Dao) GetAuth(appKey, appSecret string) (model.Auth, error) {
	auth := model.Auth{AppKey: appKey, AppSecret: appSecret}
	return auth.Get(d.engine)
}
