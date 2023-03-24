package mpesa

const endpoint = "https://sandbox.safaricom.co.ke"

type (
	StkPushRequest struct {
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

	StkPushStatusRequest struct {
		Password          string
		Timestamp         string
		CheckoutRequestID string
		PhoneNumber       string
		BusinessShortCode string
	}

	CustomerToBusinessRequest struct {
		ShortCode     string
		CommandID     string
		Amount        string
		Msisdn        string
		BillRefNumber string
	}

	CustomerToBusinessResponse struct {
		ConversationID          string
		OriginatorCoversationID string
		ResponseDescription     string
	}

	LipaNaMpesaOnlineApiResponse struct {
		MerchantRequestID   string
		CheckoutRequestID   string
		ResponseCode        string
		ResponseDescription string
		CustomerMessage     string
	}

	StkPushTransactionQueryResponse struct {
		ResponseCode        string
		ResponseDescription string
		MerchantRequestID   string
		CheckoutRequestID   string
		ResultCode          string
		ResultDesc          string
	}

	BusinessToBusinessRequest struct {
		Initiator              string
		SecurityCredential     string
		CommandID              string
		SenderIdentifierType   string
		RecieverIdentifierType string
		Amount                 string
		PartyA                 string
		PartyB                 string
		AccountReference       string
		Remarks                string
		QueueTimeOutURL        string
		ResultURL              string
	}
)
