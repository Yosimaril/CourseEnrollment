import { z } from "zod";
import { BaseTimestampDto } from "./baseTimestampDto.ts";
import { CoursePlanItemStatusEnum } from "../constants/coursePlanItemStatusEnum";

export const CoursePlanItemDto = BaseTimestampDto.extend({
    course_plan_id: z.number(),
    course_id: z.number(),
    status: z.enum(CoursePlanItemStatusEnum),
    remarks: z.string().optional(),
});

export type CoursePlanItemDtoType = z.infer<typeof CoursePlanItemDto>;
