import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { PairwiseAlignment, PairwiseAlignmentCreateData } from '../SeqbenchMcpTypes';
declare class PairwiseAlignmentEntity extends SeqbenchMcpEntityBase<PairwiseAlignment> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: PairwiseAlignmentEntity): PairwiseAlignmentEntity;
    create(this: any, reqdata?: PairwiseAlignmentCreateData, ctrl?: Control): Promise<PairwiseAlignmentEntity>;
}
export { PairwiseAlignmentEntity };
