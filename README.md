# mth-api using redis and mysql


### Database configuration, redis, port, etc in the config.yml file
### Install Libarary
- go get github.com/gomodule/redigo/redis
- go get gopkg.in/go-playground/validator.v9
- go get github.com/go-sql-driver/mysql
- go get github.com/dgrijalva/jwt-go

## Create Database mysql
create database mth

## Create Table mtr_product
```
create table mtr_product(
  productId int not null PRIMARY key AUTO_INCREMENT,
  sku varchar(30),
  productName varchar(100),
  productDesc text,
  price decimal(19,2),
  quantity int(10),
  createdDate timestamp,
  createdBy varchar(60),
  updatedDate timestamp,
  updatedBy varchar(60),
  UNIQUE(sku)
);
```
## Create Table mtr_user 
```
create table mtr_user(
userId int not null PRIMARY key AUTO_INCREMENT,
userName varchar(60),
userPassword varchar(150),
createdDate timestamp,
createdBy varchar(60),
updatedDate timestamp,
updatedBy varchar(60)
);
```

## Insert user
```
insert into mtr_user(userName,userPassword,createdDate,createdBy)
VALUES('test',sha2('test123',256),now(),'system')
```

# GetToken (token is used as authentication for api requests)
## Request Token 
```
curl -X POST \
  http://localhost:3000/api/token \
  -H 'Content-Type: application/json' \
  -d '{
	"userName":"test",
	"password":"test123"
}'
```

## response token
```
{
    "responseCode": "000",
    "responseDesc": "Success",
    "payload": [
        {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU2MzM2OTF9.KicGueJacOaQ0KjPqCeFPNO8ZmdjRHHhTgLPpyQtu8o"
        }
    ]
}
```

# Create Product
## Request 
```
curl -X POST \
  http://localhost:3000/api/create \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU2MzM2OTF9.KicGueJacOaQ0KjPqCeFPNO8ZmdjRHHhTgLPpyQtu8o' \
  -H 'Content-Type: application/json' \
  -d '{
	
	"sku":"Prd01",
	"productName":"Tas",
	"productDesc":"Tas desc",
	"quantity":10,
	"price":1000,
	"userId":"1"
}'
```

## Response 
```
{
    "responseCode": "000",
    "responseDesc": "Success"
}
```

# Update Product
## Request
```
curl -X POST \
  http://localhost:3000/api/update \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU2MzI5OTh9.cFr4QPIVayKYShW0J4m-h6CjKyXiTOjB3Hd54NcQzZw' \
  -H 'Content-Type: application/json' \
  -d '{
	
	"sku":"Prd01",
	"productName":"Tas a",
	"productDesc":"Tas a desc",
	"quantity":10,
	"price":1000,
	"userId":"1"
}'
```
## Response
```
{
    "responseCode": "000",
    "responseDesc": "Success"
}
```
# Detail Product
## Request
```
curl -X POST \
  http://localhost:3000/api/detail \
  -H 'Accept: */*' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU2MzM2OTF9.KicGueJacOaQ0KjPqCeFPNO8ZmdjRHHhTgLPpyQtu8o' \
  -H 'Content-Type: application/json' \
  -d '{
	
	"sku":"Pr01"
}'
```

## Response 
```
{
    "responseCode": "000",
    "responseDesc": "Success",
    "payload": [
        {
            "ProductID": 1,
            "SKU": "Pr01",
            "ProductName": "baju",
            "ProductDesc": "baju desc",
            "Price": 1,
            "Quantity": 10,
            "CreatedDate": "2019-08-13 00:25:58",
            "CreatedBy": "23324",
            "UpdatedDate": "2019-08-13 00:00:00",
            "UpdatedBy": "23324"
        }
    ]
}
```
# List Product
## Request
```
curl -X POST \
  http://localhost:3000/api/list \
  -H 'Accept: */*' \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU2MzM2OTF9.KicGueJacOaQ0KjPqCeFPNO8ZmdjRHHhTgLPpyQtu8o' \
  -H 'Cache-Control: no-cache' \
  -H 'Connection: keep-alive' \
  -H 'Content-Type: application/json' 
  ```
 ## Response
 ```
  {
    "responseCode": "000",
    "responseDesc": "Success",
    "payload": [
        {
            "ProductID": 1,
            "SKU": "Pr01",
            "ProductName": "baju",
            "ProductDesc": "baju desc",
            "Price": 1,
            "Quantity": 10,
            "CreatedDate": "2019-08-13 00:25:58",
            "CreatedBy": "23324",
            "UpdatedDate": "2019-08-13 00:00:00",
            "UpdatedBy": "23324"
        },
        {
            "ProductID": 3,
            "SKU": "Pr02",
            "ProductName": "baju",
            "ProductDesc": "baju desc",
            "Price": 1,
            "Quantity": 10,
            "CreatedDate": "2019-08-13 00:27:13",
            "CreatedBy": "23324",
            "UpdatedDate": "2019-08-13 00:00:00",
            "UpdatedBy": "23324"
        }
    ]
}
```
# Delete Product
## Request
```
curl -X POST \
  http://localhost:3000/api/delete \
  -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU2MzM2OTF9.KicGueJacOaQ0KjPqCeFPNO8ZmdjRHHhTgLPpyQtu8o' \
  -H 'Content-Type: application/json' \
  -d '{
	
	"sku":"Prd01"
}'
```

## Response
```
{
    "responseCode": "000",
    "responseDesc": "Success"
}
```
