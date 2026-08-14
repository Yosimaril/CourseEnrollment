import { z } from "zod";
import { BaseTimestampDto } from "./baseTimestampDto";
import { CoursePlanItemStatusEnum } from "../constants/coursePlanItemStatusEnum";
import { CourseDtoSchema } from "./courseDto";

export const CoursePlanItemDtoSchema = BaseTimestampDto.extend({
    course_plan_id: z.number(),
    course_id: z.number(),
    status: z.enum(CoursePlanItemStatusEnum),
    remarks: z.string().nullable().optional(),
    course: CourseDtoSchema.optional(),
});

export type CoursePlanItemDto = z.infer<typeof CoursePlanItemDtoSchema>;
