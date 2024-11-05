package auth

// Define the struct for your claims
type UserClaims struct {
	ID        int    `mapstructure:"id"`
	Email     string `mapstructure:"email"`
	FirstName string `mapstructure:"fname"`
	LastName  string `mapstructure:"lname"`
	Issuer    string `mapstructure:"iss"`
	Exp       int64  `mapstructure:"exp"`
	Iat       int64  `mapstructure:"iat"`
}
