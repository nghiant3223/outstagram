import React, { Component } from 'react';
import { connect } from 'react-redux';

import socket from '../../Socket';
import * as storyActions from '../../actions/story.action';

import StoryFeed from '../../components/StoryFeed/StoryFeed';
import StoryModal from '../../components/StoryModal/StoryModal';
import CreatorModal from '../../components/CreatorModal/CreatorModal';

import './HomePage.css';
import StoryFeedManager from '../../StoryFeedManager';

class HomePage extends Component {
    componentDidMount() {
        const { getStories, user } = this.props;

        getStories();
        socket.open({ userID: user.id });

        socket.on("STORY.SERVER.POST_STORY", (message) => {
            const storyFeedManager = StoryFeedManager.getInstance();
            storyFeedManager.prependUserStory(message.actorID, ...message.data)
                .then(this.updateStoryFeed)
                .catch((e) => console.log(e));
        });
    }

    updateStoryFeed = () => {
        this.storyFeed.updateStoryFeed();
    }

    render() {
        return (
            <div>
                <StoryFeed ref={(cmp) => { if (cmp) { this.storyFeed = cmp } }} />
                <StoryModal updateStoryFeed={this.updateStoryFeed} />
                <CreatorModal updateStoryFeed={this.updateStoryFeed} />
            </div>
        );
    }
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

const mapDispatchToProps = (dispatch) => ({
    getStories: () => dispatch(storyActions.getStories())
});

export default connect(mapStateToProps, mapDispatchToProps)(HomePage);