import { CoursePlanItemStatusEnum } from "~/constants/coursePlanItemStatusEnum";

export const CoursePlanItemStatusMapper = {
    [CoursePlanItemStatusEnum.PENDING]: {
        label: "Pending",
        color: "yellow",
    },
    [CoursePlanItemStatusEnum.APPROVED]: {
        label: "Approved",
        color: "green",
    },
    [CoursePlanItemStatusEnum.REJECTED]: {
        label: "Rejected",
        color: "red",
    },
};
