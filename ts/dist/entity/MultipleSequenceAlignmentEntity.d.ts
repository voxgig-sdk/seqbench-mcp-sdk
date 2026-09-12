import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { MultipleSequenceAlignment, MultipleSequenceAlignmentCreateData } from '../SeqbenchMcpTypes';
declare class MultipleSequenceAlignmentEntity extends SeqbenchMcpEntityBase<MultipleSequenceAlignment> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: MultipleSequenceAlignmentEntity): MultipleSequenceAlignmentEntity;
    create(this: any, reqdata?: MultipleSequenceAlignmentCreateData, ctrl?: Control): Promise<MultipleSequenceAlignmentEntity>;
}
export { MultipleSequenceAlignmentEntity };
