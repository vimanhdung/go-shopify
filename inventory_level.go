package goshopify

import (
	"fmt"
	"time"
)

const inventoryLevelsBasePath = "inventory_levels"

// InventoryLevelService is an interface for interacting with the
// inventory levels endpoints of the Shopify API
// See https://help.shopify.com/en/api/reference/inventory/inventorylevel
type InventoryLevelService interface {
	List(interface{}) ([]InventoryLevel, error)
	PostAdjusts(AdjustsInventoryLevel) (*InventoryLevel, error)
	Delete(int64, int64) error
	Post(ConnectInventoryLevel) (*InventoryLevel, error)
	PostSet(SetInventoryLevel) (*InventoryLevel, error)
}

// InventoryLevelServiceOp is the default implementation of the InventoryLevelService interface
type InventoryLevelServiceOp struct {
	client *Client
}

// InventoryLevel represents a Shopify inventory item
type InventoryLevel struct {
	InventoryItemID   int64        `json:"inventory_item_id"`
	LocationID        int64        `json:"location_id"`
	Available         int32        `json:"available"`
	UpdatedAt         *time.Time `json:"updated_at"`
	AdminGraphqlAPIID string     `json:"admin_graphql_api_id"`
}

// InventoryLevelResource is used for handling single item requests and responses
type InventoryLevelResource struct {
	InventoryLevel *InventoryLevel `json:"inventory_item"`
}

// InventoryLevelsResource is used for handling multiple item responses
type InventoryLevelsResource struct {
	InventoryLevels []InventoryLevel `json:"inventory_levels"`
}

type AdjustsInventoryLevel struct {
	LocationID          int64 `json:"location_id"`
	InventoryItemID     int64 `json:"inventory_item_id"`
	AvailableAdjustment int64 `json:"available_adjustment"`
}

type ConnectInventoryLevel struct {
	LocationID      int64 `json:"location_id"`
	InventoryItemID int64 `json:"inventory_item_id"`
}

type SetInventoryLevel struct {
	LocationID      int64 `json:"location_id"`
	InventoryItemID int64 `json:"inventory_item_id"`
	Available       int64 `json:"available"`
}

// List inventory items
func (s *InventoryLevelServiceOp) List(options interface{}) ([]InventoryLevel, error) {
	path := fmt.Sprintf("%s.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelsResource)
	err := s.client.Get(path, resource, options)
	return resource.InventoryLevels, err
}

// PostAdjusts Create a inventory level
func (s *InventoryLevelServiceOp) PostAdjusts(adjustsInventoryLevel AdjustsInventoryLevel) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/adjust.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, adjustsInventoryLevel, resource)
	return resource.InventoryLevel, err
}

// Delete a inventory level
func (s *InventoryLevelServiceOp) Delete(inventoryItemID, locationID int64) error {
	return s.client.Delete(fmt.Sprintf("%s.json?inventory_item_id=%d&location_id=%d", inventoryLevelsBasePath, inventoryItemID, locationID))
}

// Post connect inventory level
func (s *InventoryLevelServiceOp) Post(connectInventoryLevel ConnectInventoryLevel) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/connect.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, connectInventoryLevel, resource)
	return resource.InventoryLevel, err
}

// PostSet set inventory level
func (s *InventoryLevelServiceOp) PostSet(setInventoryLevel SetInventoryLevel) (*InventoryLevel, error) {
	path := fmt.Sprintf("%s/set.json", inventoryLevelsBasePath)
	resource := new(InventoryLevelResource)
	err := s.client.Post(path, setInventoryLevel, resource)
	return resource.InventoryLevel, err
}
