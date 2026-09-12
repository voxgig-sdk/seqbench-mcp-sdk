import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ExportPlateLayout, ExportPlateLayoutCreateData } from '../SeqbenchMcpTypes';
declare class ExportPlateLayoutEntity extends SeqbenchMcpEntityBase<ExportPlateLayout> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ExportPlateLayoutEntity): ExportPlateLayoutEntity;
    create(this: any, reqdata?: ExportPlateLayoutCreateData, ctrl?: Control): Promise<ExportPlateLayoutEntity>;
}
export { ExportPlateLayoutEntity };
