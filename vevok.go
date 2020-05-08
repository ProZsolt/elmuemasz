package elmuemasz

import "net/url"

type vevokResponse struct {
	D struct {
		Results []Vevo `json:"results"`
	} `json:"d"`
}

type Vevo struct {
	Metadata               Metadata `json:"__metadata"`
	EszlaCsatolva          bool     `json:"EszlaCsatolva"`
	EszlaAktiv             bool     `json:"EszlaAktiv"`
	EszlaDatum             string   `json:"EszlaDatum"`
	LetrehozasDatum        string   `json:"LetrehozasDatum"`
	Nev                    string   `json:"Nev"`
	ID                     string   `json:"Id"`
	EFelszolitas           bool     `json:"EFelszolitas"`
	DigitFizEml            bool     `json:"DigitFizEml"`
	EszlaKorabban          bool     `json:"EszlaKorabban"`
	RegFajta               string   `json:"RegFajta"`
	Adatvedelem            bool     `json:"Adatvedelem"`
	BizonylatSzama         string   `json:"BizonylatSzama"`
	DigitFizEmlModosithato bool     `json:"DigitFizEmlModosithato"`
	FelhasznaloiAzon       string   `json:"FelhasznaloiAzon"`
	FogyMeroAzon           string   `json:"FogyMeroAzon"`
	FizMod                 string   `json:"FizMod"`
}

func (s Service) Vevok() ([]Vevo, error) {
	var resp vevokResponse
	err := s.get("/Vevok", url.Values{}, &resp)
	if err != nil {
		return []Vevo{}, err
	}
	return resp.D.Results, nil
}
