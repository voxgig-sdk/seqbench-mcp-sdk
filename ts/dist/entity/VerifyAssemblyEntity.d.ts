import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { VerifyAssembly, VerifyAssemblyCreateData } from '../SeqbenchMcpTypes';
declare class VerifyAssemblyEntity extends SeqbenchMcpEntityBase<VerifyAssembly> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: VerifyAssemblyEntity): VerifyAssemblyEntity;
    create(this: any, reqdata?: VerifyAssemblyCreateData, ctrl?: Control): Promise<VerifyAssemblyEntity>;
}
export { VerifyAssemblyEntity };
