import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ProteaseDigestion, ProteaseDigestionCreateData } from '../SeqbenchMcpTypes';
declare class ProteaseDigestionEntity extends SeqbenchMcpEntityBase<ProteaseDigestion> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ProteaseDigestionEntity): ProteaseDigestionEntity;
    create(this: any, reqdata?: ProteaseDigestionCreateData, ctrl?: Control): Promise<ProteaseDigestionEntity>;
}
export { ProteaseDigestionEntity };
