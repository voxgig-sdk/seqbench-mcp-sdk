import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { GeneDossier, GeneDossierCreateData } from '../SeqbenchMcpTypes';
declare class GeneDossierEntity extends SeqbenchMcpEntityBase<GeneDossier> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: GeneDossierEntity): GeneDossierEntity;
    create(this: any, reqdata?: GeneDossierCreateData, ctrl?: Control): Promise<GeneDossierEntity>;
}
export { GeneDossierEntity };
