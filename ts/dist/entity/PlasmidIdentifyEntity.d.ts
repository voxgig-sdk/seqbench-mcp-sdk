import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { PlasmidIdentify, PlasmidIdentifyCreateData } from '../SeqbenchMcpTypes';
declare class PlasmidIdentifyEntity extends SeqbenchMcpEntityBase<PlasmidIdentify> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: PlasmidIdentifyEntity): PlasmidIdentifyEntity;
    create(this: any, reqdata?: PlasmidIdentifyCreateData, ctrl?: Control): Promise<PlasmidIdentifyEntity>;
}
export { PlasmidIdentifyEntity };
