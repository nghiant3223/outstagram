import * as actionTypes from '../constants/actionTypes';

export const openModal = (reactableID) =>
    ({ type: actionTypes.OPEN_REACTOR_MODAL, payload: reactableID });

export const closeModal = () =>
    ({ type: actionTypes.CLOSE_REACTOR_MODAL });