import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { AsoDesign, AsoDesignCreateData } from '../SeqbenchMcpTypes';
declare class AsoDesignEntity extends SeqbenchMcpEntityBase<AsoDesign> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: AsoDesignEntity): AsoDesignEntity;
    create(this: any, reqdata?: AsoDesignCreateData, ctrl?: Control): Promise<AsoDesignEntity>;
}
export { AsoDesignEntity };
