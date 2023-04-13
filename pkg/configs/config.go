package configs

import (
	"bytes"
	"fmt"
	"github.com/spf13/viper"
	"io"
	"log"
	"net/http"
)

func NewServerConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("cmd/server/")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}

func SetWebhook() {
	client := &http.Client{}
	webHookUrl := viper.GetString("localtunnel.webHookUrl")
	reqBody := []byte(`{"webHookUrl": "` + webHookUrl + `"}`)
	request, err := http.NewRequest(
		"POST",
		"https://api.monobank.ua/personal/webhook",
		bytes.NewBuffer(reqBody),
	)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("X-Token", viper.GetString("monobank.token"))
	resp, err := client.Do(request)
	defer resp.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("WebHook: ", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("response body: \n", string(reqBody))
	log.Println("response body: \n", string(body))
}
