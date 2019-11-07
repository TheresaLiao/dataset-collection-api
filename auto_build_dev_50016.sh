container_name="task5-4-6-th"
image_name="task5-4-6-th:191107-th"
port1="50016"


docker stop $container_name
docker rm -f $container_name
docker build -t $image_name -f docker/Dockerfile .
docker run --name $container_name \
           -ti -d \
           --network=datasetbridge \
           -p $port1:80 \
           -v /home/ccma/dataset_doc_dev:/tmp \
           $image_name

docker rmi -f $(docker images | grep "<none>")
sudo docker ps -a | grep Exit | cut -d ' ' -f 1 | xargs sudo docker rm
