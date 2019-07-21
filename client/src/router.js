import React from 'react';
import { Redirect } from 'react-router-dom';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import MainLayout from './layouts/MainLayout/MainLayout';
import HomePage from './containers/HomePage/HomePage';
import NotFoundPage from './containers/NotFoundPage/NotFoundPage';
import EntrancePage from './containers/EntrancePage/EntrancePage';
import ProfilePage from './containers/ProfilePage/ProfilePage';
import MessagePage from './containers/MessagePage/MessagePage';

const PrivateRoute = ({ component: Component, isAuthenticated, ...rest }) =>
    <Route {...rest} render={props => isAuthenticated ? <Component {...props} /> : <Redirect to='/login' />} />;

export default function router({ isAuthenticated }) {
    return (
        <BrowserRouter>
            <Switch>
                <Route path='/(login|register)' exact component={EntrancePage} isAuthenticated={isAuthenticated} />
                <MainLayout>
                    <Switch>
                        <PrivateRoute path='/' exact component={HomePage} isAuthenticated={isAuthenticated} />
                        <PrivateRoute path='/:username' exact component={ProfilePage} isAuthenticated={isAuthenticated} />
                        <PrivateRoute path='/messages/:username' exact component={MessagePage} isAuthenticated={isAuthenticated} />
                        <Route component={NotFoundPage} />
                    </Switch>
                </MainLayout>
            </Switch>
        </BrowserRouter>
    );
}