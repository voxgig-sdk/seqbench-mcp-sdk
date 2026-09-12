import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ConstructAutofix, ConstructAutofixCreateData } from '../SeqbenchMcpTypes';
declare class ConstructAutofixEntity extends SeqbenchMcpEntityBase<ConstructAutofix> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ConstructAutofixEntity): ConstructAutofixEntity;
    create(this: any, reqdata?: ConstructAutofixCreateData, ctrl?: Control): Promise<ConstructAutofixEntity>;
}
export { ConstructAutofixEntity };
