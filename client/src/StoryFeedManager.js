class StoryFeedManager {
    constructor(storyBoards) {
        this.activeStoryBoards = new DoublyLinkedList();
        this.inactiveStoryBoards = new DoublyLinkedList();

        let i = 0;
        while (i < storyBoards.length && storyBoards[i].hasNewStory) {
            this.activeStoryBoards.append(storyBoards[i++]);
        }

        while (i < storyBoards.length) {
            this.inactiveStoryBoards.append(storyBoards[i++]);
        }

        
    }


}