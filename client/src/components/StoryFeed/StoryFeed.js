import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux';

import StoryFeedManager from '../../StoryFeedManager';

import StoryCard from './StoryCard/StoryCard';
import StoryCardPlaceholder from './StoryCardPlaceholder/StoryCardPlaceholder';

import "./StoryFeed.css";

class StoryFeed extends Component {
    render() {
        const { isLoading } = this.props;

        return (
            <div className="StoryFeed">
                <div className="StoryFeed__Header">
                    <b>Story</b>
                    <i>See all</i>
                </div>

                <div className="StoryFeed__Main">
                    {isLoading ?
                        (
                            <Fragment>
                                <StoryCardPlaceholder />
                                <StoryCardPlaceholder />
                                <StoryCardPlaceholder />
                            </Fragment>
                        ) :
                        StoryFeedManager.getInstance().map((sbNode) => <StoryCard key={sbNode.getValue().userID} sbNode={sbNode} />)
                    }
                </div>
            </div>
        );
    }
}

const mapStateToProps = ({ storyReducer: { isLoading } }) => ({ isLoading });

export default connect(mapStateToProps)(StoryFeed);