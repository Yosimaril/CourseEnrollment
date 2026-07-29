export const CoursePlanStatusMapper = {
    DRAFT: "Draft",
    PARTIALLY_APPROVED: "Partially Approved",
    APPROVED: "Approved",
    REJECTED: "Rejected",
    SUBMITTED: "Submitted",
} as const;

import { CoursePlanStatusEnum } from "~/constants/coursePlanStatusEnum";

export const CoursePlanItemStatusMapper = {
    [CoursePlanStatusEnum.DRAFT]: "Draft",
    [CoursePlanStatusEnum.APPROVED]: "Approved",
    [CoursePlanStatusEnum.PARTIALLY_APPROVED]: "Partially Approved",
    [CoursePlanStatusEnum.REJECTED]: "Rejected",
    [CoursePlanStatusEnum.SUBMITTED]: "Submitted",
};
