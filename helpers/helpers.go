package helpers

import (
	"context"
	"crypto/sha1"
	"encoding/hex"
	"github.com/graphikDB/graphik/graphik-client-go"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
)

func Hash(val []byte) string {
	h := sha1.New()
	h.Write(val)
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}

func GetClient(ctx context.Context) (*graphik.Client, error) {
	host := viper.GetString("host")
	token := viper.GetString("auth.access_token")
	if host == "" {
		return nil, errors.New("config: empty host")
	}
	if token == "" {
		return nil, errors.New("config: empty auth.access_token")
	}
	tokenSource := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken:  token,
		RefreshToken: viper.GetString("auth.refresh_token"),
	})
	return graphik.NewClient(ctx, host,
		graphik.WithTokenSource(tokenSource),
		graphik.WithRetry(3),
	)
}
