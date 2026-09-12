import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { IdMapSubmit, IdMapSubmitCreateData } from '../SeqbenchMcpTypes';
declare class IdMapSubmitEntity extends SeqbenchMcpEntityBase<IdMapSubmit> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: IdMapSubmitEntity): IdMapSubmitEntity;
    create(this: any, reqdata?: IdMapSubmitCreateData, ctrl?: Control): Promise<IdMapSubmitEntity>;
}
export { IdMapSubmitEntity };
