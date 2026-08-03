import type { BaseIdDto } from "./baseIdDto.ts";
import type { BaseTimestampDto } from "./baseTimestampDto.ts";
import type { UserRoleEnum } from "../constants/userRoleEnum";

export interface UserDto extends BaseIdDto, BaseTimestampDto {
    username: string;
    email: string;
    password?: string;
    role: UserRoleEnum;
    nrp?: string;
    max_credits?: number;
}
