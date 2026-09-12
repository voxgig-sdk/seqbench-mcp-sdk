import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { Translate, TranslateCreateData } from '../SeqbenchMcpTypes';
declare class TranslateEntity extends SeqbenchMcpEntityBase<Translate> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: TranslateEntity): TranslateEntity;
    create(this: any, reqdata?: TranslateCreateData, ctrl?: Control): Promise<TranslateEntity>;
}
export { TranslateEntity };
