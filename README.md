# sut-order-go
Microservice for managing user order

## How to run in local?

1. Create file env and name it `dev.env`. its content can be seen in code block below. 
```
PORT=
PRODUCT_HOST=
STORAGE_HOST=
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
