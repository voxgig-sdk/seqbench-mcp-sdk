import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { GeneModel, GeneModelCreateData } from '../SeqbenchMcpTypes';
declare class GeneModelEntity extends SeqbenchMcpEntityBase<GeneModel> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: GeneModelEntity): GeneModelEntity;
    create(this: any, reqdata?: GeneModelCreateData, ctrl?: Control): Promise<GeneModelEntity>;
}
export { GeneModelEntity };
