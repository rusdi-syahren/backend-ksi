package middleware

import (
	"crypto/rsa"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestBearerVerify(t *testing.T) {
	expiredToken := "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiI0OWZlNzc1NC1mODJlLTQ5MDctOTIyOC03YzI2YTVjZDYyNDQiLCJhdXRob3Jpc2VkIjp0cnVlLCJkZXZpY2VJZCI6IjAwMTYiLCJleHAiOjE1MDY0MTE5MDMsImlhdCI6MTUwNjQxMTMwMywiaXNzIjoiYmhpbm5la2EuY29tIiwic3ViIjoiVVNSMTYxMDAwOTc4In0.q7ieeMa5vyXyXildm8cWM3EJGdN6baQuy3DUwfeMNaWwkMOZXwOHf8d3x1QY4xJ_5NYj-UTQD7966EymtPFbHKZCLg2YL9FI_nWcNHKSPVfo4Yvz9LGd5eh_IX6xFg5xpDLLYnPn1FlWSAGiMaxn-sf7f87nNsfmeLmBffQIQiLNoaJXLl4d-ZTgpYNnxoZZ8Nf0ORukqhL8y0aU0XSbRwA37K2086qrPO6ZJCwPhhX9zucii0st5633OqAyYdTtsQ-A0DEujD-PNeSWMbx1Bsi4E8y-UjKZbonDydbWn_qOmlSABleR1vfSUU4d0tTMgJEt7EcOFKHnJIeDLKy68g"
	validAccessToken := "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXZpY2VJZCI6IjY4OTIyIiwiYXV0aG9yaXNlZCI6dHJ1ZSwiaWF0IjoxNDkxMjA2OTE5LCJhdWQiOiI0OWZlNzc1NC1mODJlLTQ5MDctOTIyOC03YzI2YTVjZDYyNDQiLCJpc3MiOiJiaGlubmVrYS5jb20iLCJzdWIiOiJVU1IwOTEwMDExNzgifQ.XzHUptvHe7y8FEUYaP4QWP22g-hu8VvQKC2-dM2GW9ZdSqaw7V3HmTBavU3aLnt8y75zHTA-MAEkauSPUJ7IFxI-YUME2IX1_UZjkr6NGZ4ntcU00p4mfGqd0KjDZLx0hzBzFjVPWuKB9N-ksFqOiKPIGTmC36xM_H-8nGM3x_I7u3bGlXMZ40KblkShDctqwphQG4QtHf7oq-AQFwQtmtg63x-egVpwyUdz1tn5pVi0r0R_427ZlBBKXOqtAWwxuecB1ITesQIxGbifTUwdS87b_wVcI68SnNfUdU0tQqk8Yj3B85Cq_CDBaq86HIkDwRLIx0vDQC8awl8bVRP0tA"
	anonymousAccessToken := "Bearer eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJkZXZpY2VJZCI6IjY4OTIyIiwiYXV0aG9yaXNlZCI6ZmFsc2UsImlhdCI6MTQ5MTIxMDMzOCwiYXVkIjoiNDlmZTc3NTQtZjgyZS00OTA3LTkyMjgtN2MyNmE1Y2Q2MjQ0IiwiaXNzIjoiYmhpbm5la2EuY29tIiwic3ViIjoiNDlmZTc3NTQtZjgyZS00OTA3LTkyMjgtN2MyNmE1Y2Q2MjQ0In0.BnKC6lVf_l9PSashjq53Qe4etGafHy1lCYrGwnQdzvf-5ycGhJOQPeqVyQb9n-9ZLwfXKBL9-af8ZZUe7_doWkRBKK2uzM0CbXv35LxdAvmDjax9qCrJHadh-wdK7Vw6-lwu_-Wp-NipE89kpw_nSleUDIBA8LNEOUcp5ynxBsUvZNgVS6rdZWOMHCI8eADiwb_20hsPbAHX0HxOzg9uKr2kw7MygWPT6JZzrcLzBEewOCip5HYhMDYKEGu4Ful1almYHDOijCigyRhzwVVsUzSoI-JHMlqJvclcZzsASTNl1MJFcHNRhb-maJT5sF3Kiuy0jGhPpKr4Sc3j_68smQ"

	t.Run("Should No Error with valid access token", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, validAccessToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		verifyKey, _ := getPublicKey(ValidPublicKey)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "klinisia.id"))
		})

		mw := JWTVerify(verifyKey, true)(handler)
		err := mw(c)
		assert.NoError(t, err)
	})

	t.Run("Should No Error with Anonymous Access Token", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, anonymousAccessToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		verifyKey, _ := getPublicKey(ValidPublicKey)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "klinisia.id"))
		})

		mw := JWTVerify(verifyKey, false)(handler)
		err := mw(c)
		assert.NoError(t, err)
	})

	t.Run("Should Error 'need an authorised user'", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, anonymousAccessToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		verifyKey, _ := getPublicKey(ValidPublicKey)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "klinisia.id"))
		})

		mw := JWTVerify(verifyKey, true)(handler)
		err := mw(c)
		assert.Error(t, err)
	})

	t.Run("Should Error 'Access Token Expired'", func(t *testing.T) {
		e := echo.New()
		req := httptest.NewRequest(echo.GET, "/", nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, expiredToken)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		verifyKey, _ := getPublicKey(ValidPublicKey)

		handler := echo.HandlerFunc(func(c echo.Context) error {
			return c.JSON(http.StatusOK, c.String(http.StatusOK, "klinisia.id"))
		})

		mw := JWTVerify(verifyKey, true)(handler)
		err := mw(c)
		assert.Error(t, err)
	})

}

const ValidPublicKey = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAvSd5tzIKV7rlA4uKn1vl
9cg9hNzjWzbT6cvJy5TED2pilUw6LJZhV+ieV08BX2eoG17ygbs8qs7jAcHPzMWw
MGCIayy8XBNG36diPV9ukFdpLczeov0f6gP093w/C2Y6cLQRN3iBlToZKIR6qf0i
PoFMaqiFa8Ys2OmeEdL2egNm+IxGXxyRB9NOwWGjvt5w7PC41+iIGA/AV9EH7FVe
7bcnBsSGXy3kCTneI/X0pcZq1M7cYEPvzXOtq35xzDrmMSoSPo3O06GyPZNA7S4A
iMpw83U1XNmUsVq7lpXP6sROuxEmPfIVunz13DqVZXOTrtkJoONSgNFJ0VbLKUwb
eQIDAQAB
-----END PUBLIC KEY-----
`

func getPublicKey(publicKey string) (*rsa.PublicKey, error) {
	r := strings.NewReader(publicKey)
	verifyBytes, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}

	verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
	if err != nil {
		return nil, err
	}
	return verifyKey, nil
}
