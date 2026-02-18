package driver

import (
	"fmt"

	"github.com/infobloxopen/architecture-workshops2/pkg/report"
)

// Score computes a 0–100 score for a run based on error rate and latency.
func Score(data *report.RunData, s *Scenario) (int, string) {
	score := 100

	// Error rate penalty (up to -40)
	errRate := 0.0
	if data.Requests > 0 {
		errRate = float64(data.Failures) / float64(data.Requests)
	}
	if errRate > s.MaxErrRate {
		penalty := int(40 * errRate)
		if penalty > 40 {
			penalty = 40
		}
		score -= penalty
	}

	// P95 latency penalty (up to -40)
	if data.Latencies.P95 > s.MaxP95Ms {
		ratio := data.Latencies.P95 / s.MaxP95Ms
		penalty := int(40 * (ratio - 1))
		if penalty > 40 {
			penalty = 40
		}
		score -= penalty
	}

	// P99 extreme penalty (-10)
	if data.Latencies.P99 > s.MaxP95Ms*2 {
		score -= 10
	}

	if score < 0 {
		score = 0
	}

	line := fmt.Sprintf("SCORE %s: %d/100 | p95=%.0fms errRate=%.1f%% reqs=%d",
		s.Name, score, data.Latencies.P95, errRate*100, data.Requests)

	return score, line
}
