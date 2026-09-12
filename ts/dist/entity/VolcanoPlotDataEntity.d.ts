import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { VolcanoPlotData, VolcanoPlotDataCreateData } from '../SeqbenchMcpTypes';
declare class VolcanoPlotDataEntity extends SeqbenchMcpEntityBase<VolcanoPlotData> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: VolcanoPlotDataEntity): VolcanoPlotDataEntity;
    create(this: any, reqdata?: VolcanoPlotDataCreateData, ctrl?: Control): Promise<VolcanoPlotDataEntity>;
}
export { VolcanoPlotDataEntity };
