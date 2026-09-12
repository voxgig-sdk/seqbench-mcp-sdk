import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { BaseEditingDesign, BaseEditingDesignCreateData } from '../SeqbenchMcpTypes';
declare class BaseEditingDesignEntity extends SeqbenchMcpEntityBase<BaseEditingDesign> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: BaseEditingDesignEntity): BaseEditingDesignEntity;
    create(this: any, reqdata?: BaseEditingDesignCreateData, ctrl?: Control): Promise<BaseEditingDesignEntity>;
}
export { BaseEditingDesignEntity };
