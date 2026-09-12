import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { MeltingTemperature, MeltingTemperatureCreateData } from '../SeqbenchMcpTypes';
declare class MeltingTemperatureEntity extends SeqbenchMcpEntityBase<MeltingTemperature> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: MeltingTemperatureEntity): MeltingTemperatureEntity;
    create(this: any, reqdata?: MeltingTemperatureCreateData, ctrl?: Control): Promise<MeltingTemperatureEntity>;
}
export { MeltingTemperatureEntity };
