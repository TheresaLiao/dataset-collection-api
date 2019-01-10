package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AccountVo struct {
	Account string `form:"account" json:"account" binding:"required"`
}

func createVolumnHandler(c *gin.Context) {
	log.Info("===================")
	log.Info("POST")
	var accountVo AccountVo
	c.BindJSON(&accountVo)
	log.Info("account:" + accountVo.Account)
	if accountVo.Account == "" {
		log.Errorf("Account :" + accountVo.Account)
		showErrorMsg(c, "Missing parameter error", http.StatusUnauthorized)
		return
	}
	//Create NFS user
	log.Info("Create NFS user")
	urlStr := urlHeader + configPara.NfsHost + configPara.NfsPort + configPara.NfsUrl
	log.Info("urlStr : " + urlStr)
	httpPostRespNfs := nfsApiPost(accountVo.Account, urlStr)
	log.Info(strconv.Itoa(httpPostRespNfs.StatusCode))
	if httpPostRespNfs.StatusCode != 200 {
		log.Errorf("Create nfs error : " + accountVo.Account)
		log.Errorf(httpPostRespNfs.Context)
		c.String(httpPostRespNfs.StatusCode, httpPostRespNfs.Context)
		return
	}
	c.String(http.StatusOK, "")
	return
}

func nfsApiPost(acccount string, apiUrl string) (httpResp HttpResp) {
	log.Info("start nfsApiPost")
	log.Info("apiUrl:" + apiUrl)
	log.Info("acccount :" + acccount)

	accountVo := AccountVo{acccount}
	accountByte, err := json.Marshal(accountVo)
	if err != nil {
		log.Errorf("Error formate error ", err)
		return HttpResp{http.StatusUnauthorized, "call storage fail"}
	}
	body := bytes.NewReader(accountByte)

	req, err := http.NewRequest("POST", apiUrl, body)
	if err != nil {
		log.Errorf("Error NewRequest", err)
		return HttpResp{http.StatusUnauthorized, "call storage fail"}
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf("DefaultClient fail", err)
		return HttpResp{http.StatusUnauthorized, "call storage fail"}
	}
	defer resp.Body.Close()

	context := convertBody2Str(resp)
	return HttpResp{resp.StatusCode, context}
}
