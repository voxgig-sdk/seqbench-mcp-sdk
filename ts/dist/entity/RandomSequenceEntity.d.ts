import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { RandomSequence, RandomSequenceCreateData } from '../SeqbenchMcpTypes';
declare class RandomSequenceEntity extends SeqbenchMcpEntityBase<RandomSequence> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: RandomSequenceEntity): RandomSequenceEntity;
    create(this: any, reqdata?: RandomSequenceCreateData, ctrl?: Control): Promise<RandomSequenceEntity>;
}
export { RandomSequenceEntity };
