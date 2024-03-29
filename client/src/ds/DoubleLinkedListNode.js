export default class DoublyLinkedListNode {
    constructor(value, next = null, previous = null) {
        this.value = value;
        this.next = next;
        this.previous = previous;
    }

    toString(callback) {
        return callback ? callback(this.value) : `${this.value}`;
    }

    getValue() {
        return this.value;
    }

    getNext() {
        return this.next;
    }

    setNext(node) {
        this.next = node;
    }

    getPrevious() {
        return this.previous;
    }

    setPrevious(node) {
        this.previous = node;
    }
}