package middleware

import (
	"product-api/db"
	"product-api/form"
	"product-api/model"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

var IdentityKey = "username"

func PayloadFunc(data interface{}) jwt.MapClaims {
	if v, ok := data.(*model.User); ok {
		return jwt.MapClaims{
			IdentityKey: v.Username,
			"fullname":  v.Fullname,
		}
	}
	return jwt.MapClaims{}
}

func IdentityHandler(c *gin.Context) interface{} {
	claims := jwt.ExtractClaims(c)
	return &model.User{
		Username: claims[IdentityKey].(string),
		Fullname: claims["fullname"].(string),
	}
}

func Authenticator(c *gin.Context) (interface{}, error) {
	var loginVals form.Login
	if err := c.ShouldBind(&loginVals); err != nil {
		return "", jwt.ErrMissingLoginValues
	}
	username := loginVals.Username
	password := loginVals.Password

	userVals := model.User{}
	db.DB.First(&userVals, "username = ? && password = ?", username, password)

	if username == userVals.Username && password == userVals.Password {
		return &model.User{
			Username: username,
			Fullname: userVals.Fullname,
		}, nil
	}

	return nil, jwt.ErrFailedAuthentication
}

func Authorizator(data interface{}, c *gin.Context) bool {
	claims := jwt.ExtractClaims(c)
	if v, ok := data.(*model.User); ok && v.Username == claims[IdentityKey] {
		return true
	}

	return false
}

func Unauthorized(c *gin.Context, code int, message string) {
	c.JSON(code, gin.H{
		"code":    code,
		"message": message,
	})
}
