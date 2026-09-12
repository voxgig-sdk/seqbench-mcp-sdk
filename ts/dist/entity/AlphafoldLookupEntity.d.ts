import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { AlphafoldLookup, AlphafoldLookupCreateData } from '../SeqbenchMcpTypes';
declare class AlphafoldLookupEntity extends SeqbenchMcpEntityBase<AlphafoldLookup> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: AlphafoldLookupEntity): AlphafoldLookupEntity;
    create(this: any, reqdata?: AlphafoldLookupCreateData, ctrl?: Control): Promise<AlphafoldLookupEntity>;
}
export { AlphafoldLookupEntity };
