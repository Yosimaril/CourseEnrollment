import type { BaseTimestamp } from "./baseTimestamp";
import { CoursePlanItemStatusEnum } from "../constants/coursePlanItemStatusEnum";
import type { Course } from "./course";

export interface CoursePlanItem extends BaseTimestamp {
    coursePlanId: number;
    courseId: number;
    status: CoursePlanItemStatusEnum;
    remarks?: string | null;
    course?: Course;
}
