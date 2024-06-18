package repo

import (
	"awesomeProject/internal/pkg/model"
	"context"
	"fmt"
	"github.com/jackc/pgtype/pgxtype"
)

const (
	create = "insert into cat (name, age) values ($1, $2)"
	find   = "SELECT * FROM cat WHERE id = $1"
	update = "UPDATE cat SET name = $1, ag = $2 WHERE id = $3"
	delete = "DELETE FROM cat WHERE id = $1"
)

type CatRepository struct {
	db pgxtype.Querier
}

// в zuzu возращалась ссылка *
func NewCatRepository(db pgxtype.Querier) *CatRepository {
	return &CatRepository{
		db: db,
	}
}

func (r *CatRepository) Save(ctx context.Context, c model.Cat) error {
	_, err := r.db.Exec(ctx, create, c.Name, c.Age)
	if err != nil {
		err = fmt.Errorf("error happened in db.Exec: %w", err)
		return err
	}
	return nil
}

/*func (r *CatRepository) findCat(id uuid.UUID) (model.Cat, error) {
	row := r.db.QueryRow(r.ctx, find, id)
	if err := row.Scan(id); err != nil {
		return (nil, err)
	}
	return (row.Scan(id), nil)
}
*/
