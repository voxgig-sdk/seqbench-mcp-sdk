import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { PrimeEditingDesign, PrimeEditingDesignCreateData } from '../SeqbenchMcpTypes';
declare class PrimeEditingDesignEntity extends SeqbenchMcpEntityBase<PrimeEditingDesign> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: PrimeEditingDesignEntity): PrimeEditingDesignEntity;
    create(this: any, reqdata?: PrimeEditingDesignCreateData, ctrl?: Control): Promise<PrimeEditingDesignEntity>;
}
export { PrimeEditingDesignEntity };
