package dbrepo

import (
	"context"
	"time"

	"github.com/tomonar/booking/internal/models"
)

func (m *postgresDBRepo) AllUsers() bool {
	return true
}

// InsertReservation inserts a reservation into the database
func (m *postgresDBRepo) InsertReservation(res models.Reservation) error {

	// トランザクションが3秒で終わらない場合は、キャンセルする
	// トランザクションの途中でネットワーク接続が切れるなどのケースを想定して。
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	stmt := `insert into reservations (first_name, last_name, email, phone,
					start_date, end_date, room_id, created_at, updated_at)
					values ($1, $2, $3, $4, $5, $6, $7, $8, $9)`

	_, err := m.DB.ExecContext(ctx, stmt,
		res.FirstName,
		res.LastName,
		res.Email,
		res.Phone,
		res.StartDate,
		res.EndDate,
		res.RoomID,
		time.Now(),
		time.Now(),
	)

	if err != nil {
		return err
	}

	return nil
}
