container_name="task5-4-5-TH"
image_name="datacollection-api:1.0.0-TH"

data_org="/home/ccma/sdd/dataset_doc_dev_Theresa"
data_out="/tmp"

host_port="50015"

## remove old image & build image
docker stop $container_name
docker rm -f $container_name
docker build -t $image_name -f docker/Dockerfile .

## run container
docker run -ti -d --name $container_name \
                  --network=datasetbridge \
                  -p $host_port:80 \
                  -v $data_org:$data_out \
                  $image_name

## clean none image 
docker rmi -f $(docker images | grep "<none>")
sudo docker ps -a | grep Exit | cut -d ' ' -f 1 | xargs sudo docker rm
