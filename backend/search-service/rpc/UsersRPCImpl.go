package rpc

import (
	"encoding/json"
	"github.com/SKFE396/search-service/config"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type UsersRPCImpl struct {
	usersRPC UsersRPC
}

func (u *UsersRPCImpl) GetUserInfos(uids []int64) (result []UserInfo, err error) {
	var requestBody struct {
		Uids []string `json:"uids"`
	}
	requestBody.Uids = make([]string, len(uids))
	for i, v := range uids {
		requestBody.Uids[i] = strconv.FormatInt(v, 10)
	}
	var bodyBytes []byte
	bodyBytes, err = json.Marshal(requestBody)
	if err != nil {
		return
	}
	bodyReader := strings.NewReader(string(bodyBytes))
	var response *http.Response
	response, err = http.Post(config.UserServiceUrl+"infoList", "application/json", bodyReader)
	if err != nil {
		return
	}
	var responseBodyJson []byte
	responseBodyJson, err = ioutil.ReadAll(response.Body)
	var resultObj struct {
		Result []UserInfo `json:"result"`
	}
	err = json.Unmarshal(responseBodyJson, &resultObj)
	if err != nil {
		return
	}
	return resultObj.Result, nil
}

func (u *UsersRPCImpl) ParseToken(token string) (successful bool, uid int64, role int8) {
	request, err := http.NewRequest("GET", config.UserServiceUrl+"checkToken?token="+token, nil)
	if err == nil {
		request.Header.Set("Accept", "application/json")
		client := http.Client{}
		var response *http.Response
		response, err = client.Do(request)
		if err == nil {
			var responseBodyJson []byte
			responseBodyJson, err = ioutil.ReadAll(response.Body)
			var response struct {
				Successful bool  `json:"successful"`
				Uid        int64 `json:"uid"`
				Role       int8  `json:"role"`
			}
			err = json.Unmarshal(responseBodyJson, &response)
			if err == nil {
				return response.Successful, response.Uid, response.Role
			}
		}
	}
	log.Error("Failed to call checkToken RPC")
	return false, 0, 0
}