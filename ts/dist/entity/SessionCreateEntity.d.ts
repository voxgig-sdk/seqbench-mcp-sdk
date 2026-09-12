import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SessionCreate, SessionCreateCreateData } from '../SeqbenchMcpTypes';
declare class SessionCreateEntity extends SeqbenchMcpEntityBase<SessionCreate> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SessionCreateEntity): SessionCreateEntity;
    create(this: any, reqdata?: SessionCreateCreateData, ctrl?: Control): Promise<SessionCreateEntity>;
}
export { SessionCreateEntity };
