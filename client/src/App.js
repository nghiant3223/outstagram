import React, { Component } from 'react';
import { connect } from 'react-redux';

import * as authActions from './actions/auth.action';
import Router from './router/index';

import './App.css';
import Loading from './components/Loading/Loading';

class App extends Component {
    componentDidMount = () => {
        const { getMe } = this.props;
        getMe();
    }

    render() {
        const { isLoading, isAuthenticated } = this.props;

        if (isLoading) {
            return <div className="InitialLoaderContainer"><Loading /></div>
        }

        return <Router isAuthenticated={isAuthenticated} />;
    }
}

const mapStateToProps = ({ authReducer: { isAuthenticated }, uiReducer: { isLoading } }) => ({ isAuthenticated, isLoading });

const mapDispatchToProps = (dispatch) => ({
    getMe: () => dispatch(authActions.getMe())
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
