package sql

import(
	"context"
	"database/sql"
	"fmt"
	"goWebApi/models"
	_ "github.com/denisenkom/go-mssqldb"
	"log"
	"time"
)

var dbContext = context.Background()

func OpenConnection() *sql.DB {
	db, err := sql.Open("sqlserver", "sqlserver://sa:12QWaszx!!@localhost?database=ManagementDb&connection+timeout=30")
	if err != nil{
		log.Fatal(err)
	}
	fmt.Println("testtttt")
	return db
}

func GetAllUser(db *sql.DB) (*config.UserList, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	rows, err := db.QueryContext(ctx,"SELECT Name,LastName,Email,Age,IsActive FROM [dbo].[Users];")
	if err != nil{
		log.Fatal(err)
	}

	defer rows.Close()
	userList := new(config.UserList)
	fmt.Println("Test")
	for rows.Next() {
		var _name string
		var _lastName string
		var _email string
		var _age int
		var _isActive bool
		result := new(config.User)

		err := rows.Scan(&_name, &_lastName, &_email, &_age, &_isActive)
		if err != nil{
			return nil, err
		} else{
			result.Name = _name
			result.LastName = _lastName
			result.Email = _email
			result.Age = _age
			result.IsActive = _isActive
			userList.UserList = append(userList.UserList, *result)
		}
	}
	return userList, nil
}

func AddUser(db *sql.DB, user config.User) error {

	query, err := db.Prepare("INSERT INTO Users (Name,LastName,Email,Password,Age,IsActive) VALUES(@p1,@p2,@p3,@p4,@p5,@p6)")
	if err != nil{
		log.Fatal("Add User Error: ", err)
	}

	var ctx context.Context
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	defer query.Close()
	rows := query.QueryRowContext(ctx, user.Name, user.LastName, user.Email, user.Password,user.Age, user.IsActive)
	if rows.Err() != nil {
		log.Fatal("Could not insert SqlDB2")
	}

	return nil
}