import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { HgvsConvert, HgvsConvertCreateData } from '../SeqbenchMcpTypes';
declare class HgvsConvertEntity extends SeqbenchMcpEntityBase<HgvsConvert> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: HgvsConvertEntity): HgvsConvertEntity;
    create(this: any, reqdata?: HgvsConvertCreateData, ctrl?: Control): Promise<HgvsConvertEntity>;
}
export { HgvsConvertEntity };
