import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ConstructQc, ConstructQcCreateData } from '../SeqbenchMcpTypes';
declare class ConstructQcEntity extends SeqbenchMcpEntityBase<ConstructQc> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ConstructQcEntity): ConstructQcEntity;
    create(this: any, reqdata?: ConstructQcCreateData, ctrl?: Control): Promise<ConstructQcEntity>;
}
export { ConstructQcEntity };
