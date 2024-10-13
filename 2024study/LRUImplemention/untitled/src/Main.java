import utils.LRUCache;
import utils.Node;

public class Main {
    public static void main(String[] args) {
        LRUCache lruCache = new LRUCache(5);
        Node node1 = new Node(1,3);
        Node node2 = new Node(3,2);
        Node node3 = new Node(5,4);
        Node node4 = new Node(7,6);
        Node node5 = new Node(9,8);
        Node node6 = new Node(10,11);
        lruCache.put(node1);
        lruCache.put(node2);
        lruCache.put(node3);
        lruCache.put(node4);
        lruCache.put(node5);
        lruCache.showCache();
        System.out.println("===================");
        lruCache.get(3);
        lruCache.showCache();
        lruCache.put(node6);
        System.out.println("===================");
        lruCache.showCache();


    }
}