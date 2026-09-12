import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { PrimeEditingTwinDesign, PrimeEditingTwinDesignCreateData } from '../SeqbenchMcpTypes';
declare class PrimeEditingTwinDesignEntity extends SeqbenchMcpEntityBase<PrimeEditingTwinDesign> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: PrimeEditingTwinDesignEntity): PrimeEditingTwinDesignEntity;
    create(this: any, reqdata?: PrimeEditingTwinDesignCreateData, ctrl?: Control): Promise<PrimeEditingTwinDesignEntity>;
}
export { PrimeEditingTwinDesignEntity };
