package kapis3_1_1

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (k *KSInfo) GetTokenInfo() (*OauthTokenResp, error) {
	var body *OauthTokenResp
	u := k.URL + "/oauth/token"
	req := new(Request)
	req.SetURL(u)

	//body :=
	req.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	req.SetBody(fmt.Sprintf("grant_type=password&username=%s&password=%s", k.Username, k.Password))
	resp, err := req.POST()
	if err != nil {
		return body, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("调用失败,返回的http状态码为%d", resp.StatusCode))
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	_ = json.Unmarshal(bodyByte, &body)
	return body, nil
}

// GetTokenInfo2 3.2版本
func (k *KSInfo) GetTokenInfo2() (*OauthTokenResp, error) {
	var body *OauthTokenResp
	u := k.URL + "/oauth/token"
	req := new(Request)
	req.SetURL(u)
	log.Println("url-------->", u)
	log.Println("body-------->", fmt.Sprintf("grant_type=password&username=%s&password=%s", k.Username, k.Password))

	//body :=
	req.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	req.SetBody(fmt.Sprintf("grant_type=password&username=%s&password=%s&client_id=kubesphere&client_secret=kubesphere", k.Username, k.Password))
	resp, err := req.POST()
	if err != nil {
		return body, err
	}
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("调用失败,返回的http状态码为%d", resp.StatusCode))
	}
	defer resp.Body.Close()
	bodyByte, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return body, err
	}
	_ = json.Unmarshal(bodyByte, &body)
	return body, nil
}
