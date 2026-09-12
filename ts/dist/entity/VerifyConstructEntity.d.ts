import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { VerifyConstruct, VerifyConstructCreateData } from '../SeqbenchMcpTypes';
declare class VerifyConstructEntity extends SeqbenchMcpEntityBase<VerifyConstruct> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: VerifyConstructEntity): VerifyConstructEntity;
    create(this: any, reqdata?: VerifyConstructCreateData, ctrl?: Control): Promise<VerifyConstructEntity>;
}
export { VerifyConstructEntity };
