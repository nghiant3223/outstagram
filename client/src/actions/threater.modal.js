import * as actionTypes from '../constants/actionTypes';

export const openModal = (post) =>
    ({ type: actionTypes.OPEN_THEATER_MODAL, payload: post });

export const closeModal = () =>
    ({ type: actionTypes.CLOSE_THEATER_MODAL });