import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { OrthologMap, OrthologMapCreateData } from '../SeqbenchMcpTypes';
declare class OrthologMapEntity extends SeqbenchMcpEntityBase<OrthologMap> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: OrthologMapEntity): OrthologMapEntity;
    create(this: any, reqdata?: OrthologMapCreateData, ctrl?: Control): Promise<OrthologMapEntity>;
}
export { OrthologMapEntity };
