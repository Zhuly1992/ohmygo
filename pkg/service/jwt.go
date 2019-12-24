package service

import (
	"crypto/rsa"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mygo/form"
	"path/filepath"
	"strings"
	"time"
)

type JwtService struct {
}

func NewJwtService() *JwtService {
	return &JwtService{}
}

const (
	StandardClaimsName = "standardClaims"
	TokenPrefix        = "Zhuly "
)

type StandardClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

var (
	publicKey  *rsa.PublicKey
	privateKey *rsa.PrivateKey
	err        error
)

//token解析
func (svc *JwtService) ParseToken(ctx *gin.Context) {
	jwtToken := ctx.GetHeader("jwtToken")
	var standardClaims StandardClaims
	jwtToken = strings.TrimPrefix(jwtToken, TokenPrefix)
	claims, _ := jwt.ParseWithClaims(jwtToken, &standardClaims, func(token *jwt.Token) (interface{}, error) {
		var err error
		if publicKey == nil {
			publicKey, err = createVerifyKey("/Users/zhuly/git/duapp/mygo/config/rsa/id_rsa.pub")
			if err != nil {
			}
		}
		return publicKey, nil
	})

	ctx.Set(StandardClaimsName, claims.Claims.(*StandardClaims))
	ctx.Next()
}

//获取jwt信息
func (svc *JwtService) GetJwtInfo(ctx *gin.Context) *StandardClaims {
	val, ok := ctx.Get(StandardClaimsName)
	if ok {
		claims, _ := val.(*StandardClaims)
		return claims
	}
	return nil

}

//生成token
func (svc *JwtService) GenToken(ctx *gin.Context, user form.User) string {
	standardClaims := StandardClaims{
		Username: user.Username,
		Password: user.Password,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.TimeFunc().Add(time.Duration(1) * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	privateKey, err = createSignKey("/Users/zhuly/git/duapp/mygo/config/rsa/id_rsa")
	if err != nil {
		fmt.Println(err)
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, standardClaims)
	tokenString, err := token.SignedString(privateKey)
	if err != nil {
	}

	return TokenPrefix + tokenString
}

//私钥生成签名
func createSignKey(privKeyPath string) (*rsa.PrivateKey, error) {
	path, err := filepath.Abs(privKeyPath)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, err
}

//公钥签名破解
func createVerifyKey(pubKeyPath string) (*rsa.PublicKey, error) {
	path, err := filepath.Abs(pubKeyPath)
	if err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(bytes)
	if err != nil {
		return nil, err
	}

	return publicKey, nil
}
