import * as actionTypes from '../constants/actionTypes';

export const openModal = () =>
    ({ type: actionTypes.OPEN_THEATER_MODAL });

export const closeModal = () =>
    ({ type: actionTypes.CLOSE_THEATER_MODAL });