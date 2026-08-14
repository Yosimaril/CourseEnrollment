import { z } from "zod";
import { BaseIdDto } from "./baseIdDto";
import { BaseTimestampDto } from "./baseTimestampDto";
import { UserRoleEnum } from "../constants/userRoleEnum";

export const UserDtoSchema = BaseIdDto.merge(BaseTimestampDto).extend({
    username: z.string(),
    email: z.email(),
    password: z.string().optional(),
    role: z.enum(UserRoleEnum),
    nrp: z.string().nullable().optional(),
    max_credits: z.number().nullable().optional(),
});

export type UserDto = z.infer<typeof UserDtoSchema>;
