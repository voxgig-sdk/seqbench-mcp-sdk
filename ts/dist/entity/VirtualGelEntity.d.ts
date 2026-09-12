import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { VirtualGel, VirtualGelCreateData } from '../SeqbenchMcpTypes';
declare class VirtualGelEntity extends SeqbenchMcpEntityBase<VirtualGel> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: VirtualGelEntity): VirtualGelEntity;
    create(this: any, reqdata?: VirtualGelCreateData, ctrl?: Control): Promise<VirtualGelEntity>;
}
export { VirtualGelEntity };
