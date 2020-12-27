package auth

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/graphikDB/graphikctl/helpers"
	"github.com/graphikDB/graphikctl/version"
	"github.com/pkg/errors"
	"github.com/skratchdot/open-golang/open"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"golang.org/x/oauth2"
	"html/template"
	"net"
	"net/http"
	"sync"
	"time"
)

func init() {

	Auth.AddCommand(login)
}

var Auth = &cobra.Command{
	Use:     "auth",
	Short:   "authentication/authorization subcommands",
	Version: version.Version,
}

var login = &cobra.Command{
	Use:   "login",
	Short: "launch a login flow to an identity provider",
	Run: func(cmd *cobra.Command, args []string) {
		config, err := oauthConfig()
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		state := helpers.Hash([]byte(uuid.New().String()))
		if !viper.InConfig("auth.code_verifier") {
			verifier := uuid.New().String()
			viper.Set("auth.code_verifier", verifier)
		}

		challenge := helpers.Hash([]byte(viper.GetString("auth.code_verifier")))

		link := config.AuthCodeURL(state,
			oauth2.SetAuthURLParam("code_challenge_method", "S256"),
			oauth2.SetAuthURLParam("code_challenge", challenge),
			)
		mux := http.NewServeMux()
		server := &http.Server{Addr: viper.GetString("server.port"), Handler: mux}
		mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			// get the authorization code
			code := r.URL.Query().Get("code")
			if code == "" {
				http.Error(w, "Error: could not find 'code' URL parameter", http.StatusBadRequest)
				return
			}
			stateParam := r.URL.Query().Get("state")
			if stateParam != state {
				http.Error(w, "Error: mismatching state param", http.StatusBadRequest)
				return
			}
			token, err := config.Exchange(r.Context(), code, oauth2.SetAuthURLParam("code_verifier", viper.GetString("auth.code_verifier")))
			if err != nil {
				fmt.Println(err.Error())
				http.Error(w, "Error: failed to exchange authorization code", http.StatusUnauthorized)
				return
			}
			tmpl, err := template.New("").Parse(loginTmpl)
			if err != nil {
				fmt.Println(err.Error())
				http.Error(w, "Error: failed to parse login template", http.StatusInternalServerError)
				return
			}
			accessToken := token.AccessToken
			idToken := token.Extra("id_token")
			viper.Set("auth.access_token", accessToken)
			viper.Set("auth.id_token", idToken)
			if err := viper.WriteConfig(); err != nil {
				fmt.Println(err.Error())
				http.Error(w, "Error: failed to save config", http.StatusInternalServerError)
				return
			}
			if err := tmpl.Execute(w, map[string]interface{}{
				"access_token": accessToken,
				"id_token":     idToken,
			}); err != nil {
				fmt.Println(err.Error())
				http.Error(w, "Error: failed to execute login template", http.StatusInternalServerError)
				return
			}
			go server.Close()
		})
		lis, err := net.Listen("tcp", viper.GetString("server.port"))
		if err != nil {
			fmt.Printf("failed to create listener: %s", err.Error())
			return
		}
		defer lis.Close()
		wg := &sync.WaitGroup{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			if err := server.Serve(lis); err != nil && err != http.ErrServerClosed {
				fmt.Printf("server failure: %s", err)
				return
			}
		}()
		wg.Add(1)
		time.Sleep(1 * time.Second)
		go func() {
			defer wg.Done()
			// open a browser window to the authorizationURL
			if err = open.Start(link); err != nil {
				fmt.Printf("can't open browser to URL %s: %s", link, err)
				return
			}
		}()
		wg.Wait()
	},
	Version: version.Version,
}

func oauthConfig() (*oauth2.Config, error) {
	var (
		openID       = viper.GetString("auth.open_id")
		clientId     = viper.GetString("auth.client_id")
		clientSecret = viper.GetString("auth.client_secret")
		redirect     = viper.GetString("auth.redirect")
	)

	if openID == "" {
		return nil, errors.New("config: empty auth.open_id")
	}
	if clientId == "" {
		return nil, errors.New("config: empty auth.client_id")
	}
	if redirect == "" {
		return nil, errors.New("config: empty auth.redirect")
	}
	resp, err := http.Get(openID)
	if err != nil {
		return nil, errors.Wrap(err, "config: failed to get openid metadata")
	}
	defer resp.Body.Close()
	metadata := map[string]interface{}{}
	if err := json.NewDecoder(resp.Body).Decode(&metadata); err != nil {
		return nil, err
	}
	return &oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  metadata["authorization_endpoint"].(string),
			TokenURL: metadata["token_endpoint"].(string),
		},
		RedirectURL: redirect,
		Scopes:      viper.GetStringSlice("auth.scopes"),
	}, nil
}

const loginTmpl = `
		<html>
			<body>
				<h1>Login successful!</h1>
				<details><summary>access token</summary>
				<a href="https://jwt.io/#debugger-io?token={{ .access_token }}">Debug</a>
				<p>
				{{ .access_token }}
				</p>
				</details>
				<details><summary>id token</summary>
				<a href="https://jwt.io/#debugger-io?token={{ .id_token }}">Debug</a>
				<p>
				{{ .id_token }}
				</p>
				</details>
				<h2>You can close this window and return to your terminal.</h2>
			</body>
		</html>`
