import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { PlasmidDeepAnnotate, PlasmidDeepAnnotateCreateData } from '../SeqbenchMcpTypes';
declare class PlasmidDeepAnnotateEntity extends SeqbenchMcpEntityBase<PlasmidDeepAnnotate> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: PlasmidDeepAnnotateEntity): PlasmidDeepAnnotateEntity;
    create(this: any, reqdata?: PlasmidDeepAnnotateCreateData, ctrl?: Control): Promise<PlasmidDeepAnnotateEntity>;
}
export { PlasmidDeepAnnotateEntity };
