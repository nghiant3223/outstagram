import DoublyLinkedList from './ds/DoubleLinkedList';

class StoryFeedManager {
    constructor(storyBoards) {
        this.ll = new DoublyLinkedList();
        this.activeLL = new DoublyLinkedList();
        this.inactiveLL = new DoublyLinkedList();

        this.ll.append(storyBoards[0]);

        let i = 1;
        while (i < storyBoards.length && storyBoards[i].hasNewStory) {
            this.activeLL.append(storyBoards[i++]);
        }

        while (i < storyBoards.length) {
            this.inactiveStoryBoards.append(storyBoards[i++]);
        }

        if (!this.activeLL.isEmpty()) {
            const activeLLHead = this.activeLL.getHead();

            this.ll.head.setNext(activeLLHead);
            activeLLHead.setPrevious(this.ll.head);

            if (!this.inactiveLL.isEmpty()) {
                const activeLLTail = this.activeLL.getTail();
                const inactiveLLHead = this.inactiveLL.getHead();

                activeLLTail.setNext(inactiveLLHead);
                inactiveLLHead.setPrevious(activeLLTail);
            }
        } else if (!this.inactiveLL.isEmpty()) {
            const inactiveLLHead = this.inactiveLL.getHead();
            this.ll.head.setNext(inactiveLLHead);
            inactiveLLHead.setPrevious(this.ll.head);
        }
    }

    getFirstSBNode() {
        return this.ll.getHead();
    }

    map(hof) {
        let sbNode = this.ll.getHead();
        let results = [];

        while (sbNode != null) {
            results.push(hof(sbNode));
            sbNode = sbNode.getNext();
        }

        return results;
    }
}

export default (function () {
    let storyFeedManager;

    function initStoryFeedManager(boards) {
        if (storyFeedManager === undefined) {
            storyFeedManager = new StoryFeedManager(boards);
        } else {
            throw "StoryFeedManager has already existed";
        }
    }

    function getInstance() {
        if (storyFeedManager === undefined) {
            throw "StoryFeedManager hasn't been created yet";
        } else {
            return storyFeedManager;
        }
    }

    return {
        getInstance,
        initStoryFeedManager
    }
})();