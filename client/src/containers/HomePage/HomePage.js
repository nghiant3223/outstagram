import React, { Component } from 'react';
import { connect } from 'react-redux';

import socket from '../../Socket';
import * as storyServices from '../../services/story.service';

import StoryFeed from '../../components/StoryFeed/StoryFeed';
import StoryModal from '../../components/StoryModal/StoryModal';
import CreatorModal from '../../components/CreatorModal/CreatorModal';
import Container from "../../components/Container/Container";

import './HomePage.css';

import * as uiActions from '../../actions/ui.action';
import StoryFeedManager from '../../StoryFeedManager';
import ThreaterModal from '../../components/ThreaterModal/ThreaterModal';

class HomePage extends Component {
    state = {
        isLoading: true
    }

    async componentDidMount() {
        try {
            const { data: { data: { storyBoards } } } = await storyServices.getStoryFeed();
            StoryFeedManager.initialize(storyBoards);

            socket.on("STORY.SERVER.POST_STORY", (message) => {
                const manager = StoryFeedManager.getInstance();
                manager.prependUserStory(message.actorID, ...message.data)
                    .then(this.updateStoryFeed)
                    .catch((e) => console.log(e));
            });

            socket.on("STORY.SERVER.REACT_STORY", (message) => {
                const manager = StoryFeedManager.getInstance();
                const { storyID, reactor } = message.data;
                const story = manager.getStory(storyID)

                // IMPORTANT: This should be !=, not !== because built-in `find` function return undefined when item not found in array
                if (story != null) {
                    if (story.reactors == null) {
                        story.reactors = [];
                    }
                    story.reactors.unshift(reactor);
                }
            });

            socket.on("STORY.SERVER.UNREACT_STORY", (message) => {
                const manager = StoryFeedManager.getInstance();
                const { storyID, reactor } = message.data;
                const story = manager.getStory(storyID);

                // IMPORTANT: This should be !=, not !== because built-in `find` function return undefined when item not found in array
                if (story != null) {
                    story.reactors = story.reactors.filter((_reactor) => _reactor.id !== reactor.id);
                }
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
                <ThreaterModal />
            </Container>
        );
    }
}

const mapDispatchToProps = (dispatch) => ({
    closeModal: () => dispatch(uiActions.closeStoryModal())
});

export default connect(null, mapDispatchToProps)(HomePage);