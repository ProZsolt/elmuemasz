package elmuemasz

import (
	"fmt"
	"net/url"
)

type meroallasokResponse struct {
	D struct {
		Results []Meroallas `json:"results"`
	} `json:"d"`
}

type Meroallas struct {
	Metadata             Metadata `json:"__metadata"`
	LeolvasasOkaLeiras   string   `json:"LeolvasasOkaLeiras"`
	LeolvasasModja       string   `json:"LeolvasasModja"`
	Fogyasztas           string   `json:"Fogyasztas"`
	Statusz              string   `json:"Statusz"`
	LeolvasasOka         string   `json:"LeolvasasOka"`
	Allas                string   `json:"Allas"`
	MertJellemzoLeiras   string   `json:"MertJellemzoLeiras"`
	MertJellemzo         string   `json:"MertJellemzo"`
	Gyariszam            string   `json:"Gyariszam"`
	LeolvasasDatum       string   `json:"LeolvasasDatum"`
	Felhely              string   `json:"Felhely"`
	Vevo                 string   `json:"Vevo"`
	UtolsoLeolvasasDatum string   `json:"UtolsoLeolvasasDatum"`
	UtolsoAllas          string   `json:"UtolsoAllas"`
	Szamlalo             string   `json:"Szamlalo"`
}

func (s Service) Meroallasok(felhely Felhely) ([]Meroallas, error) {
	var resp meroallasokResponse
	path := fmt.Sprintf("/Felhelyek(Vevo='%s',Id='%s')/Meroallasok",
		felhely.Vevo,
		felhely.ID,
	)
	err := s.get(path, url.Values{}, &resp)
	if err != nil {
		return []Meroallas{}, err
	}
	return resp.D.Results, nil
}
