package student

type Student struct {
	ID int64 `bun:",pk" json:"id"`

	FirstName string `bun:",notnull" json:"first_name"`
	LastName  string `bun:",notnull" json:"last_name"`

	Active bool `bun:",notnull" json:"active"`

	Role Role `bun:"type:tinyint,notnull" json:"role"`
}

type Role int

func (r Role) IsValid() bool {
	return r >= RoleStudent && r <= RoleAdmin
}

const (
	RoleStudent Role = iota
	RoleRedactor
	RoleAdmin
)
