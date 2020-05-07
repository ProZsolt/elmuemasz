package elmuemasz

type eSzamlakResponse struct {
	D struct {
		Results []ESzamla `json:"results"`
	} `json:"d"`
}

type ESzamla struct {
	Metadata           Metadata `json:"__metadata"`
	Vevo               string   `json:"Vevo"`
	VevoNev            string   `json:"VevoNev"`
	FizetoNev          string   `json:"FizetoNev"`
	SzamlaMod          string   `json:"SzamlaMod"`
	SzamlaKezbesito    string   `json:"SzamlaKezbesito"`
	ErvenyKezdet       string   `json:"ErvenyKezdet"`
	BankszamlaSzam     string   `json:"BankszamlaSzam"`
	BankszamlaNev      string   `json:"BankszamlaNev"`
	EszamlaEmail       bool     `json:"EszamlaEmail"`
	Jatek              bool     `json:"Jatek"`
	TajekoztatoElfogad bool     `json:"TajekoztatoElfogad"`
	ReklamElfogad      bool     `json:"ReklamElfogad"`
	SzamlaModKod       string   `json:"SzamlaModKod"`
}

func (s Service) ESzamlak() ([]ESzamla, error) {
	var resp eSzamlakResponse
	err := s.get("/ESzamlak", &resp)
	if err != nil {
		return []ESzamla{}, err
	}
	return resp.D.Results, nil
}
