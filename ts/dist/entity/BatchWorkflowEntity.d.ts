import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { BatchWorkflow, BatchWorkflowLoadMatch, BatchWorkflowCreateData } from '../SeqbenchMcpTypes';
declare class BatchWorkflowEntity extends SeqbenchMcpEntityBase<BatchWorkflow> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: BatchWorkflowEntity): BatchWorkflowEntity;
    load(this: any, reqmatch?: BatchWorkflowLoadMatch, ctrl?: Control): Promise<BatchWorkflowEntity>;
    create(this: any, reqdata?: BatchWorkflowCreateData, ctrl?: Control): Promise<BatchWorkflowEntity>;
}
export { BatchWorkflowEntity };
