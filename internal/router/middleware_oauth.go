package router

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"simple-store/internal/app"
	"simple-store/internal/domain/common"
	"simple-store/internal/router/api/reponse"

	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type OAuthMiddleware struct {
	app *app.Application
}

func NewOAuthMiddleware(app *app.Application) *OAuthMiddleware {
	return &OAuthMiddleware{
		app: app,
	}
}

// Scopes: OAuth 2.0 scopes provide a way to limit the amount of access that is granted to an access token.
var googleOauthConfig = &oauth2.Config{
	RedirectURL:  "http://localhost:8080/api/v1/callback",
	ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	Scopes:       []string{"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint:     google.Endpoint,
}

const oauthGoogleUrlAPI = "https://www.googleapis.com/oauth2/v2/userinfo?access_token="

// SetGeneralMiddlewares add general-purpose middlewares
func (m *OAuthMiddleware) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		cookieToken, err := c.Cookie("oauthstate")
		if err != nil || cookieToken != c.GetHeader("State") {
			oauthState, err := generateStateOauthCookie(c, c.Request.URL.Path) // 需要產生 state 防止CSRF
			if err != nil {
				reponse.RespondWithError(c,
					common.NewError(common.ErrorCodeParameterInvalid, errors.New("invalid parameter"), common.WithMsg("invalid parameter")))
				return
			}
			c.Redirect(http.StatusTemporaryRedirect, googleOauthConfig.AuthCodeURL(oauthState))
			c.AbortWithStatus(http.StatusTemporaryRedirect)
			return
		} else {
			data, err := m.app.CustomerService.GetOuathCode(ctx, cookieToken)
			if err != nil {
				reponse.RespondWithError(c,
					common.NewError(common.ErrorCodeParameterInvalid, errors.New("invalid parameter"), common.WithMsg("invalid parameter")))
				return
			}

			type googleUser struct {
				Id             string `json:"id"`
				Email          string `json:"email"`
				Verified_email bool   `json:"given_name"`
				Picture        string `json:"picture"`
			}
			var googleuserInfo googleUser

			err = json.Unmarshal([]byte(data), &googleuserInfo)
			if err != nil {
				reponse.RespondWithError(c,
					common.NewError(common.ErrorCodeParameterInvalid, errors.New("invalid parameter"), common.WithMsg("invalid parameter")))
				return
			}
			c.Set("googleEmail", googleuserInfo.Email)
			c.Next()
		}
	}
}
func (m *OAuthMiddleware) Callback(c *gin.Context) {
	ctx := c.Request.Context()
	oauthState, _ := c.Cookie("oauthstate")
	fmt.Println(oauthState)
	if c.Query("state") != oauthState {
		msg := "invalid oauth google state"
		log.Println(msg)
		reponse.RespondWithError(c,
			common.NewError(common.ErrorCodeAuthNotAuthenticated, errors.New(msg), common.WithMsg(msg)))
		return

	}
	// fmt.Println(c.Query("code"))
	data, err := getUserDataFromGoogle(ctx, c.Query("code"))
	if err != nil {
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	err = m.app.CustomerService.SetOuathCode(ctx, oauthState, string(data))
	if err != nil {
		msg := "invalid oauth google state"
		reponse.RespondWithError(c,
			common.NewError(common.ErrorCodeAuthNotAuthenticated, errors.New(msg), common.WithMsg(msg)))
		return
	}
	// targetURL := c.Request.URL.Path
	// targetURL := c.Get("targetURL")
	targetURL, err := decodeStateOauthCookie(c, c.Query("state"))
	if err != nil {
		msg := "invalid oauth google state"
		reponse.RespondWithError(c,
			common.NewError(common.ErrorCodeAuthNotAuthenticated, errors.New(msg), common.WithMsg(msg)))
		return
	}
	c.Set("email", data)
	c.Redirect(http.StatusSeeOther, targetURL)
}
func generateStateOauthCookie(c *gin.Context, targeturl string) (string, error) {
	var expiration = 3600

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	b = append(b, []byte(targeturl)...)
	state := base64.URLEncoding.EncodeToString(b)
	c.SetCookie("oauthstate", state, expiration, "/", "localhost", false, true)
	return state, err
}
func decodeStateOauthCookie(c *gin.Context, state string) (string, error) {

	b, err := base64.URLEncoding.DecodeString(state)
	if err != nil {
		return "", err
	}
	userOriginalURL := string(b[16:])

	return userOriginalURL, nil
}
func getUserDataFromGoogle(ctx context.Context, code string) ([]byte, error) {

	token, err := googleOauthConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("code exchange wrong: %s", err.Error())
	}

	response, err := http.Get(oauthGoogleUrlAPI + token.AccessToken)
	if err != nil {
		return nil, fmt.Errorf("failed getting user info: %s", err.Error())
	}

	defer response.Body.Close()
	contents, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed read response: %s", err.Error())
	}

	return contents, nil
}
