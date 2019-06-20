import React, { Component } from 'react';
import { connect } from 'react-redux';

import './HomePage.css';
import StoryFeed from '../../components/StoryFeed/StoryFeed';

class HomePage extends Component {
    render() {
        return (
            <div>
                <StoryFeed />
            </div>
        );
    }
}

const mapStateToProps = ({ auth: { user } }) => ({ user });

export default connect(mapStateToProps)(HomePage);
