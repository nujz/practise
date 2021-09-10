class Node:
    def __init__(self, key=0, value=0):
        self.key = key
        self.value = value
        self.prev = None
        self.next = None


class LRUCache:
    def __init__(self, capacity: int):
        self.cache = dict()
        self.size = 0
        self.capacity = capacity

        self.head = Node()
        self.tail = Node()
        self.head.next = self.tail
        self.tail.prev = self.head

    def get(self, key: int) -> int:
        if key not in self.cache:
            return -1
        node = self.cache[key]
        self.moveToHead(node)
        return node.value

    def put(self, key: int, value: int):
        if key not in self.cache:
            node = Node(key, value)
            self.cache[key] = node
            self.addToHead(node)
            if self.size == self.capacity:
                removed = self.removeTail()
                self.cache.pop(removed.key)
            else:
                self.size += 1
        else:
            node = self.cache[key]
            node.value = value
            self.moveToHead(node)

    def addToHead(self, node: Node):
        node.prev = self.head
        node.next = self.head.next
        self.head.next.prev = node
        self.head.next = node

    def removeNode(self, node: Node):
        node.prev.next = node.next
        node.next.prev = node.prev
        node.next = None
        node.prev = None

    def moveToHead(self, node):
        self.removeNode(node)
        self.addToHead(node)

    def removeTail(self) -> Node:
        node = self.tail.prev
        self.removeNode(node)
        return node
