function bubbleSort(list, start, end) {
  if (start === end)
    return list

  for (let i = 0; i < list.length; i++) {

    if (list[i] > list[i + 1]) {
      const tmp = list[i]
      list[i] = list[i + 1]
      list[i + 1] = tmp
    }
  }
  end -= 1
  return bubbleSort(list, start, end)
}
/*const list = [25,
3,
24,
5,
78,
12,
29,
84,
6,
14,
2,
21,
48]
console.log(bubbleSort(list, 0, list.length - 1)) */

module.exports = bubbleSort
