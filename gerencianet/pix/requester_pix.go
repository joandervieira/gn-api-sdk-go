package pix

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

type Requester struct {
	Auth interface {
		getAccessToken() (authResponseBody, error)
	}
	Url       string
	Timeout   int
	Token     string
	TokenDue  time.Time
	NetClient interface {
		Do(req *http.Request) (*http.Response, error)
	}
}

func NewRequester(clientID string, clientSecret string, CA string, Key string, sandbox bool, timeout int) *Requester {
	auth := NewAuth(clientID, clientSecret, CA, Key, sandbox, timeout)
	var cert, _ = tls.LoadX509KeyPair(CA, Key)

	var netTransport = &http.Transport{
		TLSClientConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}
	httpClient := &http.Client{Timeout: time.Second * time.Duration(timeout), Transport: netTransport}
	var gnURL string
	if sandbox {
		gnURL = UrlSandbox
	} else {
		gnURL = UrlProduction
	}
	return &Requester{*auth, gnURL, timeout, "", time.Time{}, httpClient}
}

func Authenticate(requester *Requester) (bool, error) {
	if requester.Token == "" || requester.TokenDue.Before(time.Now()) {
		tokenData, authErr := requester.Auth.getAccessToken()
		if authErr != nil {
			return false, authErr
		}
		requester.Token = tokenData.AccessToken
		requester.TokenDue = time.Now().Local().Add(time.Second * time.Duration(tokenData.ExpiresIn))
	}
	return true, nil
}

func (requester Requester) Request(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}) (string, error) {
	return requester.RequestWithHeaders(endpoint, httpVerb, requestParams, body, nil)
}

func (requester Requester) RequestWithHeaders(endpoint string, httpVerb string, requestParams map[string]string, body map[string]interface{}, headers map[string]string) (string, error) {
	requestBody := new(bytes.Buffer)
	json.NewEncoder(requestBody).Encode(body)

	_, authErr := Authenticate(&requester)
	if authErr != nil {
		_, authErr = Authenticate(&requester)
	}
	if authErr != nil {
		return "", authErr
	}

	route := GetRoute(endpoint, requestParams)
	route += GetQueryString(requestParams)
	req, _ := http.NewRequest(httpVerb, requester.Url+route, requestBody)

	if (httpVerb == "POST" || httpVerb == "PUT") && body != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("accept", "application/json")
	req.Header.Add("api-sdk", "go-"+Version)
	req.Header.Add("Authorization", "Bearer "+requester.Token)
	for k, v := range headers {
		req.Header.Add(k, v)
	}
	res, resErr := requester.NetClient.Do(req)

	if resErr != nil {
		return "", resErr
	}

	defer res.Body.Close()

	reqResp, _ := ioutil.ReadAll(res.Body)
	response := string(reqResp)

	if res.StatusCode != http.StatusOK && res.StatusCode != http.StatusCreated {
		return "", errors.New(response)
	}

	return response, nil
}

func GetRoute(endpoint string, params map[string]string) string {
	patern, _ := regexp.Compile("\\:(\\w+)")
	variables := patern.FindAllStringSubmatch(endpoint, -1)
	for i := 0; i < len(variables); i++ {
		if value, exists := params[variables[i][1]]; exists {
			endpoint = strings.Replace(endpoint, variables[i][0], value, -1)
			delete(params, variables[i][1])
		}
	}
	return endpoint
}

func GetQueryString(params map[string]string) string {
	var query string
	for key, value := range params {
		if query != "" {
			query += "&"
		} else {
			query += "?"
		}
		query += key + "=" + url.QueryEscape(value)
	}

	return query
}
