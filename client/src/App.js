import React, { Component } from 'react';
import { connect } from 'react-redux';

import Router from './router';
import * as authActions from './actions/auth.action';
import './App.css';

class App extends Component {
    componentDidMount = () => {
        const { getMe } = this.props;
        getMe();
    }

    render() {
        const { isLoading, isAuthenticated } = this.props;

        if (isLoading) {
            return <div>Loading...</div>
        }

        return <Router isAuthenticated={isAuthenticated} />;
    }
}

const mapStateToProps = ({ auth: { isAuthenticated }, ui: { isLoading } }) => ({ isAuthenticated, isLoading });

const mapDispatchToProps = (dispatch) => ({
    getMe: () => dispatch(authActions.getMe())
});

export default connect(mapStateToProps, mapDispatchToProps)(App);
