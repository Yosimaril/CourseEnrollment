import type { CoursePlanItem } from "~/models/coursePlanItem";
import type { CoursePlanItemDto } from "~/dto/coursePlanItemDto";
import { CourseMapper } from "./courseMapper";

export class CoursePlanItemMapper {
    static toModel(dto: CoursePlanItemDto): CoursePlanItem {
        return {
            coursePlanId: dto.course_plan_id,
            courseId: dto.course_id,
            status: dto.status,
            remarks: dto.remarks ?? null,
            course: dto.course ? CourseMapper.toModel(dto.course) : undefined,
            createdAt: new Date(dto.created_at),
            updatedAt: new Date(dto.updated_at),
            deletedAt: dto.deleted_at ? new Date(dto.deleted_at) : null,
        };
    }

    static toDto(model: CoursePlanItem): CoursePlanItemDto {
        return {
            course_plan_id: model.coursePlanId,
            course_id: model.courseId,
            status: model.status,
            remarks: model.remarks ?? null,
            course: model.course ? CourseMapper.toDto(model.course) : undefined,
            created_at: model.createdAt.toISOString(),
            updated_at: model.updatedAt.toISOString(),
            deleted_at: model.deletedAt ? model.deletedAt.toISOString() : null,
        };
    }
}
