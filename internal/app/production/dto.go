package production

type ProductionSchedule struct {
	Schedule    map[string][]ProductionItem `json:"schedule"`
	Explanation string                      `json:"explanation"`
}

type ProductionItem struct {
	Product  string  `json:"product"`
	Quantity float64 `json:"quantity"`
}
