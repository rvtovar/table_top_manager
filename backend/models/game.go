package models

import (
	"backend/db"
	"time"
)

type Game struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name" binding:"required"`
	Style    string    `json:"style" binding:"required"`
	Location string    `json:"location" binding:"required"`
	DateTime time.Time `json:"date_time" binding:"required"`
	UserId   int       `json:"user_id"`
}

func (g *Game) Save() error {
	// Save the game to the database
	query := `INSERT INTO games (name, style, location, dateTime, user_id) 
				VALUES (?, ?, ?, ?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(g.Name, g.Style, g.Location, g.DateTime, g.UserId)
	if err != nil {
		return err
	}
	g.ID, err = result.LastInsertId()
	return err
	//games = append(games, g)
}

func AllGames() ([]Game, error) {
	query := `SELECT * FROM games`
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	games := []Game{}
	for rows.Next() {
		var g Game
		var dateTimeStr string
		err := rows.Scan(&g.ID, &g.Name, &g.Style, &g.Location, &dateTimeStr, &g.UserId)
		if err != nil {
			return nil, err
		}
		g.DateTime, err = time.Parse("2006-01-02 15:04:05", dateTimeStr)
		if err != nil {
			return nil, err
		}
		games = append(games, g)
	}
	return games, nil
}

func GetGame(id int64) (*Game, error) {
	query := `select * from games where id = ?`
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var g Game
	var dateTimeStr string
	err = stmt.QueryRow(id).Scan(&g.ID, &g.Name, &g.Style, &g.Location, &dateTimeStr, &g.UserId)
	if err != nil {
		return nil, err
	}
	g.DateTime, err = time.Parse("2006-01-02 15:04:05", dateTimeStr)
	if err != nil {
		return nil, err
	}
	return &g, nil

}

func (g *Game) Update() error {
	query := `update games set name = ?, style = ?, location = ?, dateTime = ?, user_id = ? where id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(g.Name, g.Style, g.Location, g.DateTime, g.UserId, g.ID)
	return err
}

func DeleteGame(id int64) error {
	query := `
	delete from games
	where id = ?
`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(id)
	return err
}
