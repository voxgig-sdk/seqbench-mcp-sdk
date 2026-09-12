import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { FindOrf, FindOrfCreateData } from '../SeqbenchMcpTypes';
declare class FindOrfEntity extends SeqbenchMcpEntityBase<FindOrf> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: FindOrfEntity): FindOrfEntity;
    create(this: any, reqdata?: FindOrfCreateData, ctrl?: Control): Promise<FindOrfEntity>;
}
export { FindOrfEntity };
