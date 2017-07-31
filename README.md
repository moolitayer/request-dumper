# request-dumper
Simple GO webserver dumping requests to STDOUT and response

```
CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main . # Compile with deps
sudo docker build -t request-dumper . # Build image
sudo docker run -d -p 9099:9099 --name dumper --net=host request-dumper # run

sudo docker login docker.io
sudo docker tag request-dumper mtayer/request-dumper
sudo docker tag mtayer/request-dumper docker.io/mtayer/request-dumper
sudo docker push docker.io/mtayer/request-dumper
```
