import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ExportOpentronsProtocol, ExportOpentronsProtocolCreateData } from '../SeqbenchMcpTypes';
declare class ExportOpentronsProtocolEntity extends SeqbenchMcpEntityBase<ExportOpentronsProtocol> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ExportOpentronsProtocolEntity): ExportOpentronsProtocolEntity;
    create(this: any, reqdata?: ExportOpentronsProtocolCreateData, ctrl?: Control): Promise<ExportOpentronsProtocolEntity>;
}
export { ExportOpentronsProtocolEntity };
