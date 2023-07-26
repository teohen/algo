const test = require('node:test')
const assert = require('node:assert')

const quicksort = require('./index.js')

test('Should sort the list', () => {
  const list = [4, 2, 12, 19, 3, 8, 13, 30, 28, 18, 21, 26]

  const sortedList = quicksort(list, 0, list.length - 1)
  const expected = [4, 2, 12, 19, 3, 8, 13, 30, 28, 18, 21, 26].sort((a, b) => {
    if (a > b)
      return 1
    else if (a < b)
      return -1
    else 
      return 0
  })

  assert.deepEqual(sortedList, expected)
})


