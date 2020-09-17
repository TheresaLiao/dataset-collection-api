containername="task5-4-0-TH"
image="postgres:11.9"

#data_org="/home/ccma/Test_Postgres"
data_org="/home/ccma/sdd/dataset-collection-api_Test_Postgres_Theresa"
data_out="/var/lib/postgresql/data"

host_port=50010
con_port=5432

docker stop $containername
docker rm $containername

docker run -d -it\
    --name $containername \
    --network=datasetbridge \
    -p $host_port:$con_port \
    -v $data_org:$data_out \
    -e POSTGRES_DB=Test_db \
    -e POSTGRES_USER=admin \
    -e POSTGRES_PASSWORD='12345' \
    $image
