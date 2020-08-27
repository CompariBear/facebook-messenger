package messenger

import "fmt"

// FacebookRequest received from Facebook server on webhook, contains messages, delivery reports and/or postbacks
type FacebookRequest struct {
	Entry []struct {
		ID        string `json:"id"`
		Messaging []struct {
			Recipient struct {
				ID string `json:"id"`
			} `json:"recipient"`
			Sender struct {
				ID      string `json:"id"`
				UserRef string `json:"user_ref,omitempty"`
			} `json:"sender"`
			Timestamp int               `json:"timestamp"`
			Message   *FacebookMessage  `json:"message,omitempty"`
			Delivery  *FacebookDelivery `json:"delivery,omitempty"`
			Postback  *FacebookPostback `json:"postback,omitempty"`
			Referral  *FacebookReferral `json:"referral,omitempty"`
		} `json:"messaging"`
		Time int `json:"time"`
	} `json:"entry"`
	Object string `json:"object"`
}

// FacebookReferral struct for all links and add references
type FacebookReferral struct {
	Ref    string `json:"ref"`
	Source string `json:"source"`
	Type   string `json:"type"`
	AdID   string `json:"ad_id,omitempty"`
	RefURI string `json:"referer_uri,omitempty"`
}

// FacebookMessage struct for text messaged received from facebook server as part of FacebookRequest struct
type FacebookMessage struct {
	Mid      string  `json:"mid"`
	Echo     bool    `json:"is_echo,omitempty"`
	AppID    int64   `json:"app_id,omitempty"`
	MetaData *string `json:"metadata,omitempty"`
	Text     string  `json:"text,omitempty"`
}

// FacebookDelivery struct for delivery reports received from Facebook server as part of FacebookRequest struct
type FacebookDelivery struct {
	Mids      []string `json:"mids"`
	Watermark int      `json:"watermark"`
}

// FacebookPostback struct for postbacks received from Facebook server  as part of FacebookRequest struct
type FacebookPostback struct {
	Payload string `json:"payload"`
}

// rawFBResponse received from Facebook server after sending the message
// if Error is null we copy this into FacebookResponse object
type rawFBResponse struct {
	MessageID   string         `json:"message_id"`
	RecipientID string         `json:"recipient_id"`
	Error       *FacebookError `json:"error"`
}

// FacebookResponse received from Facebook server after sending the message
type FacebookResponse struct {
	MessageID   string `json:"message_id"`
	RecipientID string `json:"recipient_id"`
}

// FacebookError received form Facebook server if sending messages failed
type FacebookError struct {
	Code      int    `json:"code"`
	FbtraceID string `json:"fbtrace_id"`
	Message   string `json:"message"`
	Type      string `json:"type"`
}

// Error returns Go error object constructed from FacebookError data
func (err *FacebookError) Error() error {
	return fmt.Errorf("FB Error: Type %s: %s; FB trace ID: %s", err.Type, err.Message, err.FbtraceID)
}
