container_name="swag-gen-api"
image_name="swag-gen-api:1.0"

docker stop $container_name
docker rm   $container_name

docker build -t $image_name -f ./swagger/swaggo/Dockerfile .
docker run -d -ti \
--name $container_name \
-v "$(pwd)"/src/api:/app \
$image_name swag init
