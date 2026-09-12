import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ReverseComplement, ReverseComplementCreateData } from '../SeqbenchMcpTypes';
declare class ReverseComplementEntity extends SeqbenchMcpEntityBase<ReverseComplement> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ReverseComplementEntity): ReverseComplementEntity;
    create(this: any, reqdata?: ReverseComplementCreateData, ctrl?: Control): Promise<ReverseComplementEntity>;
}
export { ReverseComplementEntity };
