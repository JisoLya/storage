package utils;

import java.util.HashMap;

/**
 * @author Soyan
 * @date 2024.8.29
 * LRU算法实现
 */

public class LRUCache {
    //缓存容量

    public DoubleList doubleList;

    public HashMap<Integer, Node> map;

    public LRUCache() {
        doubleList = new DoubleList(5);
        map = new HashMap<>();
    }

    public LRUCache(int capacity) {
        doubleList = new DoubleList(capacity);
        map = new HashMap<>();
    }

    public void put(Node node) {
        Node del = doubleList.addToHead(node);
        if (del != null){
            System.out.println("节点del被清除 key = " + del.key + ", val = " + del.val);
            map.remove(del.key);
        }
        map.put(node.key, node);

    }

    public int get(int key) {
        if (!map.containsKey(key))
            return -1;
        Node node = map.get(key);
        doubleList.moveToHead(node);
        return node.val;
    }

    public void showCache() {
        Node node = doubleList.getHead().next;
        while (true) {
            if (node == doubleList.getTail())
                return;
            System.out.println("节点key = " + node.key + "节点val = " + node.val);
            node = node.next;
        }
    }
}
