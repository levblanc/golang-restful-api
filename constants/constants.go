package constants

type contextKey int

const (
	// ContextUserID is the context key for user id
	ContextUserID contextKey = iota
)

const (
	// CookieName is the name of session cookie on browser
	CookieName = "mstream-session"
)
