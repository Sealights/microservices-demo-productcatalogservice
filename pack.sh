docker build -t productcatalogservice .
docker tag productcatalogservice:latest 159616352881.dkr.ecr.eu-west-1.amazonaws.com/microservices-demo-productcatalogservice:latest
docker push 159616352881.dkr.ecr.eu-west-1.amazonaws.com/microservices-demo-productcatalogservice:latest
