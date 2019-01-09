package main

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

//ex. 2017v001
const regTag = "^[0-9]{4}[v]([0-9]{3})$"

//ex. 2017v001
const regTagDisable = "^[0-9]{4}[v]([0-9]{3})-Disable$"

type ImageInfo struct {
	Name   string `form:"name" json:"name" binding:"required"`
	Label  string `form:"label" json:"label" binding:"required"`
	Digest string `form:"digest" json:"digest" binding:"required"`
}

type ImageVo struct {
	Images []ImageInfo `form:"images" json:"images" binding:"required"`
}

type ReposVo struct {
	Repositories []string `form:"repositories" json:"repositories" binding:"required"`
}

type TagsVo struct {
	Name string   `form:"name" json:"name" name:"required"`
	Tags []string `form:"tags" json:"tags" binding:"required"`
}

type DigestVo struct {
	SchemaVersion int      `form:"schemaVersion" json:"schemaVersion" name:"required"`
	MediaType     string   `form:"mediaType" json:"mediaType" name:"required"`
	Config        *Config  `form:"config" json:"config" name:"required"`
	Layers        []Config `form:"layers" json:"layers" name:"required"`
}
type Config struct {
	MediaType string `form:"mediaType" json:"mediaType" name:"required"`
	Size      int    `form:"size" json:"size" name:"required"`
	Digest    string `form:"digest" json:"digest" name:"required"`
}

var wg sync.WaitGroup

/**
 * @api {get} /kubeGpu/images Get Images
 * @apiName GetImages
 * @apiGroup Image
 *
 * @apiDescription Get Container status by machineId
 *
 *
 * @apiSuccessExample Success-Response:
 * HTTP/1.1 200 OK
 * {"images":["torch:201707v001","caffe:201707v001","tensorflow:201707v001","simple:201706v001","all:201706v001","all_java:201706v001"]}
 * @apiError Unauthorized
 * @apiError MethodNotAllowed  Not allowed this method
 *
 * @apiErrorExample {json} Error-Response:
 *     HTTP/1.1 401 Unauthorized
 *     {
 *       "message": "Mapping repository is empty"
 *     }
 */
func getImagesListHandler(c *gin.Context) {
	var imageInfos []ImageInfo
	now := time.Now()
	nowStr := now.Format(configPara.FormatStr2)
	f, logErr := os.OpenFile(logPath+"/"+nowStr+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if logErr != nil {
		log.Errorf("error opening file: %v", logErr)
	}
	defer f.Close()

	log.Info("===================")
	log.Info("GET All Image")
	wg.Add(len(configPara.RegistryList))
	for _, repo := range configPara.RegistryList {
		log.Info("repo : " + repo)
		go func(c *gin.Context, repo string) {
			defer wg.Done()
			//Get repo tags
			var tagsVo TagsVo
			//curl -k -sS https://100.86.2.10:32190/v2/{repo}/tags/list
			tagsUrl := urlSslHeader + configPara.RegistryDnsName + configPara.RegistryIntPort + "/" + configPara.RegistryUrlVer + "/" + repo + configPara.RegistryUrlTags
			httpGetTags := imgApiGet(tagsUrl)
			err := json.Unmarshal([]byte(httpGetTags.Context), &tagsVo)
			if err != nil {
				log.Errorf("JsonStr fail : ", httpGetTags.Context)
				showErrorMsg(c, "Json formate error ", http.StatusUnauthorized)
				return
			}

			if len(tagsVo.Tags) == 0 {
				log.Info("tags zero")
			} else {
				sort.Strings(tagsVo.Tags)
				tag := tagsVo.Tags[len(tagsVo.Tags)-1]
				isMapping, _ := regexp.MatchString(regTagDisable, tag)
				if isMapping {
					//Last Version is Disable
					log.Info("lastTag : " + tag)
					var imageInfo ImageInfo
					imageInfo.Name = repo
					imageInfo.Label = tag
					imageInfo.Digest = getDegist(repo, tag)
					imageInfos = append(imageInfos, imageInfo)

				} else {
					for _, tag := range tagsVo.Tags {
						isMapping, _ := regexp.MatchString(regTag, tag)
						if isMapping {
							var imageInfo ImageInfo
							imageInfo.Name = repo
							imageInfo.Label = tag
							imageInfo.Digest = getDegist(repo, tag)
							imageInfos = append(imageInfos, imageInfo)
						}
					}
				}
			}
			log.Info("Size: " + strconv.Itoa(len(imageInfos)))
		}(c, repo)
	}
	wg.Wait()
	sort.Slice(imageInfos, func(i, j int) bool {
		switch strings.Compare(imageInfos[i].Name, imageInfos[j].Name) {
		case -1:
			return true
		case 1:
			return false
		}
		return imageInfos[i].Name > imageInfos[j].Name
	})
	resp := ImageVo{imageInfos}
	c.JSON(http.StatusOK, resp)
}

/**
 * @api {get} /kubeGpu/image/:repository Get Image
 * @apiName GetImage
 * @apiGroup Image
 *
 * @apiDescription Get Image status by repository
 *
 * @apiSuccessExample Success-Response:
 * HTTP/1.1 200 OK
 * {"images": [{"Name": "caffe:2017v002","digest": "cc56ee8818668845b61de05af42196943a51df2c27b11e9d8d27e59f94d3b485"}]}
 * @apiError Unauthorized
 * @apiError MethodNotAllowed  Not allowed this method
 *
 * @apiErrorExample {json} Error-Response:
 *     HTTP/1.1 401 Unauthorized
 *     {
 *       "message": "Tags is empty"
 *     }
 */
func getImageHandler(c *gin.Context) {
	var imageInfos []ImageInfo

	now := time.Now()
	nowStr := now.Format(configPara.FormatStr2)
	f, logErr := os.OpenFile(logPath+"/"+nowStr+".log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if logErr != nil {
		log.Errorf("error opening file: %v", logErr)
	}
	defer f.Close()

	log.Info("===================")
	log.Info("GET Single Image")

	//Get Image Repository
	repo := c.Param("repository")

	//Get
	//curl -k -sS https://100.86.2.10:32190/v2/{Image}/tags/list
	var tagsVo TagsVo
	tagsUrl := urlSslHeader + configPara.RegistryDnsName + configPara.RegistryIntPort + "/" + configPara.RegistryUrlVer + "/" + repo + configPara.RegistryUrlTags
	log.Info("tagsUrl : " + tagsUrl)
	httpGetTags := imgApiGet(tagsUrl)
	log.Info("Context : " + httpGetTags.Context)
	err := json.Unmarshal([]byte(httpGetTags.Context), &tagsVo)
	if err != nil {
		log.Errorf("JsonStr fail : ", httpGetTags.Context)
		showErrorMsg(c, "Json formate error ", http.StatusUnauthorized)
		return
	}

	sort.Strings(tagsVo.Tags)
	lastTag := tagsVo.Tags[len(tagsVo.Tags)-1]
	log.Info("lastTag : " + lastTag)
	isMapping, _ := regexp.MatchString(regTagDisable, lastTag)
	if isMapping {
		var imageInfo ImageInfo
		imageInfo.Name = repo
		imageInfo.Label = lastTag
		imageInfo.Digest = getDegist(repo, lastTag)
		imageInfos = append(imageInfos, imageInfo)
	} else {
		for t := 0; t < len(tagsVo.Tags); t++ {
			tag := tagsVo.Tags[t]
			log.Info("tag : " + tag)
			isMapping, _ := regexp.MatchString(regTag, tag)
			if isMapping {
				var imageInfo ImageInfo
				imageInfo.Name = repo
				imageInfo.Label = tag
				imageInfo.Digest = getDegist(repo, tag)
				imageInfos = append(imageInfos, imageInfo)
			}
		}
	}

	if len(tagsVo.Tags) == 0 {
		log.Errorf("Tags is empty ")
		showErrorMsg(c, "Tags is empty", http.StatusUnauthorized)
		return
	}
	resp := ImageVo{imageInfos}
	c.JSON(http.StatusOK, resp)
}

func imgApiGet(apiUrl string) (httpResp HttpResp) {
	log.Info("start imgApiGet")
	log.Info("apiUrl:" + apiUrl)

	//Ca Setting Insecure
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	//Get
	req, err := http.NewRequest("GET", apiUrl, nil)
	if err != nil {
		log.Errorf("NewRequest error", err)
		return HttpResp{http.StatusUnauthorized, "call image fail"}
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("client.Do fail", err)
		return HttpResp{http.StatusUnauthorized, "call image fail"}
	}
	defer resp.Body.Close()

	context := convertBody2Str(resp)
	return HttpResp{resp.StatusCode, context}
}

func imgApiGetHeader(apiUrl string, headerKey string, headerValue string) (httpResp HttpResp) {
	log.Info("start imgApiGetHeader")
	log.Info("apiUrl:" + apiUrl)

	//Ca Setting Insecure
	tlsConfig := &tls.Config{
		InsecureSkipVerify: true,
	}
	transport := &http.Transport{TLSClientConfig: tlsConfig}
	client := &http.Client{Transport: transport}

	//Get
	req, err := http.NewRequest("GET", apiUrl, nil)
	log.Info(headerKey)
	log.Info(headerValue)
	req.Header.Set(headerKey, headerValue)
	if err != nil {
		log.Errorf("NewRequest error", err)
		return HttpResp{http.StatusUnauthorized, "call image fail"}
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("client.Do fail", err)
		return HttpResp{http.StatusUnauthorized, "call image fail"}
	}
	defer resp.Body.Close()
	context := convertBody2Str(resp)

	return HttpResp{resp.StatusCode, context}
}

func getDegist(repo string, tag string) (resp string) {
	log.Info("getDegist")
	//curl -k -sS --header "Accept: application/vnd.docker.distribution.manifest.v2+json"
	//-X GET https://100.86.2.10:32190/v2/{repo}/manifests/{tag}
	tagsUrl := urlSslHeader + configPara.RegistryDnsName + configPara.RegistryIntPort + "/" + configPara.RegistryUrlVer + "/" + repo + configPara.RegistryUrlManifests + "/" + tag
	log.Info("tagsUrl :" + tagsUrl)
	httpGetDegist := imgApiGetHeader(tagsUrl, "Accept", "application/vnd.docker.distribution.manifest.v2+json")
	digestVo := &DigestVo{
		Config: &Config{},
	}
	err := json.Unmarshal([]byte(httpGetDegist.Context), digestVo)
	if err != nil {
		log.Errorf("JsonStr fail : ", httpGetDegist.Context)
		//showErrorMsg(c, "Json formate error ", http.StatusUnauthorized)
		return
	}
	log.Info("Raw Digest : " + digestVo.Config.Digest)
	digestAry := strings.Split(digestVo.Config.Digest, ":")
	digest := digestAry[len(digestAry)-1]
	log.Info("Digest : " + digest)
	return digest
}
