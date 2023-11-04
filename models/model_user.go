package models

type Agreement bool

type User struct {
	Deleted           bool              `json:"deleted"`
	Banned            bool              `json:"banned"`
	Email             string            `json:"email"`
	StripeConnected   bool              `json:"stripeConnected"`
	CreatedAt         string            `json:"createdAt"`
	IdentityProviders []interface{}     `json:"identityProviders"`
	PendingEmail      string            `json:"pendingEmail"`
	EmailVerified     bool              `json:"emailVerified"`
	Activity          map[string]string `json:"activity"`
	Profile           struct {
		DisplayName     string      `json:"displayName"`
		FirstName       string      `json:"firstName"`
		Bio             string      `json:"bio"`
		AbbreviatedName string      `json:"abbreviatedName"`
		LastName        string      `json:"lastName"`
		PrivateData     interface{} `json:"privateData"`
		ProtectedData   interface{} `json:"protectedData"`
		PublicData      interface{} `json:"publicData"`
		Metadata        interface{} `json:"metadata"`
	} `json:"profile"`
}
