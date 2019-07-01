import React, { Component } from 'react';
import { connect } from 'react-redux';

import socket from '../../Socket';
import * as storyActions from '../../actions/story.action';
import * as storyServices from '../../services/story.service';

import StoryFeed from '../../components/StoryFeed/StoryFeed';
import StoryModal from '../../components/StoryModal/StoryModal';
import CreatorModal from '../../components/CreatorModal/CreatorModal';
import Container from "../../components/Container/Container";

import './HomePage.css';

import * as uiActions from '../../actions/ui.action';
import StoryFeedManager from '../../StoryFeedManager';

class HomePage extends Component {
    state = {
        isLoading: true
    }

    async componentDidMount() {
        try {
            const { data: { data: { storyBoards } } } = await storyServices.getStoryFeed();
            StoryFeedManager.initialize(storyBoards);

            socket.on("STORY.SERVER.POST_STORY", (message) => {
                const storyFeedManager = StoryFeedManager.getInstance();
                storyFeedManager.prependUserStory(message.actorID, ...message.data)
                    .then(this.updateStoryFeed)
                    .catch((e) => console.log(e));
            });
        } catch (e) {
            console.log("Error while fetching story: ", e);
        } finally {
            this.setState({ isLoading: false });
        }
    }

    componentWillUnmount() {
        StoryFeedManager.removeInstance();
        this.props.closeModal();

    }

    updateStoryFeed = () => {
        this.storyFeed.update();
    }

    render() {
        const { isLoading } = this.state;

        if (isLoading) {
            return null;
        }

        return (
            <Container>
                <StoryFeed ref={(cmp) => { if (cmp) { this.storyFeed = cmp } }} fetchingStoryFeed={this.state.isLoading} />
                <StoryModal updateStoryFeed={this.updateStoryFeed} />
                <CreatorModal updateStoryFeed={this.updateStoryFeed} />
            </Container>
        );
    }
}

const mapDispatchToProps = (dispatch) => ({
    closeModal: () => dispatch(uiActions.closeStoryModal())
});

export default connect(null, mapDispatchToProps)(HomePage);