import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SequencingReadbackVerify, SequencingReadbackVerifyCreateData } from '../SeqbenchMcpTypes';
declare class SequencingReadbackVerifyEntity extends SeqbenchMcpEntityBase<SequencingReadbackVerify> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SequencingReadbackVerifyEntity): SequencingReadbackVerifyEntity;
    create(this: any, reqdata?: SequencingReadbackVerifyCreateData, ctrl?: Control): Promise<SequencingReadbackVerifyEntity>;
}
export { SequencingReadbackVerifyEntity };
