import random


class Node:
    def __init__(self, value, size):
        self.value = value
        self.next = [None for i in range(size)]


class Skiplist:
    _maxLevel = 31
    _pFactor = 0.25

    def __init__(self):
        self.head = Node(None, Skiplist._maxLevel)
        self.currentLevel = 1

    def erase(self, num: int) -> bool:
        res = False
        cur = self.head
        for i in reversed(range(self.currentLevel)):
            cur = self._findClosest(cur, i, num)
            if cur.next[i] is not None and cur.next[i].value == num:
                res = True
                cur.next[i] = cur.next[i].next[i]
                continue
        return res

    def search(self, target: int) -> bool:
        cur = self.head
        for i in reversed(range(self.currentLevel)):
            cur = self._findClosest(cur, i, target)
            if cur.next[i] is not None and cur.next[i].value == target:
                return True
        return False

    def add(self, num: int):
        level = self._randomLevel()
        newNode = Node(num, level)

        updateNode = self.head
        for i in reversed(range(self.currentLevel)):
            updateNode = self._findClosest(updateNode, i, num)
            if i < level:
                if not updateNode.next[i]:
                    updateNode.next[i] = newNode
                else:
                    tmp = updateNode.next[i]
                    updateNode.next[i] = newNode
                    newNode.next[i] = tmp

        if level > self.currentLevel:
            for i in range(self.currentLevel, level):
                self.head.next[i] = newNode
            self.currentLevel = level

    def _randomLevel(self) -> int:
        level = 1
        while random.random() < Skiplist._pFactor and level < Skiplist._maxLevel:
            level += 1
        return level

    def _findClosest(self, node: Node, levelIndex: int, value: int) -> Node:
        while node.next[levelIndex] is not None and value > node.next[levelIndex].value:
            node = node.next[levelIndex]
        return node
