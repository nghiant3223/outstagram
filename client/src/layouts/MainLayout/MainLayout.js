import React from 'react';

import Header from '../../components/Header/Header';
import Footer from '../../components/Footer/Footer';

import './MainLayout.css';

const MainLayout = (props) => (
    <div>
        <Header />
        
        <div className="MainLayout">
            <aside className="LeftAside">
            </aside>

            <main>
                {props.children}
            </main>

            <aside className="RightAside">
            </aside>
        </div>

        <Footer />
    </div>
);

export default MainLayout;