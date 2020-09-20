package api

import (
	"crypto/md5"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	jwt "github.com/appleboy/gin-jwt/v2"

	customLogHandle "golang-gin-sample/pkg/loghandle"
	customMongoDatabase "golang-gin-sample/pkg/mongo/pkg"
	customMongoSetup "golang-gin-sample/pkg/mongo/setup"
	customValidation "golang-gin-sample/pkg/validation"
)

var (
	identityKey = "id"
	tokenHeader = os.Getenv("JWT_TOKEN_HEADER")
	jwtRealm    = os.Getenv("JWT_REALM")
	jwtKey      = os.Getenv("JWT_KEY")
)

func AuthMiddleware(chooseGroup string) *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       jwtRealm,
		Key:         []byte(jwtKey),
		Timeout:     time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*User); ok {
				return jwt.MapClaims{
					identityKey: v.UserName,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			return &User{
				UserName: claims[identityKey].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var (
				loginVals    login
				md5InputPass string
			)

			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			userID := loginVals.Username
			password := loginVals.Password

			// check account from mongo
			query := customMongoDatabase.MongoResultHelper{
				Cl:        mongoClient,
				Db:        customMongoSetup.MongoDatabase,
				Coll:      "users",
				FindField: "account",
				FindData:  userID,
			}

			check, err := query.FindOneDataFromMongo()
			customLogHandle.ErrorHandle(
				"GetDocument",
				"GetDocument",
				"GetDocument (Get User Info error)",
				err,
			)

			// check account exist
			if len(check) == 1 {
				data := []byte(password)
				has := md5.Sum(data)
				md5InputPass = fmt.Sprintf("%x", has)
			}

			// check account
			if len(check) == 1 && md5InputPass == fmt.Sprintf("%v", check[0]["password"]) {
				res := &User{
					UserName:  fmt.Sprintf("%v", check[0]["account"]),
					LastName:  fmt.Sprintf("%v", check[0]["lastname"]),
					FirstName: fmt.Sprintf("%v", check[0]["firstname"]),
				}

				return res, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			// check account group from mongo
			queryGroup := customMongoDatabase.MongoResultHelper{
				Cl:        mongoClient,
				Db:        customMongoSetup.MongoDatabase,
				Coll:      "groups",
				FindField: "name",
				FindData:  chooseGroup,
			}

			checkGroup, err := queryGroup.FindGroupInfoFromMongo()
			customLogHandle.ErrorHandle(
				"GetDocument",
				"GetDocument",
				"GetDocument (Get Group Info error)",
				err,
			)

			if v, ok := data.(*User); ok && customValidation.CheckSringInSlice(v.UserName, checkGroup[0].Memeber) {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
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
		TokenHeadName: tokenHeader,

		// TimeFunc provides the current time. You can override it to use another time value. This is useful for testing or if your server uses a different time zone than your tokens.
		TimeFunc: time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	return authMiddleware
}
