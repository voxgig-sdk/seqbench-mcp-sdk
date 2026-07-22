
import { test, describe } from 'node:test'
import { equal } from 'node:assert'


import { SeqbenchMcpSDK } from '..'


describe('exists', async () => {

  test('test-mode', async () => {
    const testsdk = await SeqbenchMcpSDK.test()
    equal(null !== testsdk, true)
  })

})
