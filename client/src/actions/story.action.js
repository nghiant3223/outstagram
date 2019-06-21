import * as actionTypes from '../constants/actionTypes';

import * as storyService from "../services/story.service";

export const getStories = () =>
    async (dispatch) => {
        try {
            const { data: { data: { storyBoards } } } = await storyService.getStoryFeed();
            // TODO: Get user's own story
            if (storyBoards !== null) {
                dispatch({ type: actionTypes.GET_STORY_FEED, payload: storyBoards });
            } else {
                dispatch({ type: actionTypes.GET_STORY_FEED, payload: [] });
            }
        } catch (e) {
            console.log(e);
        }
    }

export const setOnDisplayStoryBoardNode = (storyBoardNode) => ({ type: actionTypes.SET_ON_DISPLAY_STORY_BOARD_NODE, payload: storyBoardNode });