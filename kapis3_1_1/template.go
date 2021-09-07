package kapis3_1_1

import (
	"errors"
	"fmt"
	"io/ioutil"
	"kubesphere_sdk/lib"
	"net/http"
)

type GetAppTemplatesReq struct {
	Page     string
	PageSize string
}

type AppTemplateData struct {
	AppID            string            `json:"app_id"`
	AppVersionTypes  string            `json:"app_version_types"`
	CategorySet      []*CategorySet    `json:"category_set"`
	CreateTime       string            `json:"create_time"`
	Description      string            `json:"description"`
	Icon             string            `json:"icon"`
	Isv              string            `json:"isv"`
	LatestAppVersion *LatestAppVersion `json:"latest_app_version"`
	Name             string            `json:"name"`
	Owner            string            `json:"owner"`
	Status           string            `json:"status"`
	StatusTime       string            `json:"status_time"`
	UpdateTime       string            `json:"update_time"`
	ClusterTotal     int               `json:"cluster_total"`
}
type CategorySet struct {
	CategoryID string `json:"category_id"`
	CreateTime string `json:"create_time"`
	Locale     string `json:"locale"`
	Name       string `json:"name"`
	Status     string `json:"status"`
}
type LatestAppVersion struct {
	Active      bool   `json:"active"`
	AppID       string `json:"app_id"`
	CreateTime  string `json:"create_time"`
	Description string `json:"description"`
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	PackageName string `json:"package_name"`
	Status      string `json:"status"`
	StatusTime  string `json:"status_time"`
	UpdateTime  string `json:"update_time"`
	VersionID   string `json:"version_id"`
}
type GetAppTemplatesResp struct {
	Items      []*AppTemplateData `json:"items"`
	TotalCount int                `json:"total_count"`
}

// GetActiveAppTemplates 获取已激活的应用
func (ks *KSInfo) GetActiveAppTemplates(req *GetAppTemplatesReq) (*GetAppTemplatesResp, error) {
	var result *GetAppTemplatesResp
	page := req.Page
	pageSize := req.PageSize
	if page == "" {
		page = "1"
	}
	if pageSize == "" {
		pageSize = "20"
	}
	endpointURL := ks.URL
	p := fmt.Sprintf("/kapis/openpitrix.io/v1/apps?limit=%s&page=%s&conditions=status=active,repo_id=repo-helm&reverse=true", pageSize, page)
	endpointURL = endpointURL + p

	r := new(lib.Request)
	r.SetURL(endpointURL)
	r.SetHeader("Authorization", " Bearer "+ks.Token)
	resp, err := r.GET()
	if err != nil {
		return result, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return result, errors.New(fmt.Sprintf("err message: %s", string(b)))
	}

	if err := resp.BindJSON(&result); err != nil {
		return result, err
	}
	return result, nil
}
