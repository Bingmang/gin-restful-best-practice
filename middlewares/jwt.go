package middlewares

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"gin-restful-best-practice/conf"
	"gin-restful-best-practice/models"
	"net/http"
	"strings"
	"time"
)

type CustomClaims struct {
	UserID   uint
	Username string
	RoleID   uint
	jwt.StandardClaims
}

func CreateJWT(user models.User) (string, error) {
	expiresTime := time.Now().Add(7 * 24 * time.Hour).Unix()
	claims := CustomClaims{
		UserID:   user.ID,
		Username: user.Username,
		RoleID:   user.RoleID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresTime,
			NotBefore: time.Now().Unix(),
			Issuer:    conf.Conf().JWT_ISSUER,
		},
	}
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return tokenClaims.SignedString([]byte(conf.Conf().JWT_SECRET))
}

func ParseJWT(token string) (*CustomClaims, error) {
	jwtToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(conf.Conf().JWT_SECRET), nil
	})
	if err == nil && jwtToken != nil {
		if claim, ok := jwtToken.Claims.(*CustomClaims); ok && jwtToken.Valid {
			return claim, nil
		}
	}
	return nil, err
}

// validate token and set the jwt info into ctx (ctx.Get("jwt") can get the info)
func AuthenticateJWT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		auth := ctx.GetHeader("Authorization")
		if auth == "" || !strings.HasPrefix(auth, "Bearer ") {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "invalid token",
			})
			return
		}
		token := strings.Fields(auth)[1]
		claim, err := ParseJWT(token)
		ctx.Set("jwt", claim)
		if err != nil {
			ctx.Abort()
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Next()
	}
}
