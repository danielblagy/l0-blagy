package service

import (
	"encoding/json"
	"log"

	"github.com/danielblagy/l0-blagy/entity"
	"github.com/nats-io/stan.go"
	"github.com/patrickmn/go-cache"
	"gorm.io/gorm"
)

// NATS Streaming client that subscribes to a channel and stores incoming data in a database
// also manages cache storage
type StreamManager struct {
	db           *gorm.DB
	cacheStore   *cache.Cache
	ClusterID    string
	ClientID     string
	connection   stan.Conn
	subscription stan.Subscription
}

func NewStreamManager(db *gorm.DB, cacheStore *cache.Cache, clusterID, clientID string) *StreamManager {
	sm := &StreamManager{
		db:         db,
		cacheStore: cacheStore,
		ClusterID:  clusterID,
		ClientID:   clientID,
	}

	sm.initCacheFromDatabase()

	return sm
}

func (sm *StreamManager) ConnectAndSubscribe(channelName string) error {
	sc, err := stan.Connect(sm.ClusterID, sm.ClientID)
	if err != nil {
		return err
	}

	sm.connection = sc

	// Simple Async Subscriber
	// TODO make a durable subscription instead of a reguler one (to store missed messages when was offline)
	sub, err := sc.Subscribe(channelName, sm.handleMessage)
	if err != nil {
		return err
	}

	sm.subscription = sub

	log.Print("Connected to the NATS server and subscribed to '", channelName, "' channel")

	return nil
}

func (sm *StreamManager) Close() {
	log.Print("Closing a subscriber")
	// TODO don't unsubscribe a durable subscription
	//		call this sm.subscription.Close() instead
	sm.subscription.Unsubscribe()
	sm.connection.Close()
}

func (sm *StreamManager) initCacheFromDatabase() error {
	var orders []entity.Order
	if result := sm.db.Preload("Items").Find(&orders); result.Error != nil {
		log.Print("Failed to load data from the database")
		return result.Error
	}

	for _, order := range orders {
		sm.cacheStore.Set(order.OrderUid, order, cache.NoExpiration)
	}

	return nil
}

func (sm *StreamManager) storeData(order *entity.Order) error {
	// store in db
	result := sm.db.Create(order)
	if result.Error != nil {
		return result.Error
	}

	// store in cache
	// (store order value, not the pointer)
	sm.cacheStore.Set(order.OrderUid, *order, cache.NoExpiration)

	return nil
}

func (sm *StreamManager) handleMessage(m *stan.Msg) {
	var order entity.Order

	json.Unmarshal([]byte(m.Data), &order)
	log.Println("converted incoming json to entity\n", order)

	// TODO validate incoming data

	log.Println(sm.cacheStore.Get("b563feb7b2b84b6test"))

	if err := sm.storeData(&order); err != nil {
		log.Print("Failed to store data", err)
	}

	// was used for testing
	/*orderJson, err := json.MarshalIndent(order, "", "\t")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("converted entity to json (just for testing)\n", string(orderJson))*/

	log.Println(sm.cacheStore.Get("b563feb7b2b84b6test"))
}
