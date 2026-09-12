import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ProteinHydrophobicity, ProteinHydrophobicityCreateData } from '../SeqbenchMcpTypes';
declare class ProteinHydrophobicityEntity extends SeqbenchMcpEntityBase<ProteinHydrophobicity> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ProteinHydrophobicityEntity): ProteinHydrophobicityEntity;
    create(this: any, reqdata?: ProteinHydrophobicityCreateData, ctrl?: Control): Promise<ProteinHydrophobicityEntity>;
}
export { ProteinHydrophobicityEntity };
