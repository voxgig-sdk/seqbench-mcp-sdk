import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { MotifFinder, MotifFinderCreateData } from '../SeqbenchMcpTypes';
declare class MotifFinderEntity extends SeqbenchMcpEntityBase<MotifFinder> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: MotifFinderEntity): MotifFinderEntity;
    create(this: any, reqdata?: MotifFinderCreateData, ctrl?: Control): Promise<MotifFinderEntity>;
}
export { MotifFinderEntity };
