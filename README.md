# sut-order-go
Microservice for managing user order

## How to run in local?

1. Create file env and name it `dev.env`. its content can be seen in code block below. 
```
PORT=:50052
PRODUCT_HOST=:50054
STORAGE_HOST=:50053
DB_URL=
REDIS_ADDRESS=
REDIS_PASSWORD=
REDIS_DBNUMBER=
```

2. Execute command below
```
make init # initialize go.mod
make tidy # Tidy up go module
```

3. Adding go bin into path env variables
```
export PATH=$PATH:$(go env GOPATH)/bin
```

4. Adding folder with `pb` as name into ther project root directory

5. Generate protobuf by executing command below
```
make proto-gen
```

6. Run the application
```
make run
```

## Using docker
```
docker build --tag=sut/order-service --build-arg SERVICE=sut-order-go --build-arg PORT=50052 .
docker run -p 50052:50052 <IMAGE_ID>
```
