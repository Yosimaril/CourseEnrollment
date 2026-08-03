import type { BaseIdDto } from "./baseIdDto.ts";
import type { BaseTimestampDto } from "./baseTimestampDto.ts";

export interface StudentDto extends BaseIdDto, BaseTimestampDto {
    username: string;
    email: string;
    password?: string;
    nrp?: string;
    max_credits?: number;
}
