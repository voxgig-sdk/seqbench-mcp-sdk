import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { RestrictionSite, RestrictionSiteCreateData } from '../SeqbenchMcpTypes';
declare class RestrictionSiteEntity extends SeqbenchMcpEntityBase<RestrictionSite> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: RestrictionSiteEntity): RestrictionSiteEntity;
    create(this: any, reqdata?: RestrictionSiteCreateData, ctrl?: Control): Promise<RestrictionSiteEntity>;
}
export { RestrictionSiteEntity };
