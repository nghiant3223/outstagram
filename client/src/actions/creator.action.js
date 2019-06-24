import * as actionTypes from '../constants/actionTypes';

export const openCreatorModal = () =>
    ({ type: actionTypes.OPEN_CREATOR_MODAL });

export const closeCreatorModal = () =>
    ({ type: actionTypes.CLOSE_CREATOR_MODAL });