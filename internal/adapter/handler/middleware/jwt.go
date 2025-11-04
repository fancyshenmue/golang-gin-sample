package middleware

import (
	"context"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/fancyshenmue/golang-gin-sample/internal/infrastructure/config"
	"github.com/fancyshenmue/golang-gin-sample/internal/usecase/auth"
	"github.com/gin-gonic/gin"
)

const identityKey = "id"

// User represents the authenticated user in JWT
type User struct {
	Username  string
	FirstName string
	LastName  string
}

// LoginRequest represents login credentials
type LoginRequest struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// NewJWTMiddleware creates a new JWT middleware
func NewJWTMiddleware(cfg *config.Config, authUseCase *auth.AuthUseCase, requiredGroup string) *jwt.GinJWTMiddleware {
	middleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       cfg.JWT.Realm,
		Key:         []byte(cfg.JWT.SecretKey),
		Timeout:     time.Duration(cfg.JWT.Timeout) * time.Second,
		MaxRefresh:  time.Duration(cfg.JWT.Timeout) * time.Second,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
					"firstname": v.FirstName,
					"lastname":  v.LastName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				Username:  claims[identityKey].(string),
				FirstName: claims["firstname"].(string),
				LastName:  claims["lastname"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginReq LoginRequest
			if err := c.ShouldBind(&loginReq); err != nil {
				return nil, jwt.ErrMissingLoginValues
			}

			ctx := context.Background()
			user, err := authUseCase.Login(ctx, loginReq.Username, loginReq.Password)
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return &User{
				Username:  user.Username,
				FirstName: user.FirstName,
				LastName:  user.LastName,
			}, nil
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if v, ok := data.(*User); ok {
				ctx := context.Background()
				err := authUseCase.ValidateUserGroup(ctx, v.Username, requiredGroup)
				return err == nil
			}
			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		TokenLookup:   "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: cfg.JWT.TokenHeader,
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error: " + err.Error())
	}

	return middleware
}
