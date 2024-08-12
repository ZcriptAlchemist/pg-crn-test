package dao

import (
	"errors"
	"log"

	"github.com/suhailmshaik/pg-crn-test/config"
	"github.com/suhailmshaik/pg-crn-test/models"
	"gorm.io/gorm"
)

// ===================
// saves payout in DB
// ===================
func SavePayout(payout *models.Payouts) error {
	err := config.DB.Create(&payout).Error

	if err != nil {
		return errors.New("unable to save payout")
	}

	return nil
}

// ==========================
// Fetch All Payouts from DB
// ==========================
func FindPayouts() ([]models.Payouts, error) {
	var fetchedPayouts []models.Payouts

	result := config.DB.Find(&fetchedPayouts)

	if result.Error != nil {
		log.Printf("error in DAO layer %v \n", result.Error)
		return fetchedPayouts, errors.New("error fetching companies in dao layer")
	}

	return fetchedPayouts, nil
}

// ====================
// Fetch Payout by CRN
// ====================
func FindPayoutByCRN(crn int64) (models.Payouts, error) {
	var payout models.Payouts

	err := config.DB.Where("crn = ?", crn).First(&payout).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("payout with CRN %d not found \n", crn)
			return payout, errors.New("payout not found")
		}
		log.Printf("error in DAO layer %v \n", err)
		return payout, errors.New("error fetching payout from dao layer")
	}

	return payout, nil
}

// =========================
// Deletes a Payout via CRN
// =========================
func RemovePayoutByCRN(crn int64) error {
	var payout models.Payouts

	// First, find the payout by CRN
	err := config.DB.Where("crn = ?", crn).First(&payout).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Printf("payout with CRN %d not found \n", crn)
			return errors.New("payout not found")
		}
		log.Printf("error finding payout with CRN %d: %v \n", crn, err)
		return errors.New("error finding payout")
	}

	// If the record is found, delete it
	err = config.DB.Delete(&payout).Error
	if err != nil {
		log.Printf("error deleting payout with CRN %d: %v \n", crn, err)
		return errors.New("error deleting payout")
	}

	log.Printf("payout with CRN %d successfully deleted \n", crn)
	return nil
}
