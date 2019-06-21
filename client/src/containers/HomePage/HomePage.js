import React, { Component } from 'react';
import { connect } from 'react-redux';

import * as storyActions from '../../actions/story.action';

import StoryFeed from '../../components/StoryFeed/StoryFeed';
import StoryModal from '../../components/StoryModal/StoryModal';

import './HomePage.css';


class HomePage extends Component {
    componentDidMount() {
        const { getStories } = this.props;

        getStories();
    }


    render() {
        const { isLoading } = this.props;

        if (isLoading) {
            return null;
        }

        return (
            <div>
                <StoryFeed />
                <StoryModal />
            </div>
        );
    }
}

const mapStateToProps = ({ storyReducer: { isLoading } }) => ({ isLoading });

const mapDispatchToProps = (dispatch) => ({
    getStories: () => dispatch(storyActions.getStories())
});


export default connect(mapStateToProps, mapDispatchToProps)(HomePage);
