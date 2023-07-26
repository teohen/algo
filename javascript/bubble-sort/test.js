const test = require('node:test')
const assert = require('node:assert/strict')

const bubbleSort = require('./index.js')
const list = [4, 25, 3, 24, 5, 78, 12, 29, 84, 6, 14, 2, 21, 48]

test('Should order the list', () => {
  const result = bubbleSort(list, 0, list.length - 1)
  assert.deepEqual(result, list.sort((a, b) => {
    if (a > b)
      return 1
    else if (a < b)
      return -1
    else return 0
  }))
})


