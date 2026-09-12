import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { DnaMolarity, DnaMolarityCreateData } from '../SeqbenchMcpTypes';
declare class DnaMolarityEntity extends SeqbenchMcpEntityBase<DnaMolarity> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: DnaMolarityEntity): DnaMolarityEntity;
    create(this: any, reqdata?: DnaMolarityCreateData, ctrl?: Control): Promise<DnaMolarityEntity>;
}
export { DnaMolarityEntity };
