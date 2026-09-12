import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ExpressionHeatmapCluster, ExpressionHeatmapClusterCreateData } from '../SeqbenchMcpTypes';
declare class ExpressionHeatmapClusterEntity extends SeqbenchMcpEntityBase<ExpressionHeatmapCluster> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ExpressionHeatmapClusterEntity): ExpressionHeatmapClusterEntity;
    create(this: any, reqdata?: ExpressionHeatmapClusterCreateData, ctrl?: Control): Promise<ExpressionHeatmapClusterEntity>;
}
export { ExpressionHeatmapClusterEntity };
