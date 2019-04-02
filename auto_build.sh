docker stop task5-4
docker rm task5-4
docker build -t golang_1.11.2:1.0.0 -f docker/Dockerfile .
docker run --name task5-4 \
           -ti -d \
           --network=datasetbridge \
           -p 50010:22 -p 50011:80 \
           -v /home/ccma/dataset_doc:/tmp \
           golang_1.11.2:1.0.0
