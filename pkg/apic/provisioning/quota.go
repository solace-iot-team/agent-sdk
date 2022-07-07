package provisioning

import "github.com/Axway/agent-sdk/pkg/apic/apiserver/models/management/v1alpha1"

// Quota - interface for accessing an access requests quota
type Quota interface {
	// GetInterval returns the quota interval from within the access request
	GetInterval() QuotaInterval
	// GetIntervalString returns the string representation of the quota interval from within the access request
	GetIntervalString() string
	// GetLimit returns the quota limit from within the access request
	GetLimit() int64
}

// QuotaInterval is the quota limit
type QuotaInterval int

const (
	// Daily -
	Daily QuotaInterval = iota + 1
	// Weekly -
	Weekly
	// Monthly -
	Monthly
	// Annually -
	Annually
)

// String returns the string value of the State
func (q QuotaInterval) String() string {
	return map[QuotaInterval]string{
		Daily:    "daily",
		Weekly:   "weekly",
		Monthly:  "monthly",
		Annually: "annually",
	}[q]
}

// quotaIntervalFromString returns the quota limit represented by the string sent in
func quotaIntervalFromString(limit string) QuotaInterval {
	if q, ok := map[string]QuotaInterval{
		"daily":    Daily,
		"weekly":   Weekly,
		"monthly":  Monthly,
		"annually": Annually,
	}[limit]; ok {
		return q
	}
	return -1
}

type quota struct {
	interval QuotaInterval
	limit    int64
}

//NewQuotaFromAccessRequest create a Quota interface from an access request or nil if no quota on access request
func NewQuotaFromAccessRequest(ar *v1alpha1.AccessRequest) Quota {
	if ar.Spec.Quota == nil {
		return nil
	}
	interval := quotaIntervalFromString(ar.Spec.Quota.Interval)
	if interval == -1 {
		return nil
	}
	return &quota{
		limit:    int64(ar.Spec.Quota.Limit),
		interval: interval,
	}
}

func (q *quota) GetInterval() QuotaInterval {
	return q.interval
}

func (q *quota) GetIntervalString() string {
	return q.interval.String()
}

func (q *quota) GetLimit() int64 {
	return q.limit
}