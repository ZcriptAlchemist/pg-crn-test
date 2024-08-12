package service

import (
	"github.com/suhailmshaik/pg-crn-test/dao"
	"github.com/suhailmshaik/pg-crn-test/models"
)

// ===============
// Creates Payout
// ===============
func CreatePayout(payout *models.Payouts) error {
	err := dao.SavePayout(payout)

	if err != nil {
		return err
	}

	return nil
}

// ====================
// Fetches All Payouts
// ====================
func FetchAllPayouts() ([]models.Payouts, error) {
	allPayouts, err := dao.FindPayouts()

	if err != nil {
		return nil, err
	}

	return allPayouts, nil
}

// =========================
// Fetch Payout through CRN
// =========================
func FetchPayoutByCRN(crn int64) (models.Payouts, error) {
	payout, err := dao.FindPayoutByCRN(crn)

	if err != nil {
		return payout, err
	}

	return payout, err
}

// =========================
// Deletes Payout using CRN.
// =========================
func DeletePayoutByCRN(crn int64) error {
	err := dao.RemovePayoutByCRN(crn)

	if err != nil {
		return err
	}

	return nil
}