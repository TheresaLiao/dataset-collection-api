## Swagger api file
### auto gen swagger json file
ref : https://github.com/swaggo/swag
```shell=
$ cd ~/dataset-collection-api
$ sh swagger/swaggo/auto_build.sh

## will show swagger json file
├── src
│    └── api
│        ├── docs
│        │    ├── docs.go
│        │    ├── swagger.json
│        │    └── swagger.yaml
```

### swagger-editor show by gui
ref : https://github.com/codeasashu/swagger-editor/tree/envvars
```shell=
## create new image
$ cd ~/dataset-collection-api/swagger/
$ git clone https://github.com/codeasashu/swagger-editor/tree/master
$ cd swagger-editor
$ git checkout envvars
$ docker build -t swaggerapi/swagger-editor:envvars .

## mount json file
$ cd ~/dataset-collection-api/src/api/docs
$ docker run -d -p 80:8080 \
    -v $(pwd):/tmp \
    -e SWAGGER_FILE=/tmp/swagger.json \
    --name swagger-editor \
    swaggerapi/swagger-editor:envvars

## use browser
http://localhost:80
```
![](https://i.imgur.com/iKDW7ZZ.png)
