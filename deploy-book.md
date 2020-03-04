# 環購 deploy book
[TOC]

## Network
```shell=
docker network create datasetbridge
```

## DB
```shell=
$ docker run -d -it\
    --name Test_Postgres \
    --network=datasetbridge \
    -p 50003:5432 \
    -v /home/ccma/Test_Postgres:/var/lib/postgresql/data \
    -e POSTGRES_DB=Test_db \
    -e POSTGRES_USER=admin \
    -e POSTGRES_PASSWORD='12345' \
    postgres:latest 
 
$ docker run -d -it\
        -p 50015:5050 \
         thajeztah/pgadmin4 
```
## Run Yolo & LPR API
```shell=
$ cd yolo_with_lpr

# start yolo api
$ sh ./iclx_script/nvdocker_run_yolo_api_for_dataset.sh
$ curl -d '{"videonames":"./traintworg/video/Rf9MxTLfdik.mp4", "dirname": "./traintworg/video/Rf9MxTLfdik_yolo"}' \
    -H "Content-Type: application/json" \
    -X POST http://localhost:8080/yolo_coco_image
    
# start yolo deid lpr api
$ sh ./iclx_script/nvdocker_run_deident_api_for_dataset.sh
$ curl -d '{"filename":"./traintworg/video/456_yolo/res_00000032.jpg"}' \
    -H "Content-Type: application/json" \
    -X POST http://localhost:8080/yolo_lpr_image

# start lpr api
$ sh ./iclx_script/nvdocker_run_lpr_api_for_dataset.sh
$ curl -d '{"videonames":"./traintworg/video/04jm7VfInbo.mp4", "dirname": "./traintworg/video/04jm7VfInbo_lpr"}' \
    -H "Content-Type: application/json" \
    -X POST http://localhost:8080/yolo_lpr_image
```
## Run API
```shell=
$ cd dataset-collection-api
$ sh auto_build_dev.sh
```
## Test
```shell=
## yolo detect
$ curl -X POST \
--data-binary "@/file_path" \
http://10.174.61.1:50014/filterfun/detectImg


## download & detect
## 10.201.252.7:50014
## http://10.201.252.7:50016/dataset/queryTrainTwOrg
$ curl http://localhost:50016/dataset/queryTrainTwOrg
$ curl http://localhost:50016/filterfun/url2DownloadTrainTwOrg

## get yolo resault
$ curl http://localhost:50016/filterfun/parsingTrainYoloResult
$ curl http://localhost:50016/dataset/queryTrainYoloTag/04jm7VfInbo
$ curl --request GET \
http://localhost:50016/filterfun/getYoloImg/0R85NEB8l64/res_00016914.jpg \
--output res_00016914.jpg

$ curl --request GET \
http://140.96.0.34:50016/filterfun/getYoloImg/C4rO3gowyxk/res_00000001.jpg \
--output res_00000001.jpg

## get lpr resault
$ curl http://localhost:50016/filterfun/parsingTrainLprResult
$ curl http://localhost:50016/dataset/queryTrainLprTag/0YzQL00_b30
$ curl --request GET \
http://140.96.0.34:50016/filterfun/getLprImg/0WX9D_TR3HI/res_00000401.jpg \
--output res_00000401.jpg
```
###### tags: `環購`
