function partition(list, start, end) {
  const pivot = list[end]

  let idx = start - 1

  for (let i = start; i < end; ++i) {
    if (list[i] <= pivot) {
      idx += 1
      const tmp = list[i]
      list[i] = list[idx]
      list[idx] = tmp
    }
  }

  idx += 1
  list[end] = list[idx]
  list[idx] = pivot

  return idx
}

function qs (list, start, end) {
  if (start >= end) {
    return
  }
  const pivotIndex = partition(list, start, end)
  quicksort(list, start, pivotIndex - 1)
  quicksort(list, pivotIndex + 1, end)
}

function quicksort(list, start, end) {
  qs(list, start, end)

  return list
}


module.exports  = quicksort
