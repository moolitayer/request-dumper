# request-dumper

CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main . # Compile with deps
sudo docker build -t request-dumper . # Build image
sudo docker run -d -p 9099:9099 --name dumper --net=host request-dumper # run
