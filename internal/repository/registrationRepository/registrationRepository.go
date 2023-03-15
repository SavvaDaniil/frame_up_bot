package registrationRepository

import (
	"database/sql"
	"fmt"
	"frame_up_bot/config"
	"frame_up_bot/internal/entity"

	_ "github.com/go-sql-driver/mysql"
)

func ListAllAfter2023() (*[]entity.Registration, error) {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		panic("Error database connection")
	}
	defer database.Close()

	rows, err := database.Query("SELECT ... FROM XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX WHERE date_of_add > 'XXXXXXXXXXXXXXXXXXXXXXXX' ORDER BY id")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	var registrations []entity.Registration

	for rows.Next() {
		r := entity.Registration{}
		err := rows.Scan(
			&r.Id,
			&r.NominationId,
			&r.Name,
			&r.TeamName,
			&r.HowCal,
			&r.City,
			&r.Phone,
			&r.Email,
			&r.Links,
			&r.Dancers,
			&r.DateOfAdd,
		)
		if err != nil {
			fmt.Println(err)
		}
		registrations = append(registrations, r)
	}

	return &registrations, nil
}
