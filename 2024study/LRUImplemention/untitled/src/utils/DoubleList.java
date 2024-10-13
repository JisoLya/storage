package utils;

/**
 *  实现双向链表，最近使用的节点,移动到双向链表头部
 */
public class DoubleList {
    private Node head,tail;
    private int capacity;
    private int size = 0;
    public DoubleList(){
        this.head = new Node(0,0);
        this.tail = new Node(0,0);
        //default cap
        this.capacity = 5;
        head.next = tail;
        tail.pre = head;
    }
    public DoubleList(int capacity){
        this.head = new Node(0,0);
        this.tail = new Node(0,0);
        //custom cap
        this.capacity = capacity;
        head.next = tail;
        tail.pre = head;
    }

    public Node getHead() {
        return head;
    }

    public void setHead(Node head) {
        this.head = head;
    }

    public Node getTail() {
        return tail;
    }

    public void setTail(Node tail) {
        this.tail = tail;
    }

    public int getCapacity() {
        return capacity;
    }

    public void setCapacity(int capacity) {
        this.capacity = capacity;
    }

    public int getSize() {
        return size;
    }

    public void setSize(int size) {
        this.size = size;
    }

    public Node addToHead(Node node){
        this.size++;
        node.next = head.next;
        head.next.pre = node;
        head.next = node;
        node.pre = head;
        if (this.size > this.capacity){
            this.size--;
            return removeTail();
        }
        return null;
    }

    public void removeNode(Node node){
        node.next.pre = node.pre;
        node.pre.next = node.next;
        this.size--;
    }

    public void moveToHead(Node node){
        removeNode(node);
        addToHead(node);
    }

    public Node removeTail(){
        Node ret = this.tail.pre;
        ret.pre.next = this.tail;
        this.tail.pre = ret.pre;
        return ret;
    }

}
