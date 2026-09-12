import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { SessionGet, SessionGetCreateData } from '../SeqbenchMcpTypes';
declare class SessionGetEntity extends SeqbenchMcpEntityBase<SessionGet> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: SessionGetEntity): SessionGetEntity;
    create(this: any, reqdata?: SessionGetCreateData, ctrl?: Control): Promise<SessionGetEntity>;
}
export { SessionGetEntity };
