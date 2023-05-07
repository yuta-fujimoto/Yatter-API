package object

type (
	StatusID = int64

	Status struct {
		// The internal ID of the status
		ID StatusID `json:"-"`

		// The internal ID of the account
		AccountID AccountID `json:"-" db:"account_id"`

		// Status
		Content *string `json:"status,omitempty"`

		// The time the status was created
		CreateAt DateTime `json:"create_at,omitempty" db:"create_at"`
	}
)
