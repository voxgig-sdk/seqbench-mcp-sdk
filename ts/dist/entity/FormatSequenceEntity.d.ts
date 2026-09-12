import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { FormatSequence, FormatSequenceCreateData } from '../SeqbenchMcpTypes';
declare class FormatSequenceEntity extends SeqbenchMcpEntityBase<FormatSequence> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: FormatSequenceEntity): FormatSequenceEntity;
    create(this: any, reqdata?: FormatSequenceCreateData, ctrl?: Control): Promise<FormatSequenceEntity>;
}
export { FormatSequenceEntity };
