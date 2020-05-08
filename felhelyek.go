package elmuemasz

import "net/url"

type felhelyekResponse struct {
	D struct {
		Results []Felhely `json:"results"`
	} `json:"d"`
}

type Felhely struct {
	Metadata Metadata `json:"__metadata"`
	ID       string   `json:"Id"`
	Vevo     string   `json:"Vevo"`
	Cim      string   `json:"Cim"`
}

func (s Service) Felhelyek() ([]Felhely, error) {
	var resp felhelyekResponse
	err := s.get("/Felhelyek", url.Values{}, &resp)
	if err != nil {
		return []Felhely{}, err
	}
	return resp.D.Results, nil
}
