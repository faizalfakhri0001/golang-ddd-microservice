package domain

type User struct {
	ID        uint   `json:"id,omitempty"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Address   string `json:"address,omitempty"`
	City      string `json:"city,omitempty"`
}
