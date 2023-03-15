package entity

type Registration struct {
	Id           int
	NominationId int
	Name         *string
	TeamName     *string
	HowCal       *string
	City         *string
	Phone        *string
	Email        *string
	Links        *string
	Dancers      *string
	DateOfAdd    *string
}
