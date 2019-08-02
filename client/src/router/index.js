import React from 'react';
import { BrowserRouter, Route, Switch } from 'react-router-dom';

import MainLayout from '../layouts/MainLayout/MainLayout';
import HomePage from '../containers/HomePage/HomePage';
import NotFoundPage from '../containers/NotFoundPage/NotFoundPage';
import EntrancePage from '../containers/EntrancePage/EntrancePage';
import ProfilePage from '../containers/ProfilePage/ProfilePage';
import MessagePage from '../containers/MessagePage/MessagePage';
import PrivateRoute from './PrivateRoute';
import ScrollToTop from './ScrollToTop';

export default function router({ isAuthenticated }) {
    return (
        <BrowserRouter >
            <Switch>
                <Route path='/(login|register)' exact component={EntrancePage} isAuthenticated={isAuthenticated} />
                <MainLayout>
                    <ScrollToTop>
                        <Switch>
                            <PrivateRoute path='/' exact component={HomePage} isAuthenticated={isAuthenticated} />
                            <PrivateRoute path='/messages' exact component={MessagePage} isAuthenticated={isAuthenticated} />
                            <PrivateRoute path='/messages/:roomIdOrUsername' exact component={MessagePage} isAuthenticated={isAuthenticated} />
                            <PrivateRoute path='/:username' exact component={ProfilePage} isAuthenticated={isAuthenticated} />
                            <Route component={NotFoundPage} />
                        </Switch>
                    </ScrollToTop>
                </MainLayout>
            </Switch>
        </BrowserRouter>
    );
}