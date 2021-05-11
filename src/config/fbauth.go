package config

import (
	"context"
	"log"
	"path/filepath"

	firebase "firebase.google.com/go/v4"
	"firebase.google.com/go/v4/auth"
	"google.golang.org/api/option"
)

func InitAuth() (*auth.Client, error) {
	serviceAccountKeyFilePath, err := filepath.Abs("./serviceAccountKey.json")
	if err != nil {
		log.Fatal("Unable to load serviceAccountKeys.json file")
	}

	opt := option.WithCredentialsFile(serviceAccountKeyFilePath)

	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, err
	}

	auth, errAuth := app.Auth(context.Background())
	if errAuth != nil {
		return nil, err
	}

	return auth, nil
}
