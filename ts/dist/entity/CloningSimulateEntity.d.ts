import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { CloningSimulate, CloningSimulateCreateData } from '../SeqbenchMcpTypes';
declare class CloningSimulateEntity extends SeqbenchMcpEntityBase<CloningSimulate> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: CloningSimulateEntity): CloningSimulateEntity;
    create(this: any, reqdata?: CloningSimulateCreateData, ctrl?: Control): Promise<CloningSimulateEntity>;
}
export { CloningSimulateEntity };
