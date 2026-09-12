import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ParseGenbank, ParseGenbankCreateData } from '../SeqbenchMcpTypes';
declare class ParseGenbankEntity extends SeqbenchMcpEntityBase<ParseGenbank> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ParseGenbankEntity): ParseGenbankEntity;
    create(this: any, reqdata?: ParseGenbankCreateData, ctrl?: Control): Promise<ParseGenbankEntity>;
}
export { ParseGenbankEntity };
