import { Context } from './Context';
declare class SeqbenchMcpError extends Error {
    isSeqbenchMcpError: boolean;
    sdk: string;
    code: string;
    ctx: Context;
    status: number;
    get notFound(): boolean;
    constructor(code: string, msg: string, ctx: Context);
}
export { SeqbenchMcpError };
