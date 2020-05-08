package elmuemasz

import (
	"fmt"
	"net/url"
)

type meroDiktalasResponse struct {
	D struct {
		Results []MeroDiktalas `json:"results"`
	} `json:"d"`
}

type meroDiktalasPostResponse struct {
	D MeroDiktalas `json:"d"`
}

type MeroDiktalas struct {
	Metadata             Metadata    `json:"__metadata"`
	ArAfaCheck           bool        `json:"ArAfaCheck"`
	ArAfaDiktKezdet      interface{} `json:"ArAfaDiktKezdet"`
	ArAfaDiktVege        interface{} `json:"ArAfaDiktVege"`
	Megerosites          bool        `json:"Megerosites"`
	MinAllas             string      `json:"MinAllas"`
	MaxAllas             string      `json:"MaxAllas"`
	DiktalasPeriodus     string      `json:"DiktalasPeriodus"`
	LeolvasasOka         string      `json:"LeolvasasOka"`
	AktAllas             string      `json:"AktAllas"`
	UtolsoLeolvasasDatum string      `json:"UtolsoLeolvasasDatum"`
	Allas                string      `json:"Allas"`
	MertJellemzo         string      `json:"MertJellemzo"`
	MertJellemzoLeiras   string      `json:"MertJellemzoLeiras"`
	Szamjegyszam         string      `json:"Szamjegyszam"`
	LeolvasasDatum       string      `json:"LeolvasasDatum"`
	DiktEnd              string      `json:"DiktEnd"`
	DiktStart            string      `json:"DiktStart"`
	Felhely              string      `json:"Felhely"`
	Vevo                 string      `json:"Vevo"`
	MeroAzonosito        string      `json:"MeroAzonosito"`
	Szamlalo             string      `json:"Szamlalo"`
	Mertekegyseg         string      `json:"Mertekegyseg"`
	LeolvBizAzon         string      `json:"LeolvBizAzon"`
	Rogzites             Rogzites    `json:"Rogzites"`
}

type Rogzites struct {
	Results []MeroRogzites `json:"results"`
}

type MeroRogzites struct {
	Metadata      Metadata `json:"__metadata"`
	LeolvasasOka  string   `json:"LeolvasasOka"`
	AktAllas      string   `json:"AktAllas"`
	MeroAzonosito string   `json:"MeroAzonosito"`
	Szamlalo      string   `json:"Szamlalo"`
}

func (s Service) MeroDiktalasok(felhely Felhely) ([]MeroDiktalas, error) {
	var resp meroDiktalasResponse
	path := fmt.Sprintf("/Felhelyek(Vevo='%s',Id='%s')/MeroDiktalas",
		felhely.Vevo,
		felhely.ID,
	)
	err := s.get(path, url.Values{}, &resp)
	if err != nil {
		return []MeroDiktalas{}, err
	}
	return resp.D.Results, nil
}
