package models

type Procurement struct {
	Date                  string
	Initiator             string
	Object                string
	Status                string
	Budget                int
	RequiredAgreementDate string
	RequiredDeliveryDate  string
	EconomyOrOverrun      int
	DaysInWork            int
	PurchaseCount         int
}
