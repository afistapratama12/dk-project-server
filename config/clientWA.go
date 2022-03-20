package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
)

func ClientWA() *twilio.RestClient {
	err := godotenv.Load()
	FailOnError(err, 12, "config/clientWA.go")

	accSID := os.Getenv("T_ACC_SID")
	authToken := os.Getenv("T_AUTH_TOKEN")

	client := twilio.NewRestClientWithParams(twilio.RestClientParams{
		Username:   accSID,
		Password:   authToken,
		AccountSid: accSID,
	})

	return client
}
