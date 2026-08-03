import type { BaseTimestampDto } from "./baseTimestampDto.ts";
import type { CoursePlanItemStatusEnum } from "../constants/coursePlanItemStatusEnum";

export interface CoursePlanItemDto extends BaseTimestampDto {
    course_plan_id: number;
    course_id: number;
    status: CoursePlanItemStatusEnum;
    remarks?: string;
}
