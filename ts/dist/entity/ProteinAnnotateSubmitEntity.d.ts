import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ProteinAnnotateSubmit, ProteinAnnotateSubmitCreateData } from '../SeqbenchMcpTypes';
declare class ProteinAnnotateSubmitEntity extends SeqbenchMcpEntityBase<ProteinAnnotateSubmit> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ProteinAnnotateSubmitEntity): ProteinAnnotateSubmitEntity;
    create(this: any, reqdata?: ProteinAnnotateSubmitCreateData, ctrl?: Control): Promise<ProteinAnnotateSubmitEntity>;
}
export { ProteinAnnotateSubmitEntity };
