import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { ListTool, ListToolLoadMatch } from '../SeqbenchMcpTypes';
declare class ListToolEntity extends SeqbenchMcpEntityBase<ListTool> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: ListToolEntity): ListToolEntity;
    load(this: any, reqmatch?: ListToolLoadMatch, ctrl?: Control): Promise<ListToolEntity>;
}
export { ListToolEntity };
