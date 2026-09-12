import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SirnaDesign, SirnaDesignCreateData } from '../SeqbenchMcpTypes';
declare class SirnaDesignEntity extends SeqbenchMcpEntityBase<SirnaDesign> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SirnaDesignEntity): SirnaDesignEntity;
    create(this: any, reqdata?: SirnaDesignCreateData, ctrl?: Control): Promise<SirnaDesignEntity>;
}
export { SirnaDesignEntity };
