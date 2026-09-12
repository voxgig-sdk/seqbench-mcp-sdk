import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { RnaFold, RnaFoldCreateData } from '../SeqbenchMcpTypes';
declare class RnaFoldEntity extends SeqbenchMcpEntityBase<RnaFold> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: RnaFoldEntity): RnaFoldEntity;
    create(this: any, reqdata?: RnaFoldCreateData, ctrl?: Control): Promise<RnaFoldEntity>;
}
export { RnaFoldEntity };
