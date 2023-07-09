package main

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
	GetAccountByID(int) (*Account, error)
	GetAccounts() ([]*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=123456 sslmode=disable"
	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.CreateAccountTable()
}

func (s *PostgresStore) CreateAccountTable() error {

	query := `CREATE TABLE IF NOT EXISTS ACCOUNT(
		ID         INT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
		FIRST_NAME  VARCHAR(30),
		LAST_NAME   VARCHAR(30),
		NUMBER     INT,
		BALANCE    INT,
		C TIMESTAMP, 
		INSERT_AT TIMESTAMP DEFAULT NOW()
	);`

	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) CreateAccount(a *Account) error {
	query := fmt.Sprintf(
		`INSERT INTO ACCOUNT(
		FIRST_NAME,
		LAST_NAME,
		NUMBER,
		BALANCE,
		CREATED_AT
	) VALUES (
		'%v',
		'%v', 
		%v, 
		%v,
		'%v'
	);`, a.FirstName, a.LastName, a.Number, a.Balance, a.CreatedAt.Format(time.RFC3339))

	fmt.Println(query)
	_, err := s.db.Exec(query)
	return err
}

func (s *PostgresStore) DeleteAccount(ID int) error {
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) GetAccountByID(ID int) (*Account, error) {
	query := fmt.Sprintf("SELECT * FROM ACCOUNT WHERE ID = '%v'", ID)
	row, err := s.db.Query(query)

	if err != nil {
		return nil, err
	}
	acc := new(Account)
	for row.Next() {
		err := row.Scan(
			&acc.ID,
			&acc.FirstName,
			&acc.LastName,
			&acc.Number,
			&acc.Balance,
			&acc.CreatedAt,
			&acc.InsertedAt,
		)

		if err != nil {
			return nil, err
		}
	}
	return acc, nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {

	query := "SELECT * FROM ACCOUNT"
	rows, err := s.db.Query(query)
	if err != nil {
		return nil, err
	}
	accounts := []*Account{}

	for rows.Next() {
		acc := new(Account)
		err := rows.Scan(
			&acc.ID,
			&acc.FirstName,
			&acc.LastName,
			&acc.Number,
			&acc.Balance,
			&acc.CreatedAt,
			&acc.InsertedAt,
		)

		if err != nil {
			return nil, err
		}

		accounts = append(accounts, acc)

	}

	return accounts, nil
}
