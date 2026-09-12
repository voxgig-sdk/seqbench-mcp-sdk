import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SequenceFetch, SequenceFetchCreateData } from '../SeqbenchMcpTypes';
declare class SequenceFetchEntity extends SeqbenchMcpEntityBase<SequenceFetch> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SequenceFetchEntity): SequenceFetchEntity;
    create(this: any, reqdata?: SequenceFetchCreateData, ctrl?: Control): Promise<SequenceFetchEntity>;
}
export { SequenceFetchEntity };
