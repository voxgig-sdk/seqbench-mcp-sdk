import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ReverseTranslate, ReverseTranslateCreateData } from '../SeqbenchMcpTypes';
declare class ReverseTranslateEntity extends SeqbenchMcpEntityBase<ReverseTranslate> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ReverseTranslateEntity): ReverseTranslateEntity;
    create(this: any, reqdata?: ReverseTranslateCreateData, ctrl?: Control): Promise<ReverseTranslateEntity>;
}
export { ReverseTranslateEntity };
