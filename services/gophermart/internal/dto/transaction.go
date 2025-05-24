package dto

import (
	"github.com/github.com/PiskarevSA/minimarket/pkg/convtype"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/entities"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/domain/objects"
	"github.com/github.com/PiskarevSA/minimarket/services/gophermart/internal/gen/sqlc/postgresql"
)

func GetTxsByUserIDToTxs(
	userID objects.UserID,
	rows []postgresql.GetTransactionsByUserIdRow,
) []entities.Transaction {
	txs := make([]entities.Transaction, len(rows))

	for i, row := range rows {
		txs[i].SetID(row.Id)
		txs[i].SetUserID(userID)

		orderNumber := objects.OrderNumber(row.OrderNumber)
		txs[i].SetOrderNumber(orderNumber)

		decimal, _ := convtype.NumericToDecimal(row.Sum)
		sum := objects.Amount(decimal)

		txs[i].SetSum(sum)
		txs[i].SetProcessedAt(row.ProcessedAt)
	}

	return txs
}
