import DoublyLinkedList from './ds/DoubleLinkedList';

class StoryFeedManager {
    constructor(storyBoards) {
        this.ll = new DoublyLinkedList();

        let i = 1;
        this.ll.append(storyBoards[0]);
        while (i < storyBoards.length) {
            this.ll.append(storyBoards[i++]);
        }
    }

    getFirstSBNode() {
        const head = this.ll.getHead();

        if (head.stories === null) {
            return head.getNext();
        }

        return head;
    }

    prependUserStory(story) {
        let headStories = this.ll.getHead().getValue().stories;

        if (headStories === null) {
            headStories = [];
        }

        headStories.unshift(story);
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

    setInactive(sbNode) {
        if (sbNode == this.ll.getHead()) {
            return;
        }

        const storyBoard = sbNode.getValue();
        if (storyBoard.stories.every((story) => story.seen)) {
            storyBoard.hasNewStory = false;
            this.ll.delete(storyBoard);
            this.ll.append(storyBoard);
        }
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