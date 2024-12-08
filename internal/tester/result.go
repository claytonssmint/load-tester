package tester

import "fmt"

func (r LoadTestResults) String() string {
	report := fmt.Sprintf("Total requests: %d\n", r.TotalRequests)
	report += fmt.Sprintf("Status 200: %d\n", r.Status200)
	report += "Other Status Codes:\n"
	for status, count := range r.OtherStatus {
		report += fmt.Sprintf("%d: %d\n", status, count)
	}
	return report
}
