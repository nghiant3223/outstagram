import * as actionTypes from '../constants/actionTypes';

export const openModal = (postImage) =>
    ({ type: actionTypes.OPEN_THEATER_MODAL, payload: postImage });

export const closeModal = () =>
    ({ type: actionTypes.CLOSE_THEATER_MODAL });