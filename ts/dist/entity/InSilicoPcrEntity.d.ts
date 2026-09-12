import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { InSilicoPcr, InSilicoPcrCreateData } from '../SeqbenchMcpTypes';
declare class InSilicoPcrEntity extends SeqbenchMcpEntityBase<InSilicoPcr> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: InSilicoPcrEntity): InSilicoPcrEntity;
    create(this: any, reqdata?: InSilicoPcrCreateData, ctrl?: Control): Promise<InSilicoPcrEntity>;
}
export { InSilicoPcrEntity };
