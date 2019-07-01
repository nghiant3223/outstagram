import * as actionTypes from '../constants/actionTypes';
import * as storyService from "../services/story.service";
import StoryFeedManager from '../StoryFeedManager';

export const displayStoryBoardNode = (sbNode) =>
    ({ type: actionTypes.DISPLAY_STORY_BOARD_NODE, payload: sbNode });

export const displayFirstSBNode = () =>
    ({ type: actionTypes.DISPLAY_STORY_BOARD_NODE, payload: StoryFeedManager.getInstance().get1stSBNode() });