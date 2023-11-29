package methods

import (
	"github.com/brianvoe/gofakeit/v6"
	"strconv"
	"subscriber/internal/models"
)

func GenerateMockOrder() models.Order {

	var order models.Order

	order.OrderUid = gofakeit.UUID()
	order.TrackNumber = gofakeit.Word() + "TRACK"
	order.Entry = "WBIL"

	order.Delivery.Name = gofakeit.Name()
	order.Delivery.Phone = gofakeit.Phone()
	order.Delivery.Zip = gofakeit.Zip()
	order.Delivery.City = gofakeit.City()
	order.Delivery.Address = gofakeit.Address().Street
	order.Delivery.Region = gofakeit.Address().State
	order.Delivery.Email = gofakeit.Email()
	order.Payment.Transaction = order.OrderUid
	order.Payment.RequestID = ""
	order.Payment.Currency = gofakeit.CurrencyShort()
	order.Payment.Provider = "wbpay"
	order.Payment.Amount = gofakeit.Number(0, 1000)
	order.Payment.PaymentDT = int64(gofakeit.Number(0, 10000000))
	order.Payment.Bank = "alpha"
	order.Payment.DeliveryCost = gofakeit.Number(0, 10000)
	order.Payment.GoodsTotal = gofakeit.Number(0, 500)
	order.Payment.CustomFee = 0

	itemCount := gofakeit.Number(1, 5)

	order.Items = make([]models.Item, itemCount)
	for i := 0; i < itemCount; i++ {
		order.Items[i] = GenerateMocKItem(order)
	}

	order.Locale = "en"
	order.InternalSignature = ""
	order.CustomerID = "test"
	order.DeliveryService = "meest"
	order.Shardkey = strconv.Itoa(gofakeit.Number(0, 100000))
	order.SmID = gofakeit.Number(0, 100000)
	order.DateCreated = gofakeit.Date()
	order.OofShard = "1"

	return order
}
