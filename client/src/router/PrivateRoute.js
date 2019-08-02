import React from 'react';
import { Route } from 'react-router-dom';
import { Redirect } from 'react-router-dom';

const PrivateRoute = ({ component: Component, isAuthenticated, ...rest }) =>
    <Route {...rest} render={props => isAuthenticated ? <Component {...props} /> : <Redirect to='/login' />} />;

export default PrivateRoute;