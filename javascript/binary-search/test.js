const test = require('node:test')
const assert = require('node:assert')
const binarySearch = require('./index')

const list = [1, 3, 6, 8, 10, 14, 16, 19, 22]

test('Should find the item in the list', () => {
  const result = binarySearch(list, 19, 0, list.length - 1)
  assert.equal(list[result], 19)
})

test('Should return undefined if the item is not on the list', () => {

  const result = binarySearch(list, 11, 0, list.length - 1)
  assert.equal(result, undefined)
})

