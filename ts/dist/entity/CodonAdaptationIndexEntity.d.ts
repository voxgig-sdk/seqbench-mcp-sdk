import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { CodonAdaptationIndex, CodonAdaptationIndexCreateData } from '../SeqbenchMcpTypes';
declare class CodonAdaptationIndexEntity extends SeqbenchMcpEntityBase<CodonAdaptationIndex> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: CodonAdaptationIndexEntity): CodonAdaptationIndexEntity;
    create(this: any, reqdata?: CodonAdaptationIndexCreateData, ctrl?: Control): Promise<CodonAdaptationIndexEntity>;
}
export { CodonAdaptationIndexEntity };
