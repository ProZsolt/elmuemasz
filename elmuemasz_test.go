package elmuemasz

import (
	"context"
	"os"
	"testing"

	"golang.org/x/oauth2"
)

func TestThings(t *testing.T) {
	username := os.Getenv("ELMU_USERNAME")
	password := os.Getenv("ELMU_PASSWORD")
	vevoID := os.Getenv("ELMU_VEVO")
	felhelyID := os.Getenv("ELMU_FELHELY")

	user := User{Username: username,
		Password: password,
	}
	ts, err := NewTokenSource(user)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
	client := oauth2.NewClient(context.Background(), ts)
	srv := NewService(client)
	vevok, err := srv.Vevok()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	want := vevoID
	got := vevok[0].ID
	if want != got {
		t.Errorf("Got this: %v", vevok[0])
	}
	felhelyek, err := srv.Felhelyek()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	want = felhelyID
	got = felhelyek[0].ID
	if want != got {
		t.Errorf("Got this: %v", felhelyek[0])
	}

	eSzamlak, err := srv.ESzamlak()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	want = vevoID
	got = eSzamlak[0].Vevo
	if want != got {
		t.Errorf("Got this: %v", eSzamlak[0])
	}

	filter := SzamlakFilter{
		Vevo: vevoID,
	}
	szamlak, err := srv.Szamlak(filter)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = srv.DownloadPDF(szamlak[0], "responses")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	err = srv.DownloadXML(szamlak[0], "responses")
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}
