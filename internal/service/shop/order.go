package serviceShop

import (
	"log"
	"time"

	"github.com/oogway93/golangArchitecture/internal/entity/products"
	"github.com/oogway93/golangArchitecture/internal/repository"
	"github.com/oogway93/golangArchitecture/internal/repository/postgres/models"
	"github.com/shopspring/decimal"
)

const (
	Shipped   = "Shipped"
	Delivered = "Delivered"
	PickedUp  = "Picked_up"
	Completed = "Completed"
)

type OrderShopService struct {
	repo  repository.OrderRepository
	cache repository.CacheRepository
}

func NewServiceShopOrder(repo repository.OrderRepository, cache repository.CacheRepository) *OrderShopService {
	return &OrderShopService{
		repo:  repo,
		cache: cache,
	}
}

func (s *OrderShopService) Create(userID string, requestData *products.Order) {
	deliveryModel := models.Delivery{
		FullName:      requestData.DeliveryRel.FullName,
		Postcode:      requestData.DeliveryRel.Postcode,
		Country:       requestData.DeliveryRel.Country,
		City:          requestData.DeliveryRel.City,
		DeliveryPrice: requestData.DeliveryRel.DeliveryPrice,
	}

	var orderItems []*models.OrderItem
	for _, productItem := range requestData.OrderItemsRel {
		resultProduct := s.repo.FetchProductID(productItem.ProductRel.ProductName)
		unitPrice, ok := resultProduct["price"].(decimal.Decimal)
		if !ok {
			log.Printf("Failed to convert price to decimal for product %s", productItem.ProductRel.ProductName)
			continue
		}
		orderItem := models.OrderItem{
			ProductID: resultProduct["id"].(uint),
			Quantity:  productItem.Quantity,
			UnitPrice: unitPrice,
		}
		orderItems = append(orderItems, &orderItem)

	}
	if len(orderItems) == 0 {
		log.Println("No order items found")
		return
	}

	s.repo.CreateDelivery(&deliveryModel)

	deliveryID, err := s.repo.LastRow()
	if err != nil {
		log.Fatalf("Error in getting last row from delivery: %v", err.Error())
	}

	order := s.repo.CreateOrderAndOrderItems(userID, deliveryID, orderItems)
	go s.autoUpdateStatus(order.ID)
}

func (s *OrderShopService) autoUpdateStatus(orderID uint) {
	statusUpdates := []struct {
		status string
		delay  time.Duration
	}{
		{Delivered, 10 * time.Minute},
		{Shipped, 30 * time.Minute},
		{PickedUp, 50 * time.Minute},
		{Completed, 50*time.Minute + 1*time.Second},
	}
	for _, update := range statusUpdates {
		time.Sleep(update.delay)
		s.repo.UpdateOrderStatus(orderID, update.status)
		time.Sleep(1 * time.Second)
		if update.status == "Completed" {
			s.repo.Delete(orderID)
		}
	}
}
func (s *OrderShopService) GetAll(userID string) []map[string]interface{} {
	result := s.repo.GetAll(userID)
	return result
}
func (s *OrderShopService) Get(orderID string) map[string]interface{}              { return nil }
func (s *OrderShopService) Update(orderID string, requestData *models.Order) error { return nil }
func (s *OrderShopService) Delete(orderID string) error                            { return nil }
