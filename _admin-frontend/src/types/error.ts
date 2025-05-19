import { duplicateEntry } from "../vars/error";

export class BadRequestError extends Error {
    fields: string[];
    traceId: string;

    constructor(traceId: string, message: string, fields: string[]) {
        super(message);
        this.traceId = traceId;
        this.fields = fields;
        this.name = duplicateEntry
    } 
}