import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SessionSet, SessionSetCreateData } from '../SeqbenchMcpTypes';
declare class SessionSetEntity extends SeqbenchMcpEntityBase<SessionSet> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SessionSetEntity): SessionSetEntity;
    create(this: any, reqdata?: SessionSetCreateData, ctrl?: Control): Promise<SessionSetEntity>;
}
export { SessionSetEntity };
