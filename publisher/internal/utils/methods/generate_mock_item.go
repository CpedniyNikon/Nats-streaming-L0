package methods

import (
	"github.com/brianvoe/gofakeit/v6"
	"subscriber/internal/models"
)

func GenerateMocKItem(order models.Order) models.Item {
	var Item models.Item
	Item.ChrtID = gofakeit.Number(0, 100000)
	Item.TrackNumber = order.TrackNumber
	Item.Price = gofakeit.Number(0, 100000)
	Item.RID = gofakeit.UUID()
	Item.Name = gofakeit.Name()
	Item.Sale = gofakeit.Number(0, 500)
	Item.Size = "0"
	Item.TotalPrice = order.Payment.GoodsTotal
	Item.NmID = gofakeit.Number(0, 100000)
	Item.Brand = "Vivienne Sabo"
	Item.Status = 202

	return Item
}
