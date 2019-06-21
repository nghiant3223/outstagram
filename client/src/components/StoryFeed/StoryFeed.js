import React, { Component } from 'react'
import { connect } from 'react-redux';

import StoryCard from './StoryCard/StoryCard';

import "./StoryFeed.css";

class StoryFeed extends Component {
    render() {
        const { storyBoards, currentUserID } = this.props;


        return (
            <div className="StoryFeed">
                <div className="StoryFeed__Header">
                    <b>Story</b>
                    <i>See all</i>
                </div>

                <div className="StoryFeed__Main">
                    {storyBoards.map((board) => <StoryCard key={board.userID} isMy={board.userID === currentUserID} latestStory={board.stories[0]} isActive={board.hasNewStory} text={board.fullname} />)}
                </div>
            </div>
        )
    }
}

const mapStateToProps = ({ storyReducer: { storyBoards }, authReducer: { userID: currentUserID } }) => ({ storyBoards, currentUserID });

export default connect(mapStateToProps)(StoryFeed);