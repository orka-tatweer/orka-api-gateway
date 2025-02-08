package scheduler

type TaskDTO struct {
	ID                int            `json:"id"`
	Name              string         `json:"name"`
	Duration          int            `json:"duration"`
	Dependencies      []uint         `json:"dependencies,omitempty"`
	ResourcesRequired map[string]int `json:"resources_required,omitempty"`
	Location          string         `json:"location"`
	Priority          int            `json:"priority"`
	EarliestStart     *int           `json:"earliest_start,omitempty"`
	LatestEnd         *int           `json:"latest_end,omitempty"`
	CostPerHour       *float64       `json:"cost_per_hour,omitempty"`
}

type LogisticsRequestDTO struct {
	Tasks         []TaskDTO                 `json:"tasks"`
	ResourcePool  map[string]int            `json:"resource_pool"`
	TransitMatrix map[string]map[string]int `json:"transit_matrix"`
	Objective     string                    `json:"objective"`
	Vehicles      []string                  `json:"vehicles"`
}
