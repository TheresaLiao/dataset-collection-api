# dataset-collection-api

## API
![api](picture/api.png "API")

## API Lifecycle
![api](picture/dataset_lifecycle.png "API Lifecycle")

## Quick start
```shell=
# create  network 
docker network create -d bridge datasetbridge

docker build -t golang_1.11.2:1.0.0 -f docker/Dockerfile .
docker run --name task5-4 -ti -d  --network=datasetbridge -p 50010:22 -p 50011:80 golang_1.11.2:1.0.0
```

## download file from url api example
```shell=
//call by url
$ curl --header "Content-Type: application/json" \
       --request POST \
       --data '{"filename":"test","url":"https://www.youtube.com/watch?v=JpcTvrSdBoE"}' \
       http://localhost:50011/filterfun/youtubeUrl \
       --output test.mp4

//call by caracdnt id
$ curl --header "Content-Type: application/json" \
       --request GET \
       http://localhost:50011/filterfun/youtubeUrl/caracdnt/11 \
       --output 11.tar.gz

//call by subtitle id
$ curl --header "Content-Type: application/json" \
       --request GET \
       http://localhost:50011/filterfun/youtubeUrl/subtitle/4 \
       --output 11.tar.gz
```
