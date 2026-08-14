import { z } from "zod";
import { BaseIdDto } from "./baseIdDto";
import { BaseTimestampDto } from "./baseTimestampDto";

export const CourseDtoSchema = BaseIdDto.merge(BaseTimestampDto).extend({
    code: z.string(),
    name: z.string(),
    credits: z.number(),
});

export type CourseDto = z.infer<typeof CourseDtoSchema>;
