import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { CrossDimer, CrossDimerCreateData } from '../SeqbenchMcpTypes';
declare class CrossDimerEntity extends SeqbenchMcpEntityBase<CrossDimer> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: CrossDimerEntity): CrossDimerEntity;
    create(this: any, reqdata?: CrossDimerCreateData, ctrl?: Control): Promise<CrossDimerEntity>;
}
export { CrossDimerEntity };
