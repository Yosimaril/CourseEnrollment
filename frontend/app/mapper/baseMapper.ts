import type { BaseIdDto } from "~/dto/baseIdDto"
import type { BaseTimestampDto } from "~/dto/baseTimestampDto"

export abstract class BaseMapper<Raw, Dto> {
    protected mapId(raw: BaseIdDto) {
        return {
            id: raw.id,
        };
    }

    protected mapTimestamps(raw: BaseTimestampDto) {
        return {
            createdAt: new Date(raw.created_at),
            updatedAt: new Date(raw.updated_at),
            deletedAt: new Date(raw.deleted_at),
        };
    }
}
