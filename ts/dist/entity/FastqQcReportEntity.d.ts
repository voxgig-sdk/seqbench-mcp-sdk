import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { FastqQcReport, FastqQcReportCreateData } from '../SeqbenchMcpTypes';
declare class FastqQcReportEntity extends SeqbenchMcpEntityBase<FastqQcReport> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: FastqQcReportEntity): FastqQcReportEntity;
    create(this: any, reqdata?: FastqQcReportCreateData, ctrl?: Control): Promise<FastqQcReportEntity>;
}
export { FastqQcReportEntity };
