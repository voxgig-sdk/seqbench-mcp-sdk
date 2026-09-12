import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { CrisprHdrDonor, CrisprHdrDonorCreateData } from '../SeqbenchMcpTypes';
declare class CrisprHdrDonorEntity extends SeqbenchMcpEntityBase<CrisprHdrDonor> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: CrisprHdrDonorEntity): CrisprHdrDonorEntity;
    create(this: any, reqdata?: CrisprHdrDonorCreateData, ctrl?: Control): Promise<CrisprHdrDonorEntity>;
}
export { CrisprHdrDonorEntity };
