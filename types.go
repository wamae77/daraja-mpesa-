package mpesa

const (
	APP_KEY    = "GvzjNnYgNJtwgwfLBkZh65VPwfuKvs0V"
	APP_SECRET = "oOpJICRVlyrGSAkM"
)

type (
	STKPushSimulationD struct {
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
)
