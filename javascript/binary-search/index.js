function search(list, needle, lo, hi) {
  const half = Math.floor((hi - lo) / 2 ) + lo


  const halfItem = list[half]

  if (halfItem === needle) {
     return half
  }

  if (lo === hi) {
    return undefined
  }

  if (needle < halfItem) {
    return search(list, needle, lo, half)
  } else if (needle >= halfItem) {
    return search(list, needle, half + 1, hi)
  }
}

module.exports = search
