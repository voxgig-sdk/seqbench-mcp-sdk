import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { IdMapPoll, IdMapPollCreateData } from '../SeqbenchMcpTypes';
declare class IdMapPollEntity extends SeqbenchMcpEntityBase<IdMapPoll> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: IdMapPollEntity): IdMapPollEntity;
    create(this: any, reqdata?: IdMapPollCreateData, ctrl?: Control): Promise<IdMapPollEntity>;
}
export { IdMapPollEntity };
