import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ProteinProperty, ProteinPropertyCreateData } from '../SeqbenchMcpTypes';
declare class ProteinPropertyEntity extends SeqbenchMcpEntityBase<ProteinProperty> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ProteinPropertyEntity): ProteinPropertyEntity;
    create(this: any, reqdata?: ProteinPropertyCreateData, ctrl?: Control): Promise<ProteinPropertyEntity>;
}
export { ProteinPropertyEntity };
