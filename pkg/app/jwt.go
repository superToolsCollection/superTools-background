package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/twinj/uuid"
	"strconv"
	"superTools-background/pkg/errcode"

	"superTools-background/global"
	"time"
)

/**
* @Author: super
* @Date: 2020-09-23 20:02
* @Description: jwt相关操作
**/

type User struct {
	UserId      int
	AccessUuid  string
	RefreshUuid string
}

type UserClaims struct {
	User
	jwt.StandardClaims
}

type TokenDetails struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
	RtExpires    int64
}

type AccessDetails struct {
	AccessUuid string
	UserId     int
}

func (c UserClaims) Valid() (err error) {
	if c.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return errors.New("token is expired")
	}
	if !c.VerifyIssuer(global.JWTSetting.Issuer, true) {
		return errors.New("token's issuer is wrong")
	}
	if c.User.UserId < 1 {
		return errors.New("invalid user in jwt")
	}
	return
}

func GetJWTSecret() []byte {
	return []byte(global.JWTSetting.Secret)
}

//生成token
func GenerateToken(u User) (*TokenDetails, error) {
	nowTime := time.Now()
	td := &TokenDetails{
		AtExpires:   nowTime.Add(15 * time.Minute).Unix(),
		AccessUuid:  uuid.NewV4().String(),
		RtExpires:   nowTime.Add(time.Hour * 24 * 7).Unix(),
		RefreshUuid: uuid.NewV4().String(),
	}
	claims := UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.AtExpires,
			Issuer:    global.JWTSetting.Issuer,
		},
		User: User{
			UserId:      u.UserId,
			AccessUuid:  td.AccessUuid,
			RefreshUuid: td.RefreshUuid,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(GetJWTSecret())
	if err != nil {
		return nil, err
	}
	td.AccessToken = token

	claims = UserClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: td.RtExpires,
			Issuer:    global.JWTSetting.Issuer,
		},
		User: User{
			UserId: u.UserId,
			AccessUuid:  td.AccessUuid,
			RefreshUuid: td.RefreshUuid,
		},
	}
	tokenClaims = jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err = tokenClaims.SignedString(GetJWTSecret())
	if err != nil {
		return nil, err
	}
	td.RefreshToken = token

	return td, nil
}

//解析token
func ParseToken(token string) (*User, error) {
	if token == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}
	claims := UserClaims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(GetJWTSecret()), nil
	})
	if err != nil {
		return nil, err
	}
	return &claims.User, err
}

func SaveAuth(userId int, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	rt := time.Unix(td.RtExpires, 0)
	now := time.Now()

	errAccess := global.RedisEngine.Set(context.TODO(),td.AccessUuid, strconv.Itoa(userId), at.Sub(now)).Err()
	if errAccess != nil {
		return errAccess
	}
	errRefresh := global.RedisEngine.Set(context.TODO(), td.RefreshUuid, strconv.Itoa(userId), rt.Sub(now)).Err()
	if errRefresh != nil {
		return errRefresh
	}
	return nil
}

func VerfyToken(token string, id string) *errcode.Error {
	ecode := errcode.Success
	user, err := ParseToken(token)
	if err != nil {
		switch err.(*jwt.ValidationError).Errors {
		case jwt.ValidationErrorExpired:
			ecode = errcode.UnauthorizedTokenTimeout
		default:
			ecode = errcode.UnauthorizedTokenError
		}
	} else {
		userId, _ := strconv.Atoi(id)
		if userId != user.UserId {
			ecode = errcode.UnauthorizedTokenError
		} else {
			ad := &AccessDetails{
				AccessUuid: user.AccessUuid,
				UserId:     userId,
			}
			id, AuthErr := GetAuth(ad)
			if AuthErr != nil {
				ecode = errcode.UnauthorizedTokenError
			} else if id != user.UserId {
				ecode = errcode.UnauthorizedTokenError
			}
		}
	}
	return ecode
}

func GetAuth(ad *AccessDetails) (int, error) {
	userId, err := global.RedisEngine.Get(context.TODO(), ad.AccessUuid).Result()
	if err != nil {
		return 0, err
	}
	userID, _ := strconv.ParseInt(userId, 10, 64)
	return int(userID), nil
}

func DeleteAuth(givenUuid string) (int64,error) {
	deleted, err := global.RedisEngine.Del(context.TODO(), givenUuid).Result()
	if err != nil {
		return 0, err
	}
	return deleted, nil
}