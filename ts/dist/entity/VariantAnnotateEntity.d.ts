import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { VariantAnnotate, VariantAnnotateCreateData } from '../SeqbenchMcpTypes';
declare class VariantAnnotateEntity extends SeqbenchMcpEntityBase<VariantAnnotate> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: VariantAnnotateEntity): VariantAnnotateEntity;
    create(this: any, reqdata?: VariantAnnotateCreateData, ctrl?: Control): Promise<VariantAnnotateEntity>;
}
export { VariantAnnotateEntity };
