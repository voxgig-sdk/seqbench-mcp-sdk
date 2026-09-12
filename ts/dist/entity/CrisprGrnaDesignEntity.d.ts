import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { CrisprGrnaDesign, CrisprGrnaDesignCreateData } from '../SeqbenchMcpTypes';
declare class CrisprGrnaDesignEntity extends SeqbenchMcpEntityBase<CrisprGrnaDesign> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: CrisprGrnaDesignEntity): CrisprGrnaDesignEntity;
    create(this: any, reqdata?: CrisprGrnaDesignCreateData, ctrl?: Control): Promise<CrisprGrnaDesignEntity>;
}
export { CrisprGrnaDesignEntity };
