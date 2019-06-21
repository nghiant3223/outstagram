import React, { Component } from 'react';
import { connect } from 'react-redux';

import Router from './router';
import * as authActions from './actions/auth.action';
import './App.css';
import { getStories } from './actions/story.action';

class App extends Component {
    componentDidMount = () => {
        const { getMe, getStories } = this.props;
        getMe();
        getStories();
    }

    render() {
        const { isLoading, isAuthenticated } = this.props;

        if (isLoading) {
            return <div>Loading...</div>
        }

        return <Router isAuthenticated={isAuthenticated} />;
    }
}

const mapStateToProps = ({ authReducer: { isAuthenticated }, uiReducer: { isLoading } }) => ({ isAuthenticated, isLoading });

const mapDispatchToProps = (dispatch) => ({
    getMe: () => dispatch(authActions.getMe()),
    getStories: () => dispatch(getStories())
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
