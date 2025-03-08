package auth

// Define the struct for your claims
type UserClaims struct {
	ID        int
	Email     string
	FirstName string
	LastName  string
	Issuer    string
	Exp       int64
	Iat       int64
}
