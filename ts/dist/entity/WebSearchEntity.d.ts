import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { WebSearch, WebSearchCreateData } from '../SeqbenchMcpTypes';
declare class WebSearchEntity extends SeqbenchMcpEntityBase<WebSearch> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: WebSearchEntity): WebSearchEntity;
    create(this: any, reqdata?: WebSearchCreateData, ctrl?: Control): Promise<WebSearchEntity>;
}
export { WebSearchEntity };
