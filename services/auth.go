package services

import (
	"bookstore/models"
	"fmt"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type LoggedUser struct {
	ID    uint
	Email string
	Name  string
	Role  string
}

func SetupAuth(authenticatorHandler func(c *gin.Context) (interface{}, error)) (*jwt.GinJWTMiddleware, error) {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: models.UserIdkentityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Println("PayloadFunc ok!")

			if v, ok := data.(*LoggedUser); ok {
				fmt.Println("PayloadFunc ok!")
				fmt.Println("PayloadFunc ok! v", v.Email, v.Name)
				return jwt.MapClaims{
					models.UserIdkentityKey: v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			fmt.Println("IdentityHandler")
			claims := jwt.ExtractClaims(c)
			return &LoggedUser{
				Email: claims[models.UserIdkentityKey].(string),
			}
		},
		Authenticator: authenticatorHandler,
		Authorizator: func(data interface{}, c *gin.Context) bool {
			fmt.Println("Authorizator")
			v, ok := data.(*LoggedUser)

			fmt.Println("Authorizator", data)
			fmt.Println("Authorizator v", v.Role)

			if ok && v.Email == "admin@server.com" {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			fmt.Println("Unauthorized")
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		// TokenLookup is a string in the form of "<source>:<name>" that is used
		// to extract token from the request.
		// Optional. Default value "header:Authorization".
		// Possible values:
		// - "header:<name>"
		// - "query:<name>"
		// - "cookie:<name>"
		// - "param:<name>"
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		// TokenLookup: "query:token",
		// TokenLookup: "cookie:token",

		// TokenHeadName is a string in the header. Default value is "Bearer"
		TokenHeadName: "Bearer",

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	errInit := authMiddleware.MiddlewareInit()
	if errInit != nil {
		log.Fatal("authMiddleware.MiddlewareInit() Error:" + errInit.Error())
	}
	return authMiddleware, err
}
