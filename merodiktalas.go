package elmuemasz

import (
	"fmt"
	"net/url"
	"strconv"
	"time"
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

type MeroDiktalasPayload struct {
	Vevo           string            `json:"Vevo"`
	Felhely        string            `json:"Felhely"`
	LeolvasasDatum string            `json:"LeolvasasDatum"`
	Megerosites    bool              `json:"Megerosites"`
	ArAfaCheck     bool              `json:"ArAfaCheck"`
	Rogzites       []RogzitesPayload `json:"Rogzites"`
}
type RogzitesPayload struct {
	MeroAzonosito string `json:"MeroAzonosito"`
	Szamlalo      string `json:"Szamlalo"`
	AktAllas      string `json:"AktAllas"`
	LeolvasasOka  string `json:"LeolvasasOka"`
}

func MeroDiktalasPayloadFromMeroDiktalas(md MeroDiktalas, ld time.Time, aa int) MeroDiktalasPayload {
	layout := "2006-01-02T00:00:00"
	return MeroDiktalasPayload{
		Vevo:           md.Vevo,
		Felhely:        md.Felhely,
		LeolvasasDatum: ld.Format(layout),
		Megerosites:    md.Megerosites,
		ArAfaCheck:     md.ArAfaCheck,
		Rogzites: []RogzitesPayload{
			{
				MeroAzonosito: md.MeroAzonosito,
				Szamlalo:      md.Szamlalo,
				AktAllas:      strconv.Itoa(aa),
				LeolvasasOka:  md.LeolvasasOka,
			},
		},
	}
}

func (s Service) MeroDiktalasPost(payload MeroDiktalasPayload) (MeroDiktalas, error) {
	var resp meroDiktalasPostResponse
	path := fmt.Sprintf("/Felhelyek(Vevo='%s',Id='%s')/MeroDiktalas",
		payload.Vevo,
		payload.Felhely,
	)
	err := s.post(path, payload, &resp)
	if err != nil {
		return MeroDiktalas{}, err
	}
	return resp.D, nil
}
