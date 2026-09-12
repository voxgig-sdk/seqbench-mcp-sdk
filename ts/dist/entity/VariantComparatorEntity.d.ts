import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { VariantComparator, VariantComparatorCreateData } from '../SeqbenchMcpTypes';
declare class VariantComparatorEntity extends SeqbenchMcpEntityBase<VariantComparator> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: VariantComparatorEntity): VariantComparatorEntity;
    create(this: any, reqdata?: VariantComparatorCreateData, ctrl?: Control): Promise<VariantComparatorEntity>;
}
export { VariantComparatorEntity };
