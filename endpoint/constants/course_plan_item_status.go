package constants

type CoursePlanItemStatus string

const (
	CoursePlanItemPending  CoursePlanItemStatus = "PENDING"
	CoursePlanItemApproved CoursePlanItemStatus = "APPROVED"
	CoursePlanItemRejected CoursePlanItemStatus = "REJECTED"
)
