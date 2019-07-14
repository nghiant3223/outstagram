import React from 'react';

import Header from '../../components/Header/Header';
import Footer from '../../components/Footer/Footer';

import './MainLayout.css';
import ReactorModal from '../../components/ReactorModal/ReactorModal';
import ThreaterModal from '../../components/ThreaterModal/ThreaterModal';

const MainLayout = (props) => (
    <div>
        <Header />
        {props.children}
        <Footer />

        <ThreaterModal />
        <ReactorModal />
    </div>
);

export default MainLayout;