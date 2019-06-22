import * as actionTypes from '../constants/actionTypes';

export function openStoryModal(sbNode) {
    return { type: actionTypes.OPEN_STORY_MODAL, payload: sbNode }
}

export function closeStoryModal() {
    return { type: actionTypes.CLOSE_STORY_MODAL }
}