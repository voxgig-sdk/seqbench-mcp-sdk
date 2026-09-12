import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { PrimerDesign, PrimerDesignCreateData } from '../SeqbenchMcpTypes';
declare class PrimerDesignEntity extends SeqbenchMcpEntityBase<PrimerDesign> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: PrimerDesignEntity): PrimerDesignEntity;
    create(this: any, reqdata?: PrimerDesignCreateData, ctrl?: Control): Promise<PrimerDesignEntity>;
}
export { PrimerDesignEntity };
