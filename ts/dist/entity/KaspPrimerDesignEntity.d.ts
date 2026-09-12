import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { KaspPrimerDesign, KaspPrimerDesignCreateData } from '../SeqbenchMcpTypes';
declare class KaspPrimerDesignEntity extends SeqbenchMcpEntityBase<KaspPrimerDesign> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: KaspPrimerDesignEntity): KaspPrimerDesignEntity;
    create(this: any, reqdata?: KaspPrimerDesignCreateData, ctrl?: Control): Promise<KaspPrimerDesignEntity>;
}
export { KaspPrimerDesignEntity };
