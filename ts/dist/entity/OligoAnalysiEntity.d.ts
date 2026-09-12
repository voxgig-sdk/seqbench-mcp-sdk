import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { OligoAnalysi, OligoAnalysiCreateData } from '../SeqbenchMcpTypes';
declare class OligoAnalysiEntity extends SeqbenchMcpEntityBase<OligoAnalysi> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: OligoAnalysiEntity): OligoAnalysiEntity;
    create(this: any, reqdata?: OligoAnalysiCreateData, ctrl?: Control): Promise<OligoAnalysiEntity>;
}
export { OligoAnalysiEntity };
