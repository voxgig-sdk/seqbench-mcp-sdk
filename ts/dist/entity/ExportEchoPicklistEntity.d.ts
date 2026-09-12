import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ExportEchoPicklist, ExportEchoPicklistCreateData } from '../SeqbenchMcpTypes';
declare class ExportEchoPicklistEntity extends SeqbenchMcpEntityBase<ExportEchoPicklist> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ExportEchoPicklistEntity): ExportEchoPicklistEntity;
    create(this: any, reqdata?: ExportEchoPicklistCreateData, ctrl?: Control): Promise<ExportEchoPicklistEntity>;
}
export { ExportEchoPicklistEntity };
