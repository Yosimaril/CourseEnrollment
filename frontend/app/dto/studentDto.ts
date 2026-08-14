import { z } from "zod";
import { BaseIdDto } from "./baseIdDto";
import { BaseTimestampDto } from "./baseTimestampDto";
import { CoursePlanDtoSchema } from "./coursePlanDto";

export const StudentDtoSchema = BaseIdDto.merge(BaseTimestampDto).extend({
    username: z.string(),
    email: z.email().optional(),
    password: z.string().nullable().optional(),
    nrp: z.string(),
    max_credits: z.number().optional(),
    course_plans: z.array(CoursePlanDtoSchema).optional(),
});

export type StudentDto = z.infer<typeof StudentDtoSchema>;
