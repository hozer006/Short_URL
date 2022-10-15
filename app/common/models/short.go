package models

func GetAllShort() ([]ShortURL, error) {
	var short []ShortURL

	tx := db.Find(&short)
	if tx.Error != nil {
		return []ShortURL{}, tx.Error
	}
	return short, nil
}

func GetShort(id uint64) (ShortURL, error) {
	var short ShortURL

	tx := db.Where("id = ?", id).First(&short)

	if tx.Error != nil {
		return ShortURL{}, tx.Error
	}

	return short, nil
}

func CreateShort(short ShortURL) error {
	tx := db.Create(&short)
	return tx.Error
}

func UpdateShort(short ShortURL) error{
	tx := db.Save(&short)
	return tx.Error
}

func DeleteShort(id uint64) error{
	tx := db.Unscoped().Delete(&ShortURL{}, id)
	return tx.Error
}

func FindByShortUrl(url string) (ShortURL, error) {
	var short ShortURL
	tx := db.Where("short_URL = ?", url).First(&short)
	return short, tx.Error
}