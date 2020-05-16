package elmuemasz

import (
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type szamlakResponse struct {
	D struct {
		Results []Szamla `json:"results"`
	} `json:"d"`
}

type Szamla struct {
	Metadata              Metadata `json:"__metadata"`
	EsedekessegTol        string   `json:"EsedekessegTol"`
	EsedekessegIg         string   `json:"EsedekessegIg"`
	SzamlaKelteTol        string   `json:"SzamlaKelteTol"`
	SzamlaKelteIg         string   `json:"SzamlaKelteIg"`
	Vevo                  string   `json:"Vevo"`
	Felhely               string   `json:"Felhely"`
	Szamlaszam            string   `json:"Szamlaszam"`
	Utalvanyszam          string   `json:"Utalvanyszam"`
	Megnevezes            string   `json:"Megnevezes"`
	SzamlaKelte           string   `json:"SzamlaKelte"`
	FizetesiHatarido      string   `json:"FizetesiHatarido"`
	TeljesitesDatum       string   `json:"TeljesitesDatum"`
	KiegyenlitesDatum     string   `json:"KiegyenlitesDatum"`
	NettoAr               string   `json:"NettoAr"`
	AfaErtek              string   `json:"AfaErtek"`
	Osszeg                string   `json:"Osszeg"`
	Egyenleg              string   `json:"Egyenleg"`
	Penznem               string   `json:"Penznem"`
	Tipus                 string   `json:"Tipus"`
	Statusz               string   `json:"Statusz"`
	StronoHivSzam         string   `json:"StronoHivSzam"`
	PostazasiIranyitoszam string   `json:"PostazasiIranyitoszam"`
	PostazasiTelepules    string   `json:"PostazasiTelepules"`
	PostazasiKozterulet   string   `json:"PostazasiKozterulet"`
	RegiRendszer          bool     `json:"RegiRendszer"`
	CsoportFelhelyCim     string   `json:"CsoportFelhelyCim"`
}

type SzamlakFilter struct {
	SzamlaKelteTol time.Time
	SzamlaKelteIg  time.Time
	EsedekessegTol time.Time
	EsedekessegIg  time.Time
	Vevo           string
	Felhely        string
}

func (f SzamlakFilter) String() string {
	layout := "datetime'2006-01-02T00:00:00'"
	var ret []string
	if !f.SzamlaKelteTol.IsZero() {
		ret = append(ret, "(SzamlaKelteTol eq "+f.SzamlaKelteTol.Format(layout)+")")
	}
	if !f.SzamlaKelteIg.IsZero() {
		ret = append(ret, "(SzamlaKelteIg eq "+f.SzamlaKelteIg.Format(layout)+")")
	}
	if !f.EsedekessegTol.IsZero() {
		ret = append(ret, "(EsedekessegTol eq "+f.EsedekessegTol.Format(layout)+")")
	}
	if !f.EsedekessegIg.IsZero() {
		ret = append(ret, "(EsedekessegIg eq "+f.EsedekessegIg.Format(layout)+")")
	}
	if f.Vevo != "" {
		ret = append(ret, "(Vevo eq '"+f.Vevo+"')")
	}
	if f.Felhely != "" {
		ret = append(ret, "(Felhely eq '"+f.Felhely+"')")
	}
	return strings.Join(ret, " and ")
}

func (s Service) Szamlak(filter SzamlakFilter) ([]Szamla, error) {
	var resp szamlakResponse
	query := url.Values{}
	query.Add("$filter", filter.String())
	err := s.get("/Szamlak", query, &resp)
	if err != nil {
		return []Szamla{}, err
	}
	return resp.D.Results, nil
}

func msDatetimeToTime(in string) time.Time {
	r, _ := regexp.Compile("/Date\\((\\d+)\\)/")
	m := r.FindStringSubmatch(in)
	i, _ := strconv.Atoi(m[1])
	ni := i * 1000000
	return time.Unix(0, int64(ni))
}

func (s Service) DownloadPDF(szamla Szamla, filename string) error {

	layout := "2006-01-02T00:00:00"
	path := fmt.Sprintf("/ESzamlaPDFXMLek(Tipus='%s',Vevo='%s',Felhely='%s',Szlaszam='%s',DatumTol=datetime'%s',DatumIg=datetime'%s')/$value",
		"P",
		szamla.Vevo,
		szamla.Felhely,
		szamla.Szamlaszam,
		msDatetimeToTime(szamla.SzamlaKelte).AddDate(0, -3, 0).Format(layout),
		msDatetimeToTime(szamla.SzamlaKelte).AddDate(0, 3, 0).Format(layout),
	)
	return s.download(path, filename)
}

func (s Service) DownloadXML(szamla Szamla, filename string) error {

	layout := "2006-01-02T00:00:00"
	path := fmt.Sprintf("/ESzamlaPDFXMLek(Tipus='%s',Vevo='%s',Felhely='%s',Szlaszam='%s',DatumTol=datetime'%s',DatumIg=datetime'%s')/$value",
		"X",
		szamla.Vevo,
		szamla.Felhely,
		szamla.Szamlaszam,
		msDatetimeToTime(szamla.SzamlaKelte).AddDate(0, -3, 0).Format(layout),
		msDatetimeToTime(szamla.SzamlaKelte).AddDate(0, 3, 0).Format(layout),
	)
	return s.download(path, filename)
}
