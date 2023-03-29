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
	RedirectURL: "http://localhost:8080/api/v1/callback",
	// ClientID:     os.Getenv("GOOGLE_OAUTH_CLIENT_ID"),
	// ClientSecret: os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
	ClientID:     "872550806905-v4snmsvth4uaso59cmgrududrqspee7q.apps.googleusercontent.com",
	ClientSecret: "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",
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
			oauthState := generateStateOauthCookie(c) // 需要產生 state 防止CSRF
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
			json.Unmarshal([]byte(data), &googleuserInfo)
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
			common.NewError(common.ErrorCodeParameterInvalid, errors.New(msg), common.WithMsg(msg)))
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
		_ = c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
}
func generateStateOauthCookie(c *gin.Context) string {
	var expiration = 3600

	b := make([]byte, 16)
	rand.Read(b)
	state := base64.URLEncoding.EncodeToString(b)
	c.SetCookie("oauthstate", state, expiration, "/", "localhost", false, true)

	return state
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
