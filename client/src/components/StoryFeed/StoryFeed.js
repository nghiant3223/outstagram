import React, { Component } from 'react'
import { connect } from 'react-redux';

import StoryCard from './StoryCard/StoryCard';

import "./StoryFeed.css";

class StoryFeed extends Component {
    render() {
        const { storyBoardLL, currentUserID } = this.props;
        const storyCards = [];

        let storyBoardNode = storyBoardLL.getHead();
        while (storyBoardNode !== null) {
            storyCards.push(<StoryCard key={storyBoardNode.getValue().userID} storyBoardNode={storyBoardNode} currentUserID={currentUserID} />);
            storyBoardNode = storyBoardNode.getNext();
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

const mapStateToProps = ({ storyReducer: { storyBoardLL }, authReducer: { userID: currentUserID } }) => ({ storyBoardLL, currentUserID, });

export default connect(mapStateToProps)(StoryFeed);