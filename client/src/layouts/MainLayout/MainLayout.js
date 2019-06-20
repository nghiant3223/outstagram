import React from 'react';
import { Grid, Image } from 'semantic-ui-react';

import Header from '../../components/Header/Header';
import Footer from '../../components/Footer/Footer';

import './MainLayout.css';
import StoryModal from '../../components/StoryModal/StoryModal';

const MainLayout = (props) => (
    <div>
        <Header />

        <Grid divided='vertically'>
            <Grid.Row>
                <Grid.Column width={3}>
                </Grid.Column>

                <Grid.Column width={10}>
                    {props.children}
                </Grid.Column>

                <Grid.Column width={3}>
                </Grid.Column>
            </Grid.Row>
        </Grid>

        <Footer />

        <StoryModal />
    </div>
);

export default MainLayout;