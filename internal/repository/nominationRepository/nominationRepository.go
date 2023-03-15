package nominationRepository

import (
	"database/sql"
	"fmt"
	"frame_up_bot/config"
	"frame_up_bot/internal/entity"

	_ "github.com/go-sql-driver/mysql"
)

func ListAllActive() (*[]entity.Nomination, error) {

	database, err := sql.Open("mysql", config.DbConnection)
	if err != nil {
		panic("Error database connection")
	}
	defer database.Close()

	rows, err := database.Query("SELECT XXXXXXXXXXXXXXX FROM XXXXXXXXXXXXXXX WHERE status = '1' ")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		} else {
			return nil, err
		}
	}
	defer rows.Close()

	var nominations []entity.Nomination

	for rows.Next() {
		n := entity.Nomination{}
		err := rows.Scan(
			&n.Id,
			&n.Name,
			&n.Status,
		)
		if err != nil {
			fmt.Println(err)
		}
		nominations = append(nominations, n)
	}

	return &nominations, nil
}
