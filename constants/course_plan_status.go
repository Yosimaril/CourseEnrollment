package constants

type CoursePlanStatus string

const (
	CoursePlanDraft             CoursePlanStatus = "DRAFT"
	CoursePlanApproved          CoursePlanStatus = "APPROVED"
	CoursePlanPartiallyApproved CoursePlanStatus = "PARTIALLY_APPROVED"
	CoursePlanRejected          CoursePlanStatus = "REJECTED"
	CoursePlanSubmitted         CoursePlanStatus = "SUBMITTED"
)
