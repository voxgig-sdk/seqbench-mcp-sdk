import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { PlasmidFullReport, PlasmidFullReportCreateData } from '../SeqbenchMcpTypes';
declare class PlasmidFullReportEntity extends SeqbenchMcpEntityBase<PlasmidFullReport> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: PlasmidFullReportEntity): PlasmidFullReportEntity;
    create(this: any, reqdata?: PlasmidFullReportCreateData, ctrl?: Control): Promise<PlasmidFullReportEntity>;
}
export { PlasmidFullReportEntity };
