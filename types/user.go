package types

type User struct {
	ID string `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string `bson:"firstName,omitempty" json:"firstName,omitempty"`
	LastName string `bson:"lastName,omitempty" json:"lastName,omitempty"`
	// Email string `bson:"email,omitempty" json:"email,omitempty"`
	// Password string `bson:"password,omitempty" json:"password,omitempty"`


}