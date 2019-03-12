// Linked lists

let linkedList = []

// Declaring the nodes
linkedList[0] = {
  value: 0,
  random: null,
}
linkedList[1] = {
  value: 1,
  next: null,
  random: linkedList[0]
}
linkedList[2] = {
  value: 2,
  random: null,
  next: null,
}
linkedList[0].next = linkedList[1]
linkedList[1].next = linkedList[2]
linkedList[0].random = linkedList[1]

let cloneLinkedList = (node, previousNodes = []) => {
  let { next, random, value } = node

  let found = previousNodes.find(({ original, copy }) => original === node)
  if (found) return found.copy

  let newNode = { value }
  previousNodes.push({ original: node, copy: newNode })

  newNode.next = next ? cloneLinkedList(next, previousNodes) : null
  newNode.random = random ? cloneLinkedList(random, previousNodes) : null

  return newNode
}

let simpleNode = {
  next: null,
}

simpleNode.random = simpleNode
console.log(cloneLinkedList(simpleNode))

console.log(cloneLinkedList(linkedList[0]))

// It's not much, but it's honest work

// BEWARE MAPS

let cloneLinkedListWithMap = (node, previousNodes = new Map()) => {
  let { next, random, value } = node

  if (previousNodes.has(node)) return previousNodes.get(node)

  let newNode = { value }
  previousNodes.set(node, newNode)

  newNode.next = next ? cloneLinkedList(next, previousNodes) : null
  newNode.random = random ? cloneLinkedList(random, previousNodes) : null

  return newNode
}

console.log('LAST ONE', cloneLinkedList(simpleNode))
console.log('ACTUAL LAST ONE', cloneLinkedList(linkedList[0]))
