import React, { Component } from 'react'

import StoryFeedManager from '../../StoryFeedManager';

import StoryCard from './StoryCard/StoryCard';

import "./StoryFeed.css";

class StoryFeed extends Component {
    render() {
        const firstSBNode = StoryFeedManager.getInstance().getFirstSBNode(), storyCards = [];

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

export default StoryFeed;