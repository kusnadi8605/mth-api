package models

import (
	"encoding/json"
	"fmt"
	conf "mth-api/config"
	dts "mth-api/datastruct"
	"time"

	"github.com/gomodule/redigo/redis"
)

//Create Product
func Create(conn *conf.Connection, productReq dts.ProductReq) error {
	arrProduct := []dts.Product{}
	strProduct := dts.Product{}

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	createdDate := time.Now().Format("2006-01-02- 15:04:05")

	stmt, err := tx.Prepare(`insert into mtr_product 
						     (sku,productName,ProductDesc,price,quantity,createdBy,createdDate)
							  values (?,?,?,?,?,?,?)`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(productReq.SKU, productReq.ProductName, productReq.ProductDesc, productReq.Price, productReq.Quantity, productReq.UserID, createdDate)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	rows, err := res.LastInsertId()
	// build data to be stored in redis
	strProduct.ProductID = rows
	strProduct.SKU = productReq.SKU
	strProduct.ProductName = productReq.ProductName
	strProduct.ProductDesc = productReq.ProductDesc
	strProduct.Quantity = productReq.Quantity
	strProduct.Price = productReq.Price
	strProduct.CreatedBy = productReq.UserID
	strProduct.CreatedDate = createdDate
	strProduct.UpdatedBy = ""
	strProduct.UpdatedDate = ""
	arrProduct = append(arrProduct, strProduct)

	if err != nil {
		return err
	}

	//save data to redis
	key := conf.Param.RedisKEY + productReq.SKU
	jsonData, err := json.Marshal(arrProduct)
	err = SetTex(key, conf.Param.RedisEXP, jsonData)
	if err != nil {
		return err
	}
	return nil
}

//Update Product
func Update(conn *conf.Connection, productReq dts.ProductReq) error {
	arrProduct := []dts.Product{}
	strProduct := dts.Product{}

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	updatedDate := time.Now().Format("2006-01-02- 15:04:05")

	stmt, err := tx.Prepare(`update mtr_product set
							 productName=?,
							 ProductDesc=?,
							 price=?,
							 quantity=?,
							 updatedBy=?,
							 updatedDate=?
							 where sku=?
							 `)

	if err != nil {
		return err
	}

	defer stmt.Close()

	res, err := stmt.Exec(productReq.ProductName, productReq.ProductDesc, productReq.Price, productReq.Quantity, productReq.UserID, updatedDate, productReq.SKU)

	if err != nil {
		tx.Rollback()
		return err
	}

	tx.Commit()

	rows, err := res.LastInsertId()
	// build data to be stored in redis
	strProduct.ProductID = rows
	strProduct.SKU = productReq.SKU
	strProduct.ProductName = productReq.ProductName
	strProduct.ProductDesc = productReq.ProductDesc
	strProduct.Quantity = productReq.Quantity
	strProduct.Price = productReq.Price
	strProduct.CreatedBy = productReq.UserID
	strProduct.CreatedDate = updatedDate
	strProduct.UpdatedBy = productReq.UserID
	strProduct.UpdatedDate = updatedDate
	arrProduct = append(arrProduct, strProduct)

	if err != nil {
		return err
	}

	//save data to redis
	key := conf.Param.RedisKEY + productReq.SKU
	jsonData, err := json.Marshal(arrProduct)
	err = SetTex(key, conf.Param.RedisEXP, jsonData)
	if err != nil {
		return err
	}

	return nil
}

//Detail product
//check data in redis, if not exist query to DB
func Detail(conn *conf.Connection, sku string) ([]dts.Product, error) {
	arrProduct := []dts.Product{}
	strProduct := dts.Product{}
	// key redis
	key := conf.Param.RedisKEY + sku
	jsonString, err := Get(key)

	// data not found
	if err == redis.ErrNil {
		fmt.Println("data dari db")
		strQuery := `select productId,sku,productName,productDesc,
					 price,quantity,createdDate,createdBy,
					 updatedDate,updatedBy
					 from mtr_product 
					 where sku =?`

		rows, err := conn.Query(strQuery, sku)
		defer rows.Close()

		if err != nil {
			return nil, err
		}

		for rows.Next() {
			err := rows.Scan(&strProduct.ProductID,
				&strProduct.SKU,
				&strProduct.ProductName,
				&strProduct.ProductDesc,
				&strProduct.Price,
				&strProduct.Quantity,
				&strProduct.CreatedDate,
				&strProduct.CreatedBy,
				&strProduct.UpdatedDate,
				&strProduct.UpdatedBy,
			)

			if err = rows.Err(); err != nil {
				return nil, err
			}

			arrProduct = append(arrProduct, strProduct)
		}

		// if data not exist
		if len(arrProduct) < 1 {
			return nil, fmt.Errorf("data not found")
		}

		//save to redis and set expired time x second,
		//see config in config.yml
		jsonData, err := json.Marshal(arrProduct)
		err = SetTex(key, conf.Param.RedisEXP, jsonData)
		if err != nil {
			return nil, err
		}

	} else if err != nil {
		return nil, err
	} else {
		jsonData := []byte(jsonString)
		var err = json.Unmarshal(jsonData, &arrProduct)
		if err != nil {
			return nil, err
		}
	}

	return arrProduct, nil
}

//Delete Product
func Delete(conn *conf.Connection, sku string) error {

	tx, err := conn.Begin()
	if err != nil {
		return err
	}

	stmt, err := tx.Prepare(`delete from mtr_product where sku=?`)

	if err != nil {
		return err
	}

	defer stmt.Close()

	stmt.Exec(sku)
	tx.Commit()

	// chek key exist in redis
	// if exist then delete
	key := conf.Param.RedisKEY + sku
	keyExist, err := Exists(key)

	if err != nil {
		return err
	}

	if keyExist {
		err := Del(key)
		if err != nil {
			return err
		}
	}
	return nil
}

//List product
func List(conn *conf.Connection) ([]dts.Product, error) {
	arrProduct := []dts.Product{}
	strProduct := dts.Product{}

	fmt.Println("data dari db")
	strQuery := `select productId,sku,productName,productDesc,
					 price,quantity,createdDate,createdBy,
					 updatedDate,updatedBy
					 from mtr_product`

	rows, err := conn.Query(strQuery)
	defer rows.Close()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(&strProduct.ProductID,
			&strProduct.SKU,
			&strProduct.ProductName,
			&strProduct.ProductDesc,
			&strProduct.Price,
			&strProduct.Quantity,
			&strProduct.CreatedDate,
			&strProduct.CreatedBy,
			&strProduct.UpdatedDate,
			&strProduct.UpdatedBy,
		)

		if err = rows.Err(); err != nil {
			return nil, err
		}

		arrProduct = append(arrProduct, strProduct)
	}
	// if data not exist
	if len(arrProduct) < 1 {
		return nil, fmt.Errorf("data not found")
	}

	return arrProduct, nil
}
