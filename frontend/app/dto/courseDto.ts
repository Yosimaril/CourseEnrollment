import type { BaseIdDto } from "./baseIdDto.ts";
import type { BaseTimestampDto } from "./baseTimestampDto.ts";

export interface CourseDto extends BaseIdDto, BaseTimestampDto {
    code: string;
    name: string;
    credits: number;
}
