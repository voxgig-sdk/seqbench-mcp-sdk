import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ProteinAnnotatePoll, ProteinAnnotatePollCreateData } from '../SeqbenchMcpTypes';
declare class ProteinAnnotatePollEntity extends SeqbenchMcpEntityBase<ProteinAnnotatePoll> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ProteinAnnotatePollEntity): ProteinAnnotatePollEntity;
    create(this: any, reqdata?: ProteinAnnotatePollCreateData, ctrl?: Control): Promise<ProteinAnnotatePollEntity>;
}
export { ProteinAnnotatePollEntity };
