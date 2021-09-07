package kapis3_1_1

import (
	"errors"
	"fmt"
	"github.com/qiuax/kubesphere_sdk/lib"
	"io/ioutil"
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
	return nil
}

type DeleteApplicationReq struct {
	Workspaces   string `json:"workspaces"`
	Cluster      string `json:"cluster"`
	Namespaces   string `json:"namespaces"`
	AppClusterID string `json:"app_cluster_id"`
}

func (ks *KSInfo) DeleteApplication(req *DeleteApplicationReq) error {
	var endpointURL = ks.URL
	if req.Cluster == "default" {
		p := fmt.Sprintf("/kapis/openpitrix.io/v1/workspaces/%s/namespaces/%s/applications/%s", req.Workspaces, req.Namespaces, req.AppClusterID)
		endpointURL = endpointURL + p
	} else {

		p := fmt.Sprintf("/kapis/openpitrix.io/v1/workspaces/%s/clusters/%s/namespaces/%s/applications/%s", req.Workspaces, req.Cluster, req.Namespaces, req.AppClusterID)
		endpointURL = endpointURL + p

	}

	r := new(lib.Request)
	r.SetURL(endpointURL)
	r.SetHeader("Content-Type", "application/x-www-form-urlencoded")
	r.SetHeader("Authorization", " Bearer "+ks.Token)
	resp, err := r.DELETE()
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		b, _ := ioutil.ReadAll(resp.Body)
		return errors.New(fmt.Sprintf("err message: %s", string(b)))
	}
	return nil
}

type GetRunningApplicationReq struct {
	Workspaces string
	Namespaces string
	Page       string
	PageSize   string
}

type GetRunningApplicationResp struct {
	Items      []*RunningApplicationData `json:"items"`
	TotalCount int                       `json:"total_count"`
}

type Cluster struct {
	AppID      string `json:"app_id"`
	ClusterID  string `json:"cluster_id"`
	CreateTime string `json:"create_time"`
	Name       string `json:"name"`
	Owner      string `json:"owner"`
	RuntimeID  string `json:"runtime_id"`
	Status     string `json:"status"`
	StatusTime string `json:"status_time"`
	VersionID  string `json:"version_id"`
	Zone       string `json:"zone"`
}
type Version struct {
	AppID     string `json:"app_id"`
	Name      string `json:"name"`
	VersionID string `json:"version_id"`
}
type App struct {
	AppID       string      `json:"app_id"`
	CategorySet interface{} `json:"category_set"`
	ChartName   string      `json:"chart_name"`
	Name        string      `json:"name"`
}
type RunningApplicationData struct {
	Name    string   `json:"name"`
	Cluster *Cluster `json:"cluster"`
	Version *Version `json:"version"`
	App     *App     `json:"app"`
}

// GetRunningApplication 正在运行中的 http://192.168.1.177:30880
func (ks *KSInfo) GetRunningApplication(req *GetRunningApplicationReq) (*GetRunningApplicationResp, error) {
	var result *GetRunningApplicationResp
	endpointURL := ks.URL
	pageSize := req.PageSize
	page := req.Page

	if pageSize == "" {
		pageSize = "100"
	}
	if page == "" {
		page = "1"
	}
	p := fmt.Sprintf("/kapis/openpitrix.io/v1/workspaces/%s/namespaces/%s/applications?limit=%s&page=%s", req.Workspaces, req.Namespaces, pageSize, page)
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
