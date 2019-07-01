import React, { Component, Fragment, } from 'react';
import { connect } from 'react-redux';

import StoryFeedManager from '../../StoryFeedManager';
import * as storyActions from '../../actions/story.action';

import StoryCard from './StoryCard/StoryCard';
import StoryCardPlaceholder from './StoryCardPlaceholder/StoryCardPlaceholder';

import "./StoryFeed.css";

class StoryFeed extends Component {
    state = {
        shouldUpdate: false
    }

    update() {
        this.setState((prevState) => ({ shouldUpdate: !prevState.shouldUpdate }));
    }

    render() {
        const { shouldUpdate } = this.state;
        const { fetchingStoryFeed, displayFirstSB } = this.props;

        return (
            <div className="StoryFeed">
                <div className="StoryFeed__Header">
                    <b>Story</b>
                    <i className="StoryFeed__Header__SeeAll" onClick={displayFirstSB}>See all</i>
                </div>

                <div className="StoryFeed__Main">
                    {fetchingStoryFeed ?
                        (
                            <Fragment>
                                <StoryCardPlaceholder />
                                <StoryCardPlaceholder />
                                <StoryCardPlaceholder />
                                <StoryCardPlaceholder />
                                <StoryCardPlaceholder />
                            </Fragment>
                        ) :
                        StoryFeedManager.getInstance().map((sbNode) => <StoryCard key={sbNode.getValue().storyBoardID} sbNode={sbNode} shouldUpdate={shouldUpdate} />)
                    }
                </div>
            </div>
        );
    }
}

const mapStateToProps = ({ storyReducer: { isLoading, isModalOpen } }) => ({ isLoading, isModalOpen });

const mapDispatchToProps = (dispatch) => ({
    displayFirstSB: () => dispatch(storyActions.displayFirstSBNode())
});

export default connect(mapStateToProps, mapDispatchToProps, null, { forwardRef: true })(StoryFeed);