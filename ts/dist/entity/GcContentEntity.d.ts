import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { GcContent, GcContentCreateData } from '../SeqbenchMcpTypes';
declare class GcContentEntity extends SeqbenchMcpEntityBase<GcContent> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: GcContentEntity): GcContentEntity;
    create(this: any, reqdata?: GcContentCreateData, ctrl?: Control): Promise<GcContentEntity>;
}
export { GcContentEntity };
