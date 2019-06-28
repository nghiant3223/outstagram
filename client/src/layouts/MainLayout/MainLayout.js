import React from 'react';

import Header from '../../components/Header/Header';
import Footer from '../../components/Footer/Footer';

import './MainLayout.css';

const MainLayout = (props) => (
    <div>
        <Header />

        {props.children}

        <Footer />

    </div>
);

export default MainLayout;