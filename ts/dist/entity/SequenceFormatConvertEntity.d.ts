import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SequenceFormatConvert, SequenceFormatConvertCreateData } from '../SeqbenchMcpTypes';
declare class SequenceFormatConvertEntity extends SeqbenchMcpEntityBase<SequenceFormatConvert> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SequenceFormatConvertEntity): SequenceFormatConvertEntity;
    create(this: any, reqdata?: SequenceFormatConvertCreateData, ctrl?: Control): Promise<SequenceFormatConvertEntity>;
}
export { SequenceFormatConvertEntity };
