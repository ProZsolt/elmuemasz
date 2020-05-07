package elmuemasz

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const baseURL string = "https://ker.elmuemasz.hu/sap/opu/odata/sap"
const serviceURL string = "https://ker.elmuemasz.hu/sap/opu/odata/sap/ZGW_UGYFELSZOLGALAT_SRV"

func NewService(client *http.Client) Service {
	return Service{client: client}
}

type Service struct {
	client *http.Client
}

func (s Service) get(path string, response interface{}) error {
	req, err := http.NewRequest(http.MethodGet, serviceURL+path, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add("sap-client", "201")
	req.URL.RawQuery = q.Encode()

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		var errorResp Error
		err = json.Unmarshal(body, &errorResp)
		if err != nil {
			return err
		}
		return errorResp
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}
	return nil
}

func (s Service) post(path string, data interface{}, response interface{}) error {
	jsonValue, err := json.Marshal(data)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodGet, serviceURL+path, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")

	q := req.URL.Query()
	q.Add("sap-client", "201")
	req.URL.RawQuery = q.Encode()

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != 201 {
		var errorResp Error
		err = json.Unmarshal(body, &errorResp)
		if err != nil {
			return err
		}
		return errorResp
	}

	err = json.Unmarshal(body, response)
	if err != nil {
		return err
	}
	return nil
}

type Metadata struct {
	ID   string `json:"id"`
	URI  string `json:"uri"`
	Type string `json:"type"`
}
