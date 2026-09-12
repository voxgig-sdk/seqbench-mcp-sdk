import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { DoubleDigest, DoubleDigestCreateData } from '../SeqbenchMcpTypes';
declare class DoubleDigestEntity extends SeqbenchMcpEntityBase<DoubleDigest> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: DoubleDigestEntity): DoubleDigestEntity;
    create(this: any, reqdata?: DoubleDigestCreateData, ctrl?: Control): Promise<DoubleDigestEntity>;
}
export { DoubleDigestEntity };
