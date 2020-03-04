IMAGE_NAME=swaggo-swag:v1.0
CONTAINER_NAME=gen_swag_file

docker stop $CONTAINER_NAME
docker rm $CONTAINER_NAME

docker build -f ./swagger/swaggo/Dockerfile -t $IMAGE_NAME .
docker run -ti \
--name $CONTAINER_NAME \
-v $(pwd)/src/api:/app \
$IMAGE_NAME swag init
