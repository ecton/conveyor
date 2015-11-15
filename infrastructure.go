package main

import "database/sql"

// Infrastructure represents an individual data center or databaes cluster location
type Infrastructure struct {
	ID   int64
	Name string
}

// GetInfrastructure returns an Infrastructure from the specified id
func GetInfrastructure(id int64) (Infrastructure, error) {
	inf := Infrastructure{}
	db, _ := db()
	err := db.Get(&inf, "SELECT * FROM infrastructures WHERE id = $1", id)
	return inf, err
}

// Save updates or inserts the infrastructure
func (inf *Infrastructure) Save() error {
	db, _ := db()
	if inf.ID == 0 {
		err := db.QueryRow("INSERT INTO infrastructures (name) VALUES ($1) RETURNING id", inf.Name).Scan(&inf.ID)
		return err
	}
	_, err := db.Exec("UPDATE infrastructures SET name = $1 WHERE id = $2", inf.Name, inf.ID)
	return err
}

// Delete removes the row from the database, and if successful sets ID to 0
func (inf *Infrastructure) Delete() error {
	db, _ := db()
	result, err := db.Exec("DELETE FROM infrastructures WHERE id = $1", inf.ID)
	if err == nil {
		var rowsAffected int64
		rowsAffected, err = result.RowsAffected()
		if rowsAffected != 1 {
			err = sql.ErrNoRows
		}
		inf.ID = 0
	}
	return err
}
