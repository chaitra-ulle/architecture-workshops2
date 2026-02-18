package driver

import "time"

// Scenario defines a load test configuration for a specific lab case.
type Scenario struct {
	Name        string
	Description string
	TargetURL   string
	Method      string
	Body        string
	RPS         int
	Duration    time.Duration
	Concurrency int
	MaxP95Ms    float64
	MaxErrRate  float64
	DBStatsURL  string
	HPAStatsURL string
	BatchURL    string
}

// Registry maps scenario names to their configs.
var Registry = map[string]*Scenario{
	"timeouts": {
		Name:        "timeouts",
		Description: "Case 1: Timeout patterns — calls to slow dependency without deadline",
		TargetURL:   "http://localhost:8080/cases/timeouts",
		Method:      "GET",
		RPS:         10,
		Duration:    30 * time.Second,
		Concurrency: 20,
		MaxP95Ms:    2500,
		MaxErrRate:  0.1,
	},
	"tx": {
		Name:        "tx",
		Description: "Case 2: DB transaction scope — holding TX across network calls",
		TargetURL:   "http://localhost:8080/cases/tx",
		Method:      "GET",
		RPS:         10,
		Duration:    30 * time.Second,
		Concurrency: 20,
		MaxP95Ms:    3000,
		MaxErrRate:  0.1,
		DBStatsURL:  "http://localhost:8080/debug/dbstats",
	},
	"bulkheads": {
		Name:        "bulkheads",
		Description: "Case 3: Bulkhead pattern — shared pool starvation",
		TargetURL:   "http://localhost:8081/batches",
		Method:      "POST",
		Body:        `{"fast": 100, "slow": 20}`,
		RPS:         1,
		Duration:    10 * time.Second,
		Concurrency: 5,
		MaxP95Ms:    500,
		MaxErrRate:  0.05,
		BatchURL:    "http://localhost:8081/batches",
	},
	"autoscale": {
		Name:        "autoscale",
		Description: "Case 4: Autoscaling — CPU-bound without HPA",
		TargetURL:   "http://localhost:8080/cases/autoscale",
		Method:      "GET",
		RPS:         20,
		Duration:    60 * time.Second,
		Concurrency: 30,
		MaxP95Ms:    5000,
		MaxErrRate:  0.1,
	},
}

// ListScenarios returns all scenario names.
func ListScenarios() []string {
	names := make([]string, 0, len(Registry))
	for name := range Registry {
		names = append(names, name)
	}
	return names
}
