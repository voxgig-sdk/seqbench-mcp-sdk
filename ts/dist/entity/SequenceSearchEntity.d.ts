import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SequenceSearch, SequenceSearchCreateData } from '../SeqbenchMcpTypes';
declare class SequenceSearchEntity extends SeqbenchMcpEntityBase<SequenceSearch> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SequenceSearchEntity): SequenceSearchEntity;
    create(this: any, reqdata?: SequenceSearchCreateData, ctrl?: Control): Promise<SequenceSearchEntity>;
}
export { SequenceSearchEntity };
