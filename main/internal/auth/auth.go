package auth

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/t3mp14r3/curly-octopus/main/internal/config"
	"go.uber.org/zap"
)

type Auth struct {
    client  *http.Client
    addr    string
    logger  *zap.Logger
}

func New(authConfig *config.AuthConfig, logger *zap.Logger) *Auth {
    return &Auth{
        client: &http.Client{Timeout: time.Duration(3) * time.Second},
        addr: authConfig.Addr,
        logger: logger,
    }
}

func (a *Auth) Generate(ctx context.Context, userID string) (string, error) {
    u, err := url.Parse(a.addr)
    
    if err != nil {
        a.logger.Error("failed to parse the request path", zap.Error(err))
        return "", err
    }

    u.Path = "/generate"
    q := u.Query()
    q.Add("login", userID)
    u.RawQuery = q.Encode()

    path := u.String()

    req, err := http.NewRequest("GET", path, nil)
    
    if err != nil {
        a.logger.Error("failed to create new request", zap.Error(err))
        return "", err
    }

    req = req.WithContext(ctx)

    resp, err := a.client.Do(req)
    
    if err != nil {
        a.logger.Error("failed to send the request", zap.Error(err))
        return "", err
    }

    defer resp.Body.Close()

    if resp.StatusCode == http.StatusBadRequest {
        a.logger.Error("request failed - bad request")
        return "", errors.New("response status: bad request")
    }

    body, err := ioutil.ReadAll(resp.Body)

    if err != nil {
        a.logger.Error("failed to read the response body", zap.Error(err))
        return "", err
    }

    return string(body), nil
}

func (a *Auth) Validate(ctx context.Context, token string) bool {
    u, err := url.Parse(a.addr)
    
    if err != nil {
        a.logger.Error("failed to parse the request path", zap.Error(err))
        return false
    }

    u.Path = "/validate"

    path := u.String()

    req, err := http.NewRequest("GET", path, nil)
    
    if err != nil {
        a.logger.Error("failed to create new request", zap.Error(err))
        return false
    }
    
    req = req.WithContext(ctx)

    req.Header.Add("Authorization", token)

    resp, err := a.client.Do(req)
    
    if err != nil {
        a.logger.Error("failed to send the request", zap.Error(err))
        return false
    }

    defer resp.Body.Close()
    
    if resp.StatusCode != http.StatusOK {
        return false
    }

    return true
}

func (a *Auth) Extract(ctx context.Context, tokenStr string) (string, error) {
    tokenStr = strings.TrimPrefix(tokenStr, "Bearer ")

    token, _, err := new(jwt.Parser).ParseUnverified(tokenStr, jwt.MapClaims{})

    if err != nil {
        a.logger.Error("failed to parse the jwt token", zap.Error(err))
        return "", err
    }

    if claims, ok := token.Claims.(jwt.MapClaims); ok {
        login := fmt.Sprint(claims["login"])

        if len(login) == 0 {
            err := errors.New("failed to get login field from claims")
            a.logger.Error("invalid token payload", zap.Error(err))
            return "", err
        }

        return login, nil
    } else {
        err := errors.New("failed to extract token claims")
        a.logger.Error("invalid token", zap.Error(err))
        return "", err
    }
}
