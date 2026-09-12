import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { PlasmidAnnotate, PlasmidAnnotateCreateData } from '../SeqbenchMcpTypes';
declare class PlasmidAnnotateEntity extends SeqbenchMcpEntityBase<PlasmidAnnotate> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: PlasmidAnnotateEntity): PlasmidAnnotateEntity;
    create(this: any, reqdata?: PlasmidAnnotateCreateData, ctrl?: Control): Promise<PlasmidAnnotateEntity>;
}
export { PlasmidAnnotateEntity };
