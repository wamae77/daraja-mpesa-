package mpesa

const (
	APP_KEY    = "Qw8bXC5mGb5AojoSJPwttePaAgrtWGIg"
	APP_SECRET = "LUC6Du3AoZOlIdF1"
)

const endpoint = "https://sandbox.safaricom.co.ke"

type (
	MpesaSTKPushBones struct {
		BusinessShortCode,
		Password,
		Timestamp,
		TransactionType,
		Amount,
		PhoneNumber,
		PartyA,
		PartyB,
		CallBackURL,
		QueueTimeOutURL,
		AccountReference,
		TransactionDesc string
	}

	AuthResponse struct {
		Access_token string
		Expires_in   string
	}

	Credetials struct {
		APP_KEY,
		APP_SECRET string
	}
)
