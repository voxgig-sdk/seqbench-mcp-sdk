import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SiteDirectedMutagenesi, SiteDirectedMutagenesiCreateData } from '../SeqbenchMcpTypes';
declare class SiteDirectedMutagenesiEntity extends SeqbenchMcpEntityBase<SiteDirectedMutagenesi> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SiteDirectedMutagenesiEntity): SiteDirectedMutagenesiEntity;
    create(this: any, reqdata?: SiteDirectedMutagenesiCreateData, ctrl?: Control): Promise<SiteDirectedMutagenesiEntity>;
}
export { SiteDirectedMutagenesiEntity };
