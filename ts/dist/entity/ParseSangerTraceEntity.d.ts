import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ParseSangerTrace, ParseSangerTraceCreateData } from '../SeqbenchMcpTypes';
declare class ParseSangerTraceEntity extends SeqbenchMcpEntityBase<ParseSangerTrace> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ParseSangerTraceEntity): ParseSangerTraceEntity;
    create(this: any, reqdata?: ParseSangerTraceCreateData, ctrl?: Control): Promise<ParseSangerTraceEntity>;
}
export { ParseSangerTraceEntity };
