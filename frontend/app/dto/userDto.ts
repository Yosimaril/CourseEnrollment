import { z } from "zod";
import { BaseIdDto } from "./baseIdDto.js";
import { BaseTimestampDto } from "./baseTimestampDto.ts";
import { UserRoleEnum } from "../constants/userRoleEnum";

export const UserDto = BaseIdDto.extend(BaseTimestampDto).extend({
    username: z.string(),
    email: z.email(),
    password: z.string().optional(),
    role: z.enum(UserRoleEnum),
    nrp: z.string().optional(),
    max_credits: z.number().optional(),
});

export type UserDtoType = z.infer<typeof UserDto>;
