import * as actionTypes from '../constants/actionTypes';
import * as storyService from "../services/story.service";
import StoryFeedManager from '../StoryFeedManager';

export const getStories = () =>
    async (dispatch) => {
        try {
            const { data: { data: { storyBoards } } } = await storyService.getStoryFeed();

            StoryFeedManager.initStoryFeedManager(storyBoards);
            dispatch({ type: actionTypes.GET_STORY_FEED });
        } catch (e) {
            console.log(e);
        }
    }

export const displayStoryBoardNode = (sbNode) =>
    ({ type: actionTypes.DISPLAY_STORY_BOARD_NODE, payload: sbNode });

export const displayFirstSBNode = () =>
    ({ type: actionTypes.DISPLAY_STORY_BOARD_NODE, payload: StoryFeedManager.getInstance().getFirstSBNode() });