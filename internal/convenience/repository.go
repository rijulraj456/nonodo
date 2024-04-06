package convenience

import (
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type ConvenienceRepositoryImpl struct {
	db sqlx.DB
}

func (c *ConvenienceRepositoryImpl) CreateTables() error {
	schema := `CREATE TABLE vouchers (
		Destination text,
		Payload 	text,
		Executed	BOOLEAN,
		InputIndex 	integer,
		OutputIndex integer);`

	// execute a query on the server
	_, err := c.db.Exec(schema)
	return err
}

func (c *ConvenienceRepositoryImpl) CreateVoucher(
	ctx context.Context, voucher *ConvenienceVoucher,
) (*ConvenienceVoucher, error) {
	insertVoucher := `INSERT INTO vouchers (
		Destination,
		Payload,
		Executed,
		InputIndex,
		OutputIndex) VALUES (?, ?, ?, ?, ?)`
	c.db.MustExec(
		insertVoucher,
		voucher.Destination,
		voucher.Payload,
		voucher.Executed,
		voucher.InputIndex,
		voucher.OutputIndex,
	)
	return voucher, nil
}

func (c *ConvenienceRepositoryImpl) VoucherCount(
	ctx context.Context,
) (uint64, error) {
	var id int
	err := c.db.Get(&id, "SELECT count(*) FROM vouchers")
	if err != nil {
		return 0, nil
	}
	return uint64(id), nil
}
func (c *ConvenienceRepositoryImpl) FindVoucherByInputAndOutputIndex(
	ctx context.Context, inputIndex uint64, outputIndex uint64,
) (*ConvenienceVoucher, error) {
	query := `SELECT * FROM vouchers WHERE inputIndex = ? and outputIndex = ?`
	stmt, err := c.db.Preparex(query)
	if err != nil {
		return nil, err
	}
	var p ConvenienceVoucher
	err = stmt.Get(&p, inputIndex, outputIndex)
	if err != nil {
		return nil, err
	}
	return &p, nil
}

func (c *ConvenienceRepositoryImpl) UpdateExecuted(
	ctx context.Context, inputIndex uint64, outputIndex uint64,
	executedValue bool,
) error {
	query := `UPDATE vouchers SET Executed = ? WHERE inputIndex = ? and outputIndex = ?`
	c.db.MustExec(query, executedValue, inputIndex, outputIndex)
	return nil
}

func (c *ConvenienceRepositoryImpl) FindAllVouchers(
	ctx context.Context,
	first *int,
	last *int,
	after *string,
	before *string,
	filter []*ConvenienceFilter,
) ([]ConvenienceVoucher, error) {
	query := `SELECT * FROM vouchers `
	if len(filter) > 0 {
		query += "WHERE "
	}
	args := []interface{}{}
	for _, filter := range filter {
		if *filter.Field == "Executed" {
			query += "Executed = ? "
			if *filter.Eq == "true" {
				args = append(args, true)
			} else if *filter.Eq == "false" {
				args = append(args, false)
			} else {
				return nil, fmt.Errorf("unexpected executed value %s", *filter.Eq)
			}
		} else {
			return nil, fmt.Errorf("unexpected field %s", *filter.Field)
		}
	}
	stmt, err := c.db.Preparex(query)
	if err != nil {
		return nil, err
	}
	var vouchers []ConvenienceVoucher
	err = stmt.Select(&vouchers, args...)
	if err != nil {
		return nil, err
	}
	return vouchers, nil
}
