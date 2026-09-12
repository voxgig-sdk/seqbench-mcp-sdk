import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { Batch, BatchLoadMatch, BatchCreateData } from '../SeqbenchMcpTypes';
declare class BatchEntity extends SeqbenchMcpEntityBase<Batch> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: BatchEntity): BatchEntity;
    load(this: any, reqmatch?: BatchLoadMatch, ctrl?: Control): Promise<BatchEntity>;
    create(this: any, reqdata?: BatchCreateData, ctrl?: Control): Promise<BatchEntity>;
}
export { BatchEntity };
