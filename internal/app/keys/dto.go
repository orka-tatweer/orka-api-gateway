package keys

type GenerateApiKeyDTO struct {
	UserID     uint   `json:"user_id"`
	WebHookUrl string `json:"webhook_url"`
	Key        string `json:"key,omitempty"` // Optional, can be generated if not provided
}
