import React, { Component } from 'react';
import { Redirect, Route } from 'react-router-dom';
import { connect } from 'react-redux';

import LoginForm from '../../components/LoginForm/LoginForm';
import RegisterForm from '../../components/RegisterForm/RegisterForm';

import './EntrancePage.css';

function entrancePage(props) {
    const { isAuthenticated } = props;

    if (isAuthenticated) {
        return <Redirect to="/" />
    }

    return (
        <div>
            <Route path="/login" component={() => <LoginForm />} />
            <Route path="/register" component={() => <RegisterForm />} />
        </div>
    );

}

const mapStateToProps = ({ authReducer: { isAuthenticated } }) => ({ isAuthenticated });

export default connect(mapStateToProps)(entrancePage);