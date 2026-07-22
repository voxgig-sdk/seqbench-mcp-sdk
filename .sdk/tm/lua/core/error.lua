-- SeqbenchMcp SDK error

local SeqbenchMcpError = {}
SeqbenchMcpError.__index = SeqbenchMcpError


function SeqbenchMcpError.new(code, msg, ctx)
  local self = setmetatable({}, SeqbenchMcpError)
  self.is_sdk_error = true
  self.sdk = "SeqbenchMcp"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function SeqbenchMcpError:error()
  return self.msg
end


function SeqbenchMcpError:__tostring()
  return self.msg
end


return SeqbenchMcpError
