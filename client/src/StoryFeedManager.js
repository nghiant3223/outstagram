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

        console.log(this.ll);
    }

    getFirstSBNode() {
        return this.ll.getHead();
    }
}

export default StoryFeedManager;