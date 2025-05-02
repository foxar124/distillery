package cosign

type Payload struct {
	Body           string `json:"body"`
	IntegratedTime int64  `json:"integratedTime"`
	LogIndex       int64  `json:"logIndex"`
	LogID          string `json:"logID"`
}

type Rekor struct {
	SignedEntryTimestamp string  `json:"SignedEntryTimestamp"`
	Payload              Payload `json:"Payload"`
}

type Bundle struct {
	Signature   string `json:"base64Signature"`
	Certificate string `json:"cert"`
	RekorBundle Rekor  `json:"rekorBundle"`
}
