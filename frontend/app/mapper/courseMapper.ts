import type { Course } from "~/models/course";
import type { CourseDto } from "~/dto/courseDto";

export class CourseMapper {
    static toModel(dto: CourseDto): Course {
        return {
            id: dto.id,
            code: dto.code,
            name: dto.name,
            credits: dto.credits,
            createdAt: new Date(dto.created_at),
            updatedAt: new Date(dto.updated_at),
            deletedAt: dto.deleted_at ? new Date(dto.deleted_at) : null,
        };
    }

    static toDto(model: Course): CourseDto {
        return {
            id: model.id,
            code: model.code,
            name: model.name,
            credits: model.credits,
            created_at: model.createdAt.toISOString(),
            updated_at: model.updatedAt.toISOString(),
            deleted_at: model.deletedAt ? model.deletedAt.toISOString() : null,
        };
    }
}
