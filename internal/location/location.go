package location

import (
	"github.com/common-go/mongo"
	"time"
)

type Location struct {
	LocationId   string               `json:"locationId,omitempty" bson:"_id,omitempty" gorm:"column:locationId;primary_key"`
	Longitude    *float64             `json:"longitude,omitempty" bson:"-" gorm:"column:longitude"`
	Location     *mongo.MongoLocation `json:"-" bson:"location,omitempty" gorm:"column:location"`
	LocationName string               `json:"locationName,omitempty" bson:"locationName,omitempty"gorm:"column:locationName"`
	Info         *LocationInfo        `json:"info,omitempty" bson:"-" gorm:"column:info"`
	Description  string               `json:"description,omitempty" bson:"description,omitempty" gorm:"column:description"`
	Type         string               `json:"type,omitempty" bson:"type,omitempty" gorm:"column:type"`
	UrlId        string               `json:"urlId,omitempty" bson:"urlId,omitempty" gorm:"column:urlId"`

	Latitude     *float64             `json:"latitude,omitempty" bson:"-" gorm:"column:latitude"`

	CreatedBy    string               `json:"createdBy,omitempty" bson:"createdBy,omitempty" gorm:"column:createdBy"`
	CreatedAt    *time.Time            `json:"createdAt,omitempty" bson:"createdAt,omitempty" gorm:"column:createdAt"`
	UpdatedBy    string               `json:"updatedBy,omitempty" bson:"updatedBy,omitempty" gorm:"column:updatedBy"`
	UpdatedAt    *time.Time            `json:"updatedAt,omitempty" bson:"updatedAt,omitempty" gorm:"column:updatedAt"`
	Version      int                  `json:"version,omitempty" bson:"version,omitempty" gorm:"column:version"`
}

func (Location) CollectionName() string {
	return "location"
}
