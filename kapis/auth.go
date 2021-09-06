package kapis

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"kubesphere_sdk/lib"
	"net/http"
)

func (k *ksInfo) GetTokenInfo() (*OauthTokenResp, error) {
	var body *OauthTokenResp
	u := k.URL + "/oauth/token"
	req := new(lib.Request)
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
