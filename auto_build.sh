docker stop task5-4
docker rm -f task5-4
docker build -t golang_1.11.2:1.0.0 -f docker/Dockerfile .
docker run --name task5-4 \
           -ti -d \
           --network=datasetbridge \
           -p 50010:22 -p 50011:80 \
           -v /home/ccma/dataset_doc_dev:/tmp \
           golang_1.11.2:1.0.0

docker rmi -f $(docker images | grep "<none>")
sudo docker ps -a | grep Exit | cut -d ' ' -f 1 | xargs sudo docker rm
