package models

import (
	"context"
	"errors"
	"time"

	"github.com/uptrace/bun"

	"nicked.io/db"
)

type DataPoint struct {
	bun.BaseModel `bun:"table:analytics"`
	Id            int64     `bun:"id,pk,autoincrement"`
	Event         string    `bun:",notnull"`
	Location      string    `bun:",notnull"`
	Page          string    `bun:",notnull"`
	Details       string    `bun:",notnull"`
	Data1         string    `bun:",notnull"`
	Data2         string    `bun:",notnull"`
	CreatedAt     time.Time `bun:",nullzero,notnull,default:current_timestamp"`
}

func CreateDataPoint(datapoint *DataPoint, ctx context.Context) error {
	if datapoint == nil {
		return errors.New("missing or invalid data point")
	}

	_, err := db.Client.NewInsert().
		Model(datapoint).
		Exec(ctx)
	if err != nil {
		return err
	}

	return nil
}
