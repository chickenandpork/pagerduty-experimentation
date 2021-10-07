package lib

import (
	"context"
	//"fmt"
	pagerduty "github.com/PagerDuty/go-pagerduty"
)

func ListCurrentOncall(ctx context.Context, authtoken string, scheduleID string) ([]pagerduty.OnCall, error) {
	// per the openapi doc, `since` and `until` default to `now()`
	opts := pagerduty.ListOnCallOptions{EscalationPolicyIDs: []string{scheduleID}}
	client := pagerduty.NewClient(authtoken)

	ocs, err := client.ListOnCallsWithContext(ctx, opts)
	if err != nil {
		return []pagerduty.OnCall{}, err
	} //else

	return ocs.OnCalls, nil
}
