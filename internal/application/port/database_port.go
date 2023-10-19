package port

import db "github.com/fbriansyah/bank-ina-test/internal/adapter/database"

type DatabasePort interface {
	db.DatabaseAdapter
}
