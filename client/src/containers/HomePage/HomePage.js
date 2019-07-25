import React, { Component } from 'react';
import { connect } from 'react-redux';

import Socket from '../../Socket';
import * as storyServices from '../../services/story.service';
import * as userServices from '../../services/user.service';

import StoryFeed from '../../components/StoryFeed/StoryFeed';
import StoryModal from '../../components/StoryModal/StoryModal';
import CreatorModal from '../../components/CreatorModal/CreatorModal';
import Container from "../../components/Container/Container";
import Post from "../../components/Post/Post";

import './HomePage.css';

import * as uiActions from '../../actions/ui.action';
import StoryFeedManager from '../../StoryFeedManager';
import PostPlaceholder from '../../components/Post/PostPlaceholder';

class HomePage extends Component {
    state = {
        posts: [],
        sinceID: 0,
        isLoading: true,
        isFetchingMorePost: false,
    }

    shouldScroll = true

    async componentDidMount() {
        document.addEventListener('scroll', this.trackScrolling);

        try {
            const { data: { data: { storyBoards } } } = await storyServices.getStoryFeed();
            const { data: { data: { posts, nextSinceID } } } = await userServices.getNewsFeed(0);

            this.setState({ posts: posts || [], sinceID: nextSinceID });
            StoryFeedManager.initialize(storyBoards);

            Socket.on("STORY.SERVER.POST_STORY", (message) => {
                const manager = StoryFeedManager.getInstance();
                manager.prependUserStory(message.actorID, ...message.data)
                    .then(this.updateStoryFeed)
                    .catch((e) => console.log(e));
            });

            Socket.on("STORY.SERVER.REACT_STORY", (message) => {
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

            Socket.on("STORY.SERVER.UNREACT_STORY", (message) => {
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
        document.removeEventListener('scroll', this.trackScrolling);
    }

    updateStoryFeed = () => {
        this.storyFeed.update();
    }

    isBottom(el) {
        return el.getBoundingClientRect().bottom <= window.innerHeight;
    }

    trackScrolling = async () => {
        const { sinceID } = this.state;
        const wrappedElement = document.getElementsByTagName("body")[0];

        if (this.isBottom(wrappedElement) && this.shouldScroll && sinceID) {
            this.shouldScroll = false;
            this.setState({ isFetchingMorePost: true });
            userServices.getNewsFeed(sinceID)
                .then(({ data: { data: { posts, nextSinceID } } }) => {
                    this.setState((prevState) => ({ posts: [...prevState.posts, ...posts], sinceID: nextSinceID }));
                    this.shouldScroll = true;
                })
                .catch((e) => console.log("Error while fetching more post"))
                .finally(() => this.setState({ isFetchingMorePost: false }));
        }
    };

    render() {
        const { isLoading, isFetchingMorePost, posts } = this.state;

        return (
            <Container white={false}>
                <StoryFeed ref={(cmp) => { if (cmp) { this.storyFeed = cmp } }} fetchingStoryFeed={isLoading} />
                <StoryModal updateStoryFeed={this.updateStoryFeed} />
                <CreatorModal updateStoryFeed={this.updateStoryFeed} />

                <Container className="HomePage__PostContainer" white={false}>
                    {isLoading ?
                        Array(3).fill(0).map((_, index) => <PostPlaceholder key={index} />)
                        :
                        posts.map((post) => <Post {...post} key={post.id} showImageGrid />)
                    }
                    {isFetchingMorePost && <PostPlaceholder />}
                </Container>
            </Container>
        );
    }
}

const mapDispatchToProps = (dispatch) => ({
    closeModal: () => dispatch(uiActions.closeStoryModal())
});

export default connect(null, mapDispatchToProps)(HomePage);