import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SessionRun, SessionRunCreateData } from '../SeqbenchMcpTypes';
declare class SessionRunEntity extends SeqbenchMcpEntityBase<SessionRun> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SessionRunEntity): SessionRunEntity;
    create(this: any, reqdata?: SessionRunCreateData, ctrl?: Control): Promise<SessionRunEntity>;
}
export { SessionRunEntity };
