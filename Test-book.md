# 環購 Test book
[TOC]

## Test API
* IP: 10.201.252.7 /10.174.61.1
* Port: 50015

### 1. Test Summary list
```shell=
## Read
curl localhost:50015/dataset/list
## responce
{
    "message": [
        {
            "title": "Car Type dataset",
            "desc": "Include all type of car video",
            "api": "/dataset/queryTrainTwOrg"
        },...
    ],
    "status": 200
}
```
### 2. Test detect
```shell=
## Read
curl localhost:50015/dataset/queryTrainYoloTag/0-7_nvNNdcM
## responce
{
    "data": [
        {
            "id": 4123975,
            "youtubeId": "0-7_nvNNdcM",
            "Object": "car",
            "filename": "res_00000001.jpg",
            "x_num": 165,
            "y_num": 292,
            "width": 106,
            "height": 40
        },...
    ],
    "status": 200
}
```

```shell=
## Read
curl localhost:50015/dataset/queryTrainLprTag/0WX9D_TR3HI
## responce
{
    "data": [
        {
            "id": 178998,
            "youtubeId": "0WX9D_TR3HI",
            "Object": "NS8680",
            "filename": "res_00000400.jpg",
            "x_num": 637,
            "y_num": 305,
            "width": 79,
            "height": 40
        },...
    ],
    "status": 200
}
```

```shell=
## Post
$ curl -X POST --data-binary "@/file_path" \
http://localhost:50015/filterfun/detectImg
## responce
{
    "filename": "filename",
    "tag": [
        {
            "confidences": [
                14
            ],
            "objectHeight": 21,
            "objectPicX": 71,
            "objectPicY": 90,
            "objectTypes": [
                "person"
            ],
            "objectWidth": 16
        },
        ...
    ]
}
```
![](https://i.imgur.com/8dkYPpo.png)

### 3. Test caracdnt list
```shell=
## Read
curl localhost:50015/dataset/caracdnt
## responce
{
    "message": {
        "title": "Car Accident dataset",
        "desc": "Include all type of Car Accident videos",
        "data": [
            {
                "id": 1,
                "keyWord": "擦撞+行車"
            },
            {
                "id": 2,
                "keyWord": "碰撞+監視器"
            },
            ...
        ]
    },
    "status": 200
}
```

```shell=
## Read
curl localhost:50015/dataset/caracdnt/1
## responce
{
    "message": [
        {
            "carAccidentID": "1606",
            "title": "1060614 疑擦撞釀行車糾紛 男子對公車司機動手拉扯",
            "youtubeId": "lk06YQv47hE",
            "url": "https://www.youtube.com/watch?v=lk06YQv47hE",
            "thumbnail": "https://i.ytimg.com/vi/lk06YQv47hE/mqdefault.jpg",
            "keyWord": "擦撞+行車",
            "collision_time": "0:0:25",
            "video_length": "1:49",
            "car_type": "car+car"
        },...
    ],
    "status": 200
}
```
### 4. Test subtitle list
```shell=
## Read
curl localhost:50015/dataset/subtitle
## responce
{
    "message": {
        "title": "Subtitle dataset",
        "desc": "Include all type of Subtitle videos",
        "data": [
            {
                "id": 1,
                "tagName": "外鄉女",
                "thumbnail": "https://i.ytimg.com/vi/B89bS2wdSAw/hqdefault.jpg?sqp=-oaymwEXCNACELwBSFryq4qpAwkIARUAAIhCGAE=&amp;rs=AOn4CLAtLjxT052Gov9XHrcAyDmMz4CtBw"
            },...
        ]
    },
    "status": 200
}
```

```shell=
## Read
curl localhost:50015/dataset/subtitle/1
## responce
{
    "message": {
        "title": "Subtitle dataset",
        "desc": "Include all type of Subtitle videos",
        "data": [
            {
                "id": 15,
                "title": "《日新蔭油》金鐘幸福好戲 外鄉女-愛人的選擇 Far And Away EP06",
                "url": "https://www.youtube.com/watch?v=-d8TlAGYFmc&list=PL02zpjjwMEjpx2-s14lxtNZY7rotQOj4_&index=7&t=0s",
                "thumbnail": "https://i.ytimg.com/vi/-d8TlAGYFmc/mqdefault.jpg",
                "srtUrl": "https://www.youtube.com/api/timedtext?v=-d8TlAGYFmc&lang=zh-TW"
            },...
        ]
    },
    "status": 200
}
```

```shell=
## download subtitle dataset
$ ls |grep subtitle |grep 4
## responce
subtitle_4
subtitle_4.tar.gz
## open browser
$  http://10.174.61.1:50015/dataset/youtubeUrl/subtitle/4

$ docker logs task5-4-5-TH
...
start Url2DownloadSubtitleTag
checkFileIsExist : /tmp/subtitle_4.tar.gz
/tmp/subtitle_4.tar.gz is exist
checkFileIsExist : /tmp/subtitle_4.tar.gz
/tmp/subtitle_4.tar.gz is exist
start respFile2Client
destFilePath : /tmp/subtitle_4.tar.gz
## Rename file 
## than untar
...
```
![](https://i.imgur.com/qkMzkGe.png)

```shell=
## download subtitle video
$ curl localhost:50015/dataset/subtitle/4
...
"id": 71,
"title": "新兵日記 Rookies' Diary Ep 29",
...

$ http://10.174.61.1:50015/dataset/youtubeUrl/subtitleById/71 
## open browser
## URL http://10.174.61.1:50015/dataset/youtubeUrl/subtitleById/71
## Rename file 
```
![](https://i.imgur.com/86Pg3oS.png)

![](https://i.imgur.com/wSWzjWV.png)


```shell=
curl localhost:50015/dataset/getSubTitleThumbnail
```

### 5. Test carType
* save **yolo** file
```shell=
## download file to client
## check
$ cd dataset_doc_dev_Theresa/traintworg/video
$ ls |grep vWhLkvyqR2U |grep yolo
## responce
...
vWhLkvyqR2U_yolo
...
$ cd vWhLkvyqR2U_yolo
$ ls |grep res_00003264
## Post download
$ curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"filename":"res_00003264.jpg","youtubeId":"vWhLkvyqR2U"}' \
    http://10.174.61.1:50015/filterfun/getYoloImg \
    --output res_00003264.jpg
```
![](https://i.imgur.com/r1A6172.jpg)

* save **lpr** file
```shell=
$ cd dataset_doc_dev_Theresa/traintworg/video
$ ls |grep vWhLkvyqR2U |grep lpr
## responce
...
vWhLkvyqR2U_lpr
...
$ cd vWhLkvyqR2U_lpr
$ ls |grep res_00003392
## Post download
$ curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"filename":"res_00003392.jpg","youtubeId":"vWhLkvyqR2U"}' \
    http://10.174.61.1:50015/filterfun/getLprImg \
    --output res_00003392.jpg
```
![](https://i.imgur.com/odQQrni.jpg)

* save video file by web
```shell=
$ cd dataset_doc_dev_Theresa/traintworg/video
$ ls |grep WrT0AE4gcS8 |grep mp4
## responce
....
WrT0AE4gcS8.mp4
...
## open browser
## URL http://10.174.61.1:50015/dataset/youtubeUrl/cartype/WrT0AE4gcS8 
## Rename file 
```
![](https://i.imgur.com/M1GNrIZ.png)
![](https://i.imgur.com/QDZBuFm.png)

```shell=
## get table : car_accident, col : keyWord list
## use keyword to search youtube list
## check new youtube list exit in train_tw_org
## download video into dataset_doc_dev_Theresa/traintworg/video
## insert one data into train_tw_org
$ curl http://localhost:50015/filterfun/youtubeUrl/getSearchByKeyWord
```

```shell=
## download & detect
## SELECT "URL" FROM train_tw_org WHERE "youtube_id" = 'NULL'
## downlaod URLs video into dataset_doc_dev_Theresa/traintworg/video
## update table train_tw_org Id by url
## triggerYoloApi & triggerLprApi
## 
$ curl http://localhost:50015/filterfun/url2DownloadTrainTwOrg
```

### 6. Test others video
```shell=
## Read
$ curl localhost:50015/dataset/queryTrainTwOrg
## responce
{
    "message": {
        "title": "Car Type dataset",
        "desc": "Include all type of car video",
        "data": [
            {
                "carAccidentID": "662",
                "title": "20160729中国交通事故合集：行车记录仪监控实拍下恐怖的最新车祸瞬间现场视频，国内车祸集锦斑马线礼让女司机高速闯红灯别车碰瓷鬼探头卡车重卡大货车。超清版 0002",
                "youtubeId": "0-7_nvNNdcM",
                "url": "https://www.youtube.com/watch?v=0-7_nvNNdcM",
                "thumbnail": "https://i.ytimg.com/vi/0-7_nvNNdcM/mqdefault.jpg",
                "keyWord": "",
                "collision_time": "",
                "video_length": "",
                "car_type": ""
            },...
        ]
    },
    "status": 200
}
```

```shell=
## Read
$ curl localhost:50015/dataset/queryTrainTwOrg/getThumbnail
## responce
{
    "message": {
        "1qzKYGAEw7c": "",
        "4y8Qaxwpmqw": "",...
    },
    "status": 200
}

## download video from youtube
$ curl --header "Content-Type: application/json" \
    --request POST \
    --data '{"filename":"S6elro0Wzo4","url":"https://www.youtube.com/watch?v=S6elro0Wzo4"}' \
    http://10.174.61.1:50015/filterfun/youtubeUrl \
    --output S6elro0Wzo4.mp4
$ ls
S6elro0Wzo4.mp4
```



### TODO
```shell=
## $ curl http://localhost:50015/filterfun/youtubeInfo/0WX9D_TR3HI

## get yolo resault
## parsing JSON into database
## INSERT INTO train_yolo_tag("youtube_id","x_num","y_num","width","height","object","filename") 
$ curl http://localhost:50015/filterfun/parsingTrainYoloResult

## get lpr resault
## parsing JSON into database
## INSERT INTO train_lpr_tag("youtube_id","x_num","y_num","width","height","plateNumber","filename") 
$ curl http://localhost:50015/filterfun/parsingTrainLprResult
```



###### tags: `環購`
