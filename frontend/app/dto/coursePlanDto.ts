import { z } from "zod";
import { BaseIdDto } from "./baseIdDto.ts";
import { BaseTimestampDto } from "./baseTimestampDto.ts";
import { CoursePlanStatusEnum } from "../constants/coursePlanStatusEnum.ts";

export const CoursePlanDto = BaseIdDto
    .extend(BaseTimestampDto)
    .extend({
        student_id: z.number(),
        status: z.enum(CoursePlanStatusEnum),
    });

export type CoursePlanDtoType = z.infer<typeof CoursePlanDto>;