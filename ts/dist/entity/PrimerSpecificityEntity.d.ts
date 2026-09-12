import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { PrimerSpecificity, PrimerSpecificityCreateData } from '../SeqbenchMcpTypes';
declare class PrimerSpecificityEntity extends SeqbenchMcpEntityBase<PrimerSpecificity> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: PrimerSpecificityEntity): PrimerSpecificityEntity;
    create(this: any, reqdata?: PrimerSpecificityCreateData, ctrl?: Control): Promise<PrimerSpecificityEntity>;
}
export { PrimerSpecificityEntity };
