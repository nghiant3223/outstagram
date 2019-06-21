import * as actionTypes from '../constants/actionTypes';

export function openStoryModal(storyBoardNode) {
    return { type: actionTypes.OPEN_STORY_MODAL, payload: storyBoardNode }
}

export function closeStoryModal() {
    return { type: actionTypes.CLOSE_STORY_MODAL }
}