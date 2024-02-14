package database

import (
	"database/sql"

	"github.com/danielzinhors/clean-go-expert/internal/entity"
)

type OrderRepository struct {
	Db *sql.DB
}

func NewOrderRepository(db *sql.DB) *OrderRepository {
	return &OrderRepository{Db: db}
}

func (r *OrderRepository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}

func (r *OrderRepository) GetTotal() (int, error) {
	var total int
	err := r.Db.QueryRow("Select count(*) from orders").Scan(&total)
	if err != nil {
		return 0, err
	}
	return total, nil
}

func (r *OrderRepository) FindAll(page, limit int, sort string) ([]entity.Order, error) {
	var orders []entity.Order
	var err error
	if sort != "" && sort != "asc" && sort != "desc" {
		sort = "asc"
	}
	if page != 0 && limit != 0 {
		err = r.Db.QueryRow("Select * from orders order by id ? limit ? offset ? ", sort, limit, page).Scan(&orders)
	} else {
		err = r.Db.QueryRow("Select * from orders order by id ? ", sort).Scan(&orders)
	}
	return orders, err
}
