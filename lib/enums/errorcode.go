package enums

const (
	// general
	UnknownError int = 1000
	// validation errors
	RequiredParamsError int = 2000
	InvalidAddressError int = 2001
	ValidationError     int = 2002
	// integration errors
	PostHogError int = 4000
	// pairing
	PairingNotFoundError int = 5000
)
