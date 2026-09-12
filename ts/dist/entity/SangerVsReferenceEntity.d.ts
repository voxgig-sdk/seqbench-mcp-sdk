import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SangerVsReference, SangerVsReferenceCreateData } from '../SeqbenchMcpTypes';
declare class SangerVsReferenceEntity extends SeqbenchMcpEntityBase<SangerVsReference> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SangerVsReferenceEntity): SangerVsReferenceEntity;
    create(this: any, reqdata?: SangerVsReferenceCreateData, ctrl?: Control): Promise<SangerVsReferenceEntity>;
}
export { SangerVsReferenceEntity };
