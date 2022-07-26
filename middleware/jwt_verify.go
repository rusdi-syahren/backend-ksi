package middleware

import (
	"crypto/rsa"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// BearerClaims data structure for claims
type BearerClaims struct {
	DeviceID       string `json:"deviceId"`
	UserAuthorized bool   `json:"authorised,bool"`
	jwt.StandardClaims
}

// {
// 	"reffId": "",
// 	"tokenType": "shortToken",
// 	"userType": "patient",
// 	"role": "",
// 	"hospitalId": "MEDISTRA",
// 	"tokenId": "583bbedbf34d4526832e13d020c2baac",
// 	"sub": "e0d8348746314fe1aac2ea099a5d46f4",
// 	"iss": "telemed",
// 	"exp": 1658475970,
// 	"iat": 1658474170
//   }
type TokenInfo struct {
	TokenTypeCode string `json:"tokenTypeCode"` //
	UserId        string `json:"userId"`
	ReffId        string `json:"reffId"`
	UserTypeCode  string `json:"userTypeCode"`
	Role          string `json:"role"`
	HospitalId    string `json:"hospitalId"`
	TokenId       string `json:"tokenId"`
}

type jwtCustomClaims struct {
	TokenInfo
	jwt.StandardClaims
}

// JWTVerify function to verify jwt access token
func JWTVerify(rsaPublicKey *rsa.PublicKey, mustAuthorized bool) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if os.Getenv("NO_TOKEN") == "1" {
				return next(c)
			}

			req := c.Request()
			header := req.Header
			auth := header.Get("Authorization")

			if auth == "" {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization is empty!")
			}

			splitToken := strings.Split(auth, " ")
			if len(splitToken) < 2 {
				return echo.NewHTTPError(http.StatusUnauthorized, "Authorization is empty!")
			}

			tokenStr := splitToken[1]

			token, err := jwt.ParseWithClaims(tokenStr, &BearerClaims{}, func(token *jwt.Token) (interface{}, error) {
				if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
					return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				}
				return rsaPublicKey, nil
			})

			if err != nil {
				return echo.NewHTTPError(http.StatusUnauthorized, "Invalid Token Format")
			}

			if claims, ok := token.Claims.(*BearerClaims); err == nil && token.Valid && ok {
				if mustAuthorized {
					if claims.UserAuthorized {
						c.Set("token", token)
						return next(c)
					}
					fmt.Printf("%+v", claims)
					return echo.NewHTTPError(http.StatusUnauthorized, "Resource need an authorised user")
				}
				c.Set("token", token)
				return next(c)
			} else if ve, ok := err.(*jwt.ValidationError); ok {
				var errorStr string
				if ve.Errors&jwt.ValidationErrorMalformed != 0 {
					errorStr = fmt.Sprintf("Invalid token format: %s", tokenStr)
				} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
					errorStr = "Token has been expired"
				} else {
					errorStr = fmt.Sprintf("Token Parsing Error: %s", err.Error())
				}
				return echo.NewHTTPError(http.StatusUnauthorized, errorStr)
			} else {
				return echo.NewHTTPError(http.StatusUnauthorized, "Unknown token error")
			}
		}
	}
}

func CrateJwtToken(tokenInfo *TokenInfo) (string, error) {

	// Set custom claims
	expiredIn := 24
	if tokenInfo.TokenTypeCode == "shortToken" {
		expiredIn = 1
	}
	claims := &jwtCustomClaims{

		TokenInfo{
			TokenTypeCode: tokenInfo.TokenTypeCode,
			UserId:        tokenInfo.UserId,
			ReffId:        tokenInfo.ReffId,
			UserTypeCode:  tokenInfo.UserTypeCode,
			Role:          tokenInfo.Role,
			HospitalId:    tokenInfo.HospitalId,
			TokenId:       tokenInfo.TokenId,
		},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiredIn)).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodRS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return t, nil
}
