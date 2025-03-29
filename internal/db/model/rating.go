package model

import "gorm.io/gorm"

type Rating struct {
	ID   int    `gorm:"primaryKey; autoIncrement" json:"id"`
	Link string `gorm:"type:varchar(255); not null" json:"link"`
	Rate int    `gorm:"type:integer; not null" json:"rate"`
}

func AddRating(db *gorm.DB, link string, rate int) (*Rating, error) {
	rating := Rating{
		Link: link,
		Rate: rate,
	}
	err := db.Create(&rating).Error
	if err != nil {
		return nil, err
	}
	return &rating, nil
}

// func GetRating(db *gorm.DB, link string) (*Rating, error) {
// 	var rating Rating
// 	err := db.Where("link = ?", link).First(&rating).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &rating, nil
// }

// func ChangeRating(db *gorm.DB, link string, rate int) (*Rating, error) {
// 	var rating Rating
// 	err := db.Where("link = ?", link).First(&rating).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	rating.Rate = rate
// 	err = db.Save(&rating).Error
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &rating, nil
// }
