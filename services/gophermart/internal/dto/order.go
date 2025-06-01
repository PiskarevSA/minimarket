package dto

import (
	"github.com/github.com/PiskarevSA/minimarket/pkg/convtype"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/gen/sqlc/postgresql"
)

func GetOrdersByUserIDToOrders(
	userID objects.UserID,
	rows []postgresql.GetOrdersByUserIDRow,
) []entities.Order {
	orders := make([]entities.Order, len(rows))

	for i, row := range rows {
		number := objects.OrderNumber(row.Number)
		orders[i].SetNumber(number)

		orders[i].SetUserID(userID)

		status := objects.OrderStatus(row.Status)
		orders[i].SetStatus(status)

		dec, _ := convtype.NumericToDecimal(row.Accrual)
		accrual := objects.Amount(dec)

		orders[i].SetAccrual(accrual)

		orders[i].SetUploadedAt(row.UploadedAt)
	}

	return orders
}
