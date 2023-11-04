package models

type Transaction struct {
	CreatedAt          string                 `json:"createdAt"`
	ProcessName        string                 `json:"processName"`
	ProcessVersion     int                    `json:"processVersion"`
	LastTransition     string                 `json:"lastTransition"`
	LastTransitionedAt string                 `json:"lastTransitionedAt"`
	PayinTotal         Price                  `json:"payinTotal"`
	PayoutTotal        Price                  `json:"payoutTotal"`
	LineItems          []LineItem             `json:"lineItems"`
	ProtectedData      map[string]interface{} `json:"protectedData"`
	Metadata           map[string]interface{} `json:"metadata"`
	Transitions        []Transition           `json:"transitions"`
}

type Transition struct {
	Transition string `json:"transition"`
	CreatedAt  string `json:"createdAt"`
	By         string `json:"by"`
}

type LineItem struct {
	Code       string   `json:"code"`
	Quantity   int      `json:"quantity"`
	Units      int      `json:"units"`
	Seats      int      `json:"seats"`
	Reversal   bool     `json:"reversal"`
	UnitPrice  Price    `json:"unitPrice"`
	LineTotal  Price    `json:"lineTotal"`
	IncludeFor []string `json:"includeFor"`
	Percentage float64  `json:"percentage"`
}
