package client

// duplicate of scheduler/dto TaskDTO to avoid cercular import errors
type TaskDTO struct {
	ID                int            `json:"id"`
	Name              string         `json:"name"`
	Duration          int            `json:"duration"`
	Dependencies      []uint         `json:"dependencies,omitempty"`
	ResourcesRequired map[string]int `json:"resources_required,omitempty"`
	Location          string         `json:"location"`
	Priority          int            `json:"priority"`
	EarliestStart     int            `json:"earliest_start,omitempty"`
	LatestEnd         int            `json:"latest_end,omitempty"`
	CostPerHour       float64        `json:"cost_per_hour,omitempty"`
}

type LogisticsRequestDTO struct {
	Tasks         []TaskDTO                 `json:"tasks"`
	ResourcePool  map[string]int32          `json:"resource_pool"`
	TransitMatrix map[string]map[string]int `json:"transit_matrix"`
	Objective     string                    `json:"objective"`
	Vehicles      []string                  `json:"vehicles"`
}

type ScheduleTask struct {
	Name      string         `json:"name"`
	Start     int            `json:"start"`
	End       int            `json:"end"`
	Resources map[string]int `json:"resources"`
	Location  string         `json:"location"`
	Vehicle   string         `json:"vehicle"`
}

type ScheduleResponse struct {
	Schedule  map[string]ScheduleTask `json:"schedule"`
	Makespan  *int                    `json:"makespan,omitempty"`
	TotalCost *float64                `json:"total_cost,omitempty"`
}

type OptimizeScheduleResponse struct {
	Result ScheduleResponse `json:"result"`
}

type ProductionPlanRequest struct {
	Demand    []DemandForecast `json:"demand"`
	Resources []Resource       `json:"resources"`
	Capacity  ProductCapacity  `json:"capacity"`
}

type DemandForecast struct {
	ProductID string `json:"product_id"`
	Quantity  int32  `json:"quantity"`
}

type Resource struct {
	MachineID string `json:"machine_id"`
	Type      string `json:"type"`
	Capacity  int32  `json:"capacity"`
}

type ProductCapacity struct {
	MaxShifts      int     `json:"max_shifts"`
	HoursPerShift  int     `json:"hours_per_shift"`
	DowntimeFactor float64 `json:"downtime_factor"`
}

type ProductSchedule struct {
	Product  string `json:"product"`
	Quantity int32  `json:"quantity"`
}

type ScheduleEntry struct {
	MachineID string            `json:"machine_id"`
	Products  []ProductSchedule `json:"products"`
}

type MachineSchedule struct {
	Schedules []ProductSchedule `json:"schedules"`
}

type ProductionResponse struct {
	Schedule    map[string]MachineSchedule `json:"schedule"`
	Explanation string                     `json:"explanation"`
}
