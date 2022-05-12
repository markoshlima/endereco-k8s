# endereco-k8s

**Pre-Requisites** <br />
GO 1.15 <br />
Docker 20.10.7

**Build** <br />
go build .

**Start** <br />
go run main.go

**Build docker container** <br />
docker build . -t markoshlima/endereco-k8s<br />
docker build . -t markoshlima/endereco-k8s --platform linux/amd64 //EKS

**Start docker container** <br />
docker run -it -p 5001:8080 -d markoshlima/endereco-k8s

**Logs docker container** <br />
docker logs [img]

**Inspect docker image** <br />
docker inspect [img] <br />
docker exec -it [img] /bin/bash

**Stop docker container** <br />
docker stop [img]