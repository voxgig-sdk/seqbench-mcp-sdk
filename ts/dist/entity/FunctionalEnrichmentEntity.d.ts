import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { FunctionalEnrichment, FunctionalEnrichmentCreateData } from '../SeqbenchMcpTypes';
declare class FunctionalEnrichmentEntity extends SeqbenchMcpEntityBase<FunctionalEnrichment> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: FunctionalEnrichmentEntity): FunctionalEnrichmentEntity;
    create(this: any, reqdata?: FunctionalEnrichmentCreateData, ctrl?: Control): Promise<FunctionalEnrichmentEntity>;
}
export { FunctionalEnrichmentEntity };
