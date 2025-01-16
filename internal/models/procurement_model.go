package models

type Procurement struct {
	ID                    int    `db:"id"`
	Date                  string `db:"date"`
	Initiator             string `db:"initiator"`
	Object                string `db:"object"`
	Status                string `db:"status"`
	Budget                int    `db:"budget"`
	RequiredAgreementDate string `db:"required_agreement_date"`
	RequiredDeliveryDate  string `db:"required_delivery_date"`
	EconomyOrOverrun      int    `db:"economy_or_overrun"`
	DaysInWork            int    `db:"days_in_work"`
	PurchaseCount         int    `db:"purchase_count"`
}
