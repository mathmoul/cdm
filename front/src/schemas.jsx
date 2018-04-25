import { schema } from 'normalizr'

export const teamsSchema = new schema.Entity(
    "teams",
    {},
    {"idAttribute": "_id"}
);