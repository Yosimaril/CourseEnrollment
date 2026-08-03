import type { BaseIdDto } from "./baseIdDto.ts";
import type { BaseTimestampDto } from "./baseTimestampDto.ts";
import type { CoursePlanStatusEnum } from "../constants/coursePlanStatusEnum.ts";

export interface CoursePlanDto extends BaseIdDto, BaseTimestampDto {
    student_id: number;
    status: CoursePlanStatusEnum;
}
