# mth-api using redis and mysql

## create database mysql
create database mth

## create table
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


## Database configuration, redis, port, etc in the config.yml file

# getToken (token is used as authentication for api requests)
## request token 
curl -X POST \
  http://localhost:3000/api/token \
  -H 'Content-Type: application/json' \
  -d '{
	"userName":"test",
	"password":"test123"
}'

## response token
{
    "responseCode": "000",
    "responseDesc": "Success",
    "payload": [
        {
            "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1NjU2MzM2OTF9.KicGueJacOaQ0KjPqCeFPNO8ZmdjRHHhTgLPpyQtu8o"
        }
    ]
}

# create Product
## request 
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

## response 
{
    "responseCode": "000",
    "responseDesc": "Success"
}

# update product
## request
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

## response
{
    "responseCode": "000",
    "responseDesc": "Success"
}




