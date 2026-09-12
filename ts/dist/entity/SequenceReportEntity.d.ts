import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SequenceReport, SequenceReportCreateData } from '../SeqbenchMcpTypes';
declare class SequenceReportEntity extends SeqbenchMcpEntityBase<SequenceReport> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SequenceReportEntity): SequenceReportEntity;
    create(this: any, reqdata?: SequenceReportCreateData, ctrl?: Control): Promise<SequenceReportEntity>;
}
export { SequenceReportEntity };
