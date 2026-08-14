import { z } from "zod";
import { BaseIdDto } from "./baseIdDto";
import { BaseTimestampDto } from "./baseTimestampDto";
import { CoursePlanStatusEnum } from "../constants/coursePlanStatusEnum";
import { UserDtoSchema } from "./userDto";
import { CoursePlanItemDtoSchema } from "./coursePlanItemDto";

export const CoursePlanDtoSchema = BaseIdDto.merge(BaseTimestampDto).extend({
    student_id: z.number(),
    status: z.enum(CoursePlanStatusEnum),
    student: UserDtoSchema.optional(),
    course_plan_items: z.array(CoursePlanItemDtoSchema),
});

export type CoursePlanDto = z.infer<typeof CoursePlanDtoSchema>;
