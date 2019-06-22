import React, { Component } from 'react'
import { connect } from 'react-redux';

import StoryCard from './StoryCard/StoryCard';

import "./StoryFeed.css";

class StoryFeed extends Component {
    render() {
        const { firstSBNode } = this.props;
        const storyCards = [];

        let sbNode = firstSBNode;
        while (sbNode !== null) {
            storyCards.push(<StoryCard key={sbNode.getValue().userID} sbNode={sbNode} />);
            sbNode = sbNode.getNext();
        }

        return (
            <div className="StoryFeed">
                <div className="StoryFeed__Header">
                    <b>Story</b>
                    <i>See all</i>
                </div>

                <div className="StoryFeed__Main">
                    {storyCards}
                </div>
            </div>
        )
    }
}

const mapStateToProps = ({ storyReducer: { storyFeedManager } }) => ({ firstSBNode: storyFeedManager.getFirstSBNode() });

export default connect(mapStateToProps)(StoryFeed);