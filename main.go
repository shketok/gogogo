package main

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // here
	//"database/sql"
	"log"
)

//`db:"user_id"`
//`db:"user_nme"`
//`db:"user_email"`
//`db:"user_address_id"`

type ITransactionSamples interface {
	CreateUserTransaction()
}

type TransactionSamples struct {
	Db *sqlx.DB
}

func NewTransactionSamples(Db *sqlx.DB) ITransactionSamples {
	return &TransactionSamples{Db}
}

func (ts *TransactionSamples) CreateUserTransaction() {
	tx := ts.Db.MustBegin()
	var restaurantId int
	var menuItemId int

	err := tx.QueryRowx(`INSERT INTO restaurant (id, restaurant_name, rating) VALUES ($1, $2, $3) RETURNING id`,
		1, "The Rest", 4.5).Scan(&restaurantId)
	if err != nil {
		log.Fatal(err)
	}

	err = tx.QueryRowx(`INSERT INTO menu_item (id,
		                                       name,
		                                       description,
		                                       price,
		                                       decimalprice,
		                                       weight,
		                                       adult,
		                                       shippingtype,
		                                       available) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9) RETURNING id`,
		2,
		"Menu Item",
		"Desc",
		123,
		"123.123",
		"23",
		true,
		"sdfsdf",
		false,
	).Scan(&menuItemId)

	err = tx.QueryRowx(`INSERT INTO restaurant_menu_item (restaurant_id, menu_item_id)
                             VALUES ($1, $2)`, 1, 2).Scan()
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	authConnect := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s",
		"0.0.0.0",
		"5433",
		"mydbuser",
		"mydbname",
		"mydbpwd")
	db, err := sqlx.Connect("postgres", authConnect)
	if err != nil {
		log.Fatal(err)
	}
	ts := NewTransactionSamples(db)
	ts.CreateUserTransaction()
}
