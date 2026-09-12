import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SeqfileStat, SeqfileStatCreateData } from '../SeqbenchMcpTypes';
declare class SeqfileStatEntity extends SeqbenchMcpEntityBase<SeqfileStat> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SeqfileStatEntity): SeqfileStatEntity;
    create(this: any, reqdata?: SeqfileStatCreateData, ctrl?: Control): Promise<SeqfileStatEntity>;
}
export { SeqfileStatEntity };
