import { z } from "zod";
import { BaseIdDto } from "./baseIdDto.ts";
import { BaseTimestampDto } from "./baseTimestampDto.ts";

export const StudentDto = BaseIdDto
    .extend(BaseTimestampDto)
    .extend({
        username: z.string(),
        email: z.string().email(),
        password: z.string().optional(),
        nrp: z.string().optional(),
        max_credits: z.number().optional(),
    });

export type StudentDtoType = z.infer<typeof StudentDto>;