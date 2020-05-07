package elmuemasz

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"

	"golang.org/x/oauth2"
)

// User represents the user credentials
type User struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

type authResponse struct {
	D struct {
		AuthCode string `json:"AuthCode"`
	} `json:"d"`
}

type tokenResponse struct {
	D struct {
		GetToken struct {
			ExpiresIn      string `json:"ExpiresIn"`
			RenewToken     string `json:"RenewToken"`
			TokenCode      string `json:"TokenCode"`
			WarningTimeOut string `json:"WarningTimeOut"`
		} `json:"GetToken"`
	} `json:"d"`
}

type renewResponse struct {
	D struct {
		RenewToken struct {
			ExpiresIn      string `json:"ExpiresIn"`
			RenewToken     string `json:"RenewToken"`
			TokenCode      string `json:"TokenCode"`
			WarningTimeOut string `json:"WarningTimeOut"`
		} `json:"RenewToken"`
	} `json:"d"`
}

// Error represents any error received from the endpoint
type Error struct {
	Err struct {
		Code    string `json:"code"`
		Message struct {
			Lang  string `json:"lang"`
			Value string `json:"value"`
		} `json:"message"`
		Innererror struct {
			Application struct {
				ComponentID      string `json:"component_id"`
				ServiceNamespace string `json:"service_namespace"`
				ServiceID        string `json:"service_id"`
				ServiceVersion   string `json:"service_version"`
			} `json:"application"`
			Transactionid   string `json:"transactionid"`
			Timestamp       string `json:"timestamp"`
			ErrorResolution struct {
				SAPTransaction string `json:"SAP_Transaction"`
				SAPNote        string `json:"SAP_Note"`
			} `json:"Error_Resolution"`
			Errordetails []struct {
				Code        string `json:"code"`
				Message     string `json:"message"`
				Propertyref string `json:"propertyref"`
				Severity    string `json:"severity"`
				Target      string `json:"target"`
			} `json:"errordetails"`
		} `json:"innererror"`
	} `json:"error"`
}

func (e Error) Error() string {
	return e.Err.Message.Value
}

const authURL string = baseURL + "/ZGW_UGYFELSZOLGALAT_LOGIN_SRV/Login?sap-client=201"
const tokenURL string = baseURL + "/ZGW_OAUTH_SRV/GetToken?sap-client=201"
const renewURL string = baseURL + "/ZGW_OAUTH_SRV/RenewToken?sap-client=201"

func getAuthCode(user User) (string, error) {
	var authCode string

	jsonValue, err := json.Marshal(user)
	if err != nil {
		return authCode, err
	}

	req, err := http.NewRequest(http.MethodPost, authURL, bytes.NewBuffer(jsonValue))
	if err != nil {
		return authCode, err
	}
	req.Header.Add("Accept", "application/json")
	req.Header.Add("X-Requested-With", "X")
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return authCode, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return authCode, err
	}

	if resp.StatusCode != 201 {
		var errorResp Error
		err = json.Unmarshal(body, &errorResp)
		if err != nil {
			return authCode, err
		}
		return authCode, errorResp
	}

	var authResp authResponse
	err = json.Unmarshal(body, &authResp)
	if err != nil {
		return authCode, err
	}

	authCode = authResp.D.AuthCode

	return authCode, nil
}

func getToken(authCode string) (oauth2.Token, error) {
	var token oauth2.Token

	req, err := http.NewRequest(http.MethodGet, tokenURL, nil)
	if err != nil {
		return token, err
	}
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add("Code", "'"+authCode+"'")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return token, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}

	if resp.StatusCode != 200 {
		var errorResp Error
		err = json.Unmarshal(body, &errorResp)
		if err != nil {
			return token, err
		}
		return token, errors.New(req.URL.RawQuery) //errorResp
	}

	var tokenResp tokenResponse
	err = json.Unmarshal(body, &tokenResp)
	if err != nil {
		return token, err
	}

	exp, err := strconv.Atoi(tokenResp.D.GetToken.ExpiresIn)
	if err != nil {
		return token, err
	}
	token = oauth2.Token{
		AccessToken:  tokenResp.D.GetToken.TokenCode,
		RefreshToken: tokenResp.D.GetToken.RenewToken,
		Expiry:       time.Now().Add(time.Duration(exp) * time.Second),
	}

	return token, nil
}

func renewToken(oldToken oauth2.Token) (oauth2.Token, error) {
	var token oauth2.Token

	req, err := http.NewRequest(http.MethodGet, renewURL, nil)
	if err != nil {
		return token, err
	}
	req.Header.Add("Accept", "application/json")

	q := req.URL.Query()
	q.Add("RefreshToken", "'"+oldToken.RefreshToken+"'")
	req.URL.RawQuery = q.Encode()

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return token, err
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return token, err
	}

	if resp.StatusCode != 200 {
		var errorResp Error
		err = json.Unmarshal(body, &errorResp)
		if err != nil {
			return token, err
		}
		return token, errorResp
	}

	var renewResp renewResponse
	err = json.Unmarshal(body, &renewResp)
	if err != nil {
		return token, err
	}

	exp, err := strconv.Atoi(renewResp.D.RenewToken.ExpiresIn)
	if err != nil {
		return token, err
	}
	token = oauth2.Token{
		AccessToken:  renewResp.D.RenewToken.TokenCode,
		RefreshToken: renewResp.D.RenewToken.RenewToken,
		Expiry:       time.Now().Add(time.Duration(exp) * time.Second),
	}

	return token, nil
}

type elmuemaszTokenSource struct {
	token *oauth2.Token
}

func (s *elmuemaszTokenSource) Token() (*oauth2.Token, error) {
	token, err := renewToken(*s.token)
	if err != nil {
		return s.token, err
	}
	s.token = &token
	return s.token, nil
}

// NewTokenSource creates a new oauth2 tken source from user credetials
func NewTokenSource(user User) (oauth2.TokenSource, error) {
	var tokenSource elmuemaszTokenSource

	ac, err := getAuthCode(user)
	if err != nil {
		return &tokenSource, err
	}

	token, err := getToken(ac)
	if err != nil {
		return &tokenSource, err
	}

	tokenSource = elmuemaszTokenSource{token: &token}

	return &tokenSource, nil
}
