import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SavePermalink, SavePermalinkCreateData } from '../SeqbenchMcpTypes';
declare class SavePermalinkEntity extends SeqbenchMcpEntityBase<SavePermalink> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SavePermalinkEntity): SavePermalinkEntity;
    create(this: any, reqdata?: SavePermalinkCreateData, ctrl?: Control): Promise<SavePermalinkEntity>;
}
export { SavePermalinkEntity };
