import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { GeneExpression, GeneExpressionCreateData } from '../SeqbenchMcpTypes';
declare class GeneExpressionEntity extends SeqbenchMcpEntityBase<GeneExpression> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: GeneExpressionEntity): GeneExpressionEntity;
    create(this: any, reqdata?: GeneExpressionCreateData, ctrl?: Control): Promise<GeneExpressionEntity>;
}
export { GeneExpressionEntity };
