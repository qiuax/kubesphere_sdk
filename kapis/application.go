package kapis

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func CreateApplication(conf string) {
	//u := "192.168.1.177:32517/kapis/openpitrix.io/v1/workspaces/testapi/clusters/admin/namespaces/default/applications"
	u := fmt.Sprintf("http://192.168.1.177:32517/kapis/openpitrix.io/v1/workspaces/testapi/namespaces/getapi/applications")
	m := map[string]string{
		"app_id":     "app-x3q3640n76rw47-store",
		"name":       "nginx-n9ueyu",
		"version_id": "appv-x795nnrnlx8qmy",
		//"conf":"replicaCount: 1\nimage:\n  html: {}\n  nginx:\n    repository: nginx\n    pullPolicy: IfNotPresent\nnameOverride: ''\nfullnameOverride: ''\nservice:\n  name: http\n  type: ClusterIP\n  port: 80\ningress:\n  enabled: false\n  annotations: {}\n  paths:\n    - /\n  hosts:\n    - nginx.local\n  tls: []\nextraVolumes: []\nextraVolumeMounts: []\nextraInitContainers: []\nreadinessProbe:\n  path: /\n  initialDelaySeconds: 5\n  periodSeconds: 3\n  failureThreshold: 6\nlivenessProbe:\n  path: /\n  initialDelaySeconds: 5\n  periodSeconds: 3\nresources: {}\nconfigurationFile: {}\nextraConfigurationFiles: {}\nnodeSelector: {}\ntolerations: []\naffinity: {}\ntests:\n  enabled: false\n",
		"conf": conf,
	}
	marshal, _ := json.Marshal(m)
	body := strings.NewReader(string(marshal))
	req, err := http.NewRequest(http.MethodPost, u, body)
	if err != nil {
		//logger.Error(logger.LOGGER_MSG, logger.String("kubesphere ListAllApplications newRequest失败,err ", err.Error()))
		log.Println("1------------------>", err.Error())
		return
	}
	tk := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImFkbWluIiwidG9rZW5fdHlwZSI6ImFjY2Vzc190b2tlbiIsImV4cCI6MTYzMDkxODQwMCwiaWF0IjoxNjMwOTExMjAwLCJpc3MiOiJrdWJlc3BoZXJlIiwibmJmIjoxNjMwOTExMjAwfQ.UUQZgkyCScZsKAW9i-wP1S_pxFIolHhyrFqEywpaVJw"
	//req.Header.Set("Authorization", tk)

	req.Header.Add("Authorization", " Bearer "+tk)
	req.Header.Set("Content-Type", "application/json")
	do, err := http.DefaultClient.Do(req)
	if err != nil {
		//logger.Error(logger.LOGGER_MSG, logger.String("kubesphere ListAllApplications Do 失败,err ", err.Error()))
		log.Println("2----------->", err.Error())
		return
	}

	if do.StatusCode != http.StatusOK {
		log.Println("3----------->", do.StatusCode)

	}
	result, err := ioutil.ReadAll(do.Body)
	if err != nil {
		log.Println("4----------->", err.Error())
		return
	}
	log.Println("============>", string(result))
}
