package token

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

var(
	ErrMissingHeader  =	errors.New("The 'Authorization' Header is empty.")
	TokenLifeInHour   = parseTokenLife()
	TokenLifeInSecond = int(time.Hour * TokenLifeInHour / time.Second)
)

// Context is the context of the JSON web token.
type Context struct {
	ID		uint64
	Role	int						//
}

func parseTokenLife() time.Duration {
	nHour:=viper.GetInt64("tokenLife")
	return time.Duration(nHour)
}

// secretFunc validates the secret format.
func secretFunc(secret string) jwt.Keyfunc {
	return func(token *jwt.Token) (i interface{}, e error) {
		// Make sure the `alg` is what we except.
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(secret), nil
	}
}

// Parse validates the token with the specified secret,
// and returns the context if the token was valid.
func Parse(tokenStr string,secret string) (*Context,error) {
	ctx:=&Context{}

	// Parse the token.
	token,err:=jwt.Parse(tokenStr,secretFunc(secret))

	// Parse error.
	if err!=nil {
		return nil,err

		// Read the token if it's valid.
	}else if claims,ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		ctx.ID=uint64(claims["id"].(float64))
		ctx.Role=int(claims["role"].(float64))
		return ctx,nil

		// Other errors.
	}else {
		return ctx,err
	}
}

// ParseRequest gets the token from the header and
// pass it to the Parse function to parses the token.
func ParseRequest(c *gin.Context) (*Context,error) {
	header := c.Request.Header.Get("Authorization")

	secret := viper.GetString("jwt_secret")

	if len(header) == 0 {
		return &Context{},ErrMissingHeader
	}

	var t string

	fmt.Sscanf(header,"Bearer %s",&t)

	return Parse(t,secret)
}

// Sign signs the context with the specified secret.
func Sign(ctx *gin.Context,c Context,secret string) (tokenString string,err error) {
	//Load the jwt secret from the Gin config if the secret isn't specified.
	if secret == "" {
		secret = viper.GetString("jwt_secret")
	}

	//The token content.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":		c.ID,
		"role":		c.Role,
		"nbf":		time.Now().Unix(),
		"iat":		time.Now().Unix(),
		"exp":		TokenLifeInSecond,
	})

	tokenString,err = token.SignedString([]byte(secret))

	return
}