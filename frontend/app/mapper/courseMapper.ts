import type { BaseIdDto } from "~/dto/baseIdDto";
import type { BaseTimestampDto } from "~/dto/baseTimestampDto";

export type RawCourse = BaseIdDto &
    BaseTimestampDto & {
        code: string;
        name: string;
        credits: number;
    };

export class CourseMapper {
    static fromJson(json: unknown): RawCourse {
        return json as RawCourse;
    }

    static toJson(dto: RawCourse): object {
        return {
            id: dto.id,
            code: dto.code,
            name: dto.name,
            credits: dto.credits,
            created_at: dto.created_at,
            updated_at: dto.updated_at,
        };
    }
}