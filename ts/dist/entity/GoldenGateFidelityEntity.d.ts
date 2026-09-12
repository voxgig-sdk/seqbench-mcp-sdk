import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { GoldenGateFidelity, GoldenGateFidelityCreateData } from '../SeqbenchMcpTypes';
declare class GoldenGateFidelityEntity extends SeqbenchMcpEntityBase<GoldenGateFidelity> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: GoldenGateFidelityEntity): GoldenGateFidelityEntity;
    create(this: any, reqdata?: GoldenGateFidelityCreateData, ctrl?: Control): Promise<GoldenGateFidelityEntity>;
}
export { GoldenGateFidelityEntity };
