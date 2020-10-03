package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/ProZsolt/elmuemasz"
	"golang.org/x/oauth2"
)

func getELMUEMASZService() (elmuemasz.Service, error) {
	username := os.Getenv("ELMU_USERNAME")
	password := os.Getenv("ELMU_PASSWORD")
	user := elmuemasz.User{
		Username: username,
		Password: password,
	}

	ts, err := elmuemasz.NewTokenSource(user)
	if err != nil {
		return elmuemasz.Service{}, err
	}
	client := oauth2.NewClient(context.Background(), ts)
	return elmuemasz.NewService(client), nil
}

func reportELMUEMASZMeter(meterReading int) error {
	srv, err := getELMUEMASZService()
	if err != nil {
		return err
	}

	felhelyek, err := srv.Felhelyek()
	if err != nil {
		return err
	}
	merodiktalasok, err := srv.MeroDiktalasok(felhelyek[0])
	if err != nil {
		return err
	}
	payload := elmuemasz.MeroDiktalasPayloadFromMeroDiktalas(merodiktalasok[0], time.Now(), meterReading)
	payload.Vevo = felhelyek[0].Vevo
	resp, err := srv.MeroDiktalasPost(payload)
	fmt.Printf("%#v", resp)
	return err
}

func main() {
	err := reportELMUEMASZMeter(15725)
	fmt.Println(err)
}
