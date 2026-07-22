
import { Context } from './Context'


class SeqbenchMcpError extends Error {

  isSeqbenchMcpError = true

  sdk = 'SeqbenchMcp'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  SeqbenchMcpError
}

