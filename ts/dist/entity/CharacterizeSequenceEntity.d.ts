import { SeqbenchMcpEntityBase } from '../SeqbenchMcpEntityBase';
import type { SeqbenchMcpSDK } from '../SeqbenchMcpSDK';
import type { Control } from '../types';
import type { CharacterizeSequence, CharacterizeSequenceCreateData } from '../SeqbenchMcpTypes';
declare class CharacterizeSequenceEntity extends SeqbenchMcpEntityBase<CharacterizeSequence> {
    constructor(client: SeqbenchMcpSDK, entopts: any);
    make(this: CharacterizeSequenceEntity): CharacterizeSequenceEntity;
    create(this: any, reqdata?: CharacterizeSequenceCreateData, ctrl?: Control): Promise<CharacterizeSequenceEntity>;
}
export { CharacterizeSequenceEntity };
