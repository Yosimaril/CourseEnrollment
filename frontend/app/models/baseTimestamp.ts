export interface BaseTimestamp {
    createdAt: Date;
    updatedAt: Date;
    deletedAt?: Date | null;
}
