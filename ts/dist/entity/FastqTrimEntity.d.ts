import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { FastqTrim, FastqTrimCreateData } from '../SeqbenchMcpTypes';
declare class FastqTrimEntity extends SeqbenchMcpEntityBase<FastqTrim> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: FastqTrimEntity): FastqTrimEntity;
    create(this: any, reqdata?: FastqTrimCreateData, ctrl?: Control): Promise<FastqTrimEntity>;
}
export { FastqTrimEntity };
