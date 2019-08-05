import * as actionTypes from '../constants/actionTypes';

export const openModal = (payload) =>
    ({ type: actionTypes.OPEN_CREATOR_MODAL, payload });

export const closeModal = () =>
    ({ type: actionTypes.CLOSE_CREATOR_MODAL });