package location

type LocationInfo struct {
	LocationInfoID string `json:"-" bson:"_id" gorm:"column:id;primary_key"`
	Rate           int32  `json:"rate" bson:"rate" gorm:"column:rate"`
	Rate1          int32  `json:"rate1" bson:"rate1" gorm:"column:rate1"`
	Rate2          int32  `json:"rate2" bson:"rate2" gorm:"column:rate2"`
	Rate3          int32  `json:"rate3" bson:"rate3" gorm:"column:rate3"`
	Rate4          int32  `json:"rate4" bson:"rate4" gorm:"column:rate4"`
	Rate5          int32  `json:"rate5" bson:"rate5" gorm:"column:rate5"`
	RateLocation   int32  `json:"rateLocation" bson:"rateLocation" gorm:"column:rateLocation"`
}

func (LocationInfo) CollectionName() string {
	return "locationInfo"
}
