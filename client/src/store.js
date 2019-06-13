import thunk from 'redux-thunk';
import logger from 'redux-logger';
import { combineReducers, compose, createStore, applyMiddleware } from 'redux';

import reducers from './reducers';

export default function () {
    const rootReducer = combineReducers(reducers);
    const composeEnhancers = window.__REDUX_DEVTOOLS_EXTENSION_COMPOSE__ || compose;
    return createStore(rootReducer, composeEnhancers(applyMiddleware(thunk, logger)));
}
