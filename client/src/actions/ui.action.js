import * as actionTypes from '../constants/actionTypes';

export function openStoryModal() {
    return { type: actionTypes.OPEN_STORY_MODAL }
}

export function closeStoryModal() {
    return { type: actionTypes.CLOSE_STORY_MODAL }
}