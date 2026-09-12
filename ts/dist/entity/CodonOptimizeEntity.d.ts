import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { CodonOptimize, CodonOptimizeCreateData } from '../SeqbenchMcpTypes';
declare class CodonOptimizeEntity extends SeqbenchMcpEntityBase<CodonOptimize> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: CodonOptimizeEntity): CodonOptimizeEntity;
    create(this: any, reqdata?: CodonOptimizeCreateData, ctrl?: Control): Promise<CodonOptimizeEntity>;
}
export { CodonOptimizeEntity };
