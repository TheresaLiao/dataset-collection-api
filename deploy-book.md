# 環購 deploy book
[TOC]

## 1.Network
```shell=
sh 1-build-network.sh
# docker network create datasetbridge
```

## 2.DB
```shell=
sh 2-build-db.sh

## DB Server
#$ docker run -d -it --name $containername --network=datasetbridge -p 50010:5432 -v $data_org:$data_out -e POSTGRES_DB=Test_db -e POSTGRES_USER=admin -e POSTGRES_PASSWORD='12345' postgres:11.9

## DB UI , ignore
#$ docker run -d -it -p 50015:5050 thajeztah/pgadmin4 
```
## 3.Run Yolo & LPR & DEID API
### 1.YOLO API
* task5-4-1-TH:8080/yolo_coco_image : transfer video into image ,and detect yolo type object output JSON
```shell=
$ cd dataset-collection-api-engine-Theresa

## start yolo api
## task5-4-1-TH, Image=datacollection-detectyolo:1.0.0-TH
$ sh ./iclx_script/task5-4-1-TH_datacollection-detectyolo_1.0.0-TH.sh

## check task5-4-1-TH
$ docker ps
CONTAINER ID IMAGE COMMAND    CREATED    STATUS    PORTS    NAMES
... datacollection-detectyolo:1.0.0-TH "python python/run_a…" 2 days ago Up 2 days task5-4-1-TH

## Test file & call API
$ docker exec -ti task5-4-1-TH bash
$ apt install curl
$ cd /home/darknet_AlexeyAB
$ ls traintworg/video/ |grep mp4
...
Rf9MxTLfdik.mp4
...

$ curl -d '{"videonames":"./traintworg/video/Rf9MxTLfdik.mp4", "dirname": "./traintworg/video/Rf9MxTLfdik_yolo"}' \
    -H "Content-Type: application/json" \
    -X POST http://localhost:8080/yolo_coco_image
```
### 2.LPR API
* task5-4-2-TH:8080/yolo_lpr_image : transfer video into image ,and detect lpr object output JSON
```shell=
$ cd dataset-collection-api-engine-Theresa

## start lpr api
## task5-4-2-TH, Image=datacollection-detectlpr:1.0.0-TH
$ sh ./iclx_scripttask5-4-2-TH_datacollection-detectlpr_1.0.0-TH.sh

## check task5-4-2-TH
$ docker ps
CONTAINER ID IMAGE COMMAND    CREATED    STATUS    PORTS    NAMES
... datacollection-detectlpr:1.0.0-TH "python python/run_a…" 2 days ago Up 2 days task5-4-2-TH

## Test file & call API
$ docker exec -ti task5-4-2-TH bash
$ apt install curl
$ cd /home/darknet_AlexeyAB
$ ls traintworg/video/ |grep mp4
...
04jm7VfInbo.mp4
...

$ curl -d '{"videonames":"./traintworg/video/04jm7VfInbo.mp4", "dirname": "./traintworg/video/04jm7VfInbo_lpr"}' \
    -H "Content-Type: application/json" \
    -X POST http://localhost:8080/yolo_lpr_image
```
![](https://i.imgur.com/oe29xOL.png)


### 3.DEID API
* task5-4-3-TH:8080/yolo_lpr_image : for deidentify with lpr for car type data
```shell=
$ cd dataset-collection-api-engine-Theresa

## start yolo deid lpr api
## task5-4-3-TH, Image=datacollection-deidentlpr:1.0.0-TH
$ sh ./iclx_script/task5-4-3-TH-datacollection-deidentlpr_1.0.0-TH.sh

## check task5-4-2-TH
$ docker ps
CONTAINER ID IMAGE COMMAND    CREATED    STATUS    PORTS    NAMES
... datacollection-deidentlpr:1.0.0-TH "python python/run_a…" 2 days ago Up 2 days 8080/tcp task5-4-3-TH

## Test file & call API
$ docker exec -ti task5-4-3-TH bash
$ apt install curl
$ cd /home/darknet_AlexeyAB
$ ls ./traintworg/video/ |grep lpr
...
zvmK-fOoR8g_lpr
...

$ ls ./traintworg/video/zvmK-fOoR8g_lpr |grep jpg

$ curl -d '{"filename":"./traintworg/video/zvmK-fOoR8g_lpr/res_00002974.jpg"}' \
    -H "Content-Type: application/json" \
    -X POST http://localhost:8080/yolo_lpr_image
```
![](https://i.imgur.com/kazab6x.png)

### 4.Just YOL API , ignore it
* http://task5-4-8-TH:8080/yolo_coco_image : detect Img
```shell=
## deidentify lpr
## task5-4-8-TH, Image=detect_yolo:1.0.0-TH
```



## 4. Run API

```shell=
$ cd dataset-collection-api-Theresa
$ sh task5-4-4-TH_datacollection-api_1.0.0-TH.sh

$ docker ps
CONTAINER ID IMAGE COMMAND    CREATED    STATUS    PORTS    NAMES
... datacollection-api:1.0.0-TH "/bin/sh -c /app/main" 5 hours ago Up 5 hours 22/tcp, 0.0.0.0:50014->80/tcp task5-4-4-TH
```
* Test
```shell=
## yolo detect
$ curl -X POST --data-binary "@/file_path" \
http://localhost:50014/filterfun/detectImg

## download & detect
## 10.201.252.7:50014
## 10.174.61.1:50014
## http://10.201.252.7:50016/dataset/queryTrainTwOrg
$ curl http://localhost:50014/dataset/queryTrainTwOrg
$ curl http://localhost:50014/filterfun/url2DownloadTrainTwOrg

## get yolo resault
$ curl http://localhost:50014/filterfun/parsingTrainYoloResult
$ curl http://localhost:50014/dataset/queryTrainYoloTag/04jm7VfInbo
$ curl --request GET \
http://localhost:50014/filterfun/getYoloImg/0R85NEB8l64/res_00016914.jpg \
--output res_00016914.jpg

$ curl --request GET \
http://140.96.0.34:50014/filterfun/getYoloImg/C4rO3gowyxk/res_00000001.jpg \
--output res_00000001.jpg

## get lpr resault
$ curl http://localhost:50014/filterfun/parsingTrainLprResult
$ curl http://localhost:50014/dataset/queryTrainLprTag/0YzQL00_b30
$ curl --request GET \
http://140.96.0.34:50014/filterfun/getLprImg/0WX9D_TR3HI/res_00000401.jpg \
--output res_00000401.jpg
```
###### tags: `環購`


