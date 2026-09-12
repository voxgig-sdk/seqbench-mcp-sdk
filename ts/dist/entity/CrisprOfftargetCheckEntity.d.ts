import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { CrisprOfftargetCheck, CrisprOfftargetCheckCreateData } from '../SeqbenchMcpTypes';
declare class CrisprOfftargetCheckEntity extends SeqbenchMcpEntityBase<CrisprOfftargetCheck> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: CrisprOfftargetCheckEntity): CrisprOfftargetCheckEntity;
    create(this: any, reqdata?: CrisprOfftargetCheckCreateData, ctrl?: Control): Promise<CrisprOfftargetCheckEntity>;
}
export { CrisprOfftargetCheckEntity };
