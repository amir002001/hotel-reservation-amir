package types

type User struct {
	Id                string `bson:"_id,omitempty"      json:"id,omitempty"`
	FirstName         string `bson:"first_name"         json:"first_name"`
	LastName          string `bson:"last_name"          json:"last_name"`
	Email             string `bson:"email"              json:"email"`
	EncryptedPassword string `bson:"encrypted_password" json:"-"`
}

type CreateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type UpdateUserParams struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
}
