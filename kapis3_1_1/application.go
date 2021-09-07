package kapis3_1_1

import (
	"errors"
	"fmt"
	"io/ioutil"
	"kubesphere_sdk/lib"
	"net/http"
)

type CreateApplicationReq struct {
	Workspaces string `json:"workspaces"`
	Namespaces string `json:"namespaces"`
	AppID      string `json:"app_id"`
	Name       string `json:"name"`
	VersionID  string `json:"version_id"`
	Conf       string `json:"conf"`
}

func (ks *KSInfo) CreateApplication(request *CreateApplicationReq) error {
	u := fmt.Sprintf("%s/kapis/openpitrix.io/v1/workspaces/%s/namespaces/%s/applications", ks.URL, request.Workspaces, request.Namespaces)
	req := new(lib.Request)
	req.SetURL(u)
	body := &struct {
		AppID     string `json:"app_id"`
		Name      string `json:"name"`
		VersionID string `json:"version_id"`
		Conf      string `json:"conf"`
	}{
		AppID:     request.AppID,
		Name:      request.Name,
		VersionID: request.VersionID,
		Conf:      request.Conf,
	}
	req.SetHeader("Content-Type", "application/json")
	req.SetHeader("Authorization", " Bearer "+ks.Token)
	req.SetJSONBody(&body)
	resp, err := req.POST()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("err message: %s", string(b)))
	}
	//result := make(map[string]string)
	//if err := resp.BindJSON(&result); err != nil {
	//	return err
	//}
	//marshal, _ := json.Marshal(result)
	//log.Println("-------->",string(marshal))
	return nil
}
