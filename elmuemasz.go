package elmuemasz

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
)

const baseURL string = "https://ker.elmuemasz.hu/sap/opu/odata/sap"
const serviceURL string = "https://ker.elmuemasz.hu/sap/opu/odata/sap/ZGW_UGYFELSZOLGALAT_SRV"

func NewService(client *http.Client) Service {
	return Service{client: client}
}

type Service struct {
	client *http.Client
}

func (s Service) get(path string, query url.Values, response interface{}) error {
	req, err := http.NewRequest(http.MethodGet, serviceURL+path, nil)
	if err != nil {
		return err
	}
	req.Header.Add("Accept", "application/json")

	query.Add("sap-client", "201")
	req.URL.RawQuery = query.Encode()

	resp, err := s.client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusOK {
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
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if resp.StatusCode != http.StatusCreated {
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

func (s Service) download(path string, filePath string) error {
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
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		var errorResp Error
		err = json.Unmarshal(body, &errorResp)
		if err != nil {
			return err
		}
		return errorResp
	}

	contentDisposition := resp.Header.Get("Content-Disposition")
	_, params, err := mime.ParseMediaType(contentDisposition)
	if err != nil {
		return err
	}
	fileName := params["filename"]

	out, err := os.Create(filepath.Join(filePath, fileName))
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, resp.Body)
	return err
}

type Metadata struct {
	ID   string `json:"id"`
	URI  string `json:"uri"`
	Type string `json:"type"`
}
