import React, { Component } from 'react'
import  StoryCard  from './StoryCard/StoryCard';

import "./StoryFeed.css";

class StoryFeed extends Component {
    render() {
        return (
            <div className="StoryFeed">
                <div className="StoryFeed__Header">
                    <b>Story</b>
                    <i>See all</i>
                </div>

                <div className="StoryFeed__Main">
                    <StoryCard text="Nguyen Trong Nghia" isMy />
                    <StoryCard text="Nguyen Trong" isActive/>
                    <StoryCard text="Nguyen" />
                </div>
            </div>
        )
    }
}

export default StoryFeed;