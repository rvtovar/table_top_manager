package db

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	DB, err = sql.Open("mysql", "vanica:admin@tcp(127.0.0.1:3306)/games")

	if err != nil {
		panic("Could not connect to database.")
	}

	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)

	createTables()
}

func createTables() {
	createUsers := `
	create table if not exists users(
		id int auto_increment primary key,
        username varchar(255) not null unique,
        email varchar(255) not null unique,
        password varchar(255) not null,
        created_at datetime not null
	);
`
	_, err := DB.Exec(createUsers)
	if err != nil {
		panic("Could not create users table")
	}

	createGames := `
	CREATE TABLE IF NOT EXISTS games (
		id INT AUTO_INCREMENT PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		style VARCHAR(255) NOT NULL,
		location VARCHAR(255) NOT NULL,
		dateTime DATETIME NOT NULL,
		user_id INT,
		foreign key (user_id) references users(id)
	);`
	_, err = DB.Exec(createGames)
	if err != nil {
		panic(err)
	}

	createGameRegistrations := `
		create table if not exists registrations(
			id int auto_increment primary key,
            game_id int not null,
            user_id int not null,
			foreign key (game_id) references games(id),
            foreign key (user_id) references users(id)
		)`

	_, err = DB.Exec(createGameRegistrations)
	if err != nil {
		panic("Could not create registrations table")
	}
}
