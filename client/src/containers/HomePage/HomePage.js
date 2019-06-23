import React, { Component } from 'react';
import { connect } from 'react-redux';

import * as storyActions from '../../actions/story.action';

import StoryFeed from '../../components/StoryFeed/StoryFeed';
import StoryModal from '../../components/StoryModal/StoryModal';

import './HomePage.css';

class HomePage extends Component {
    constructor(props) {
        super(props);
    }

    componentDidMount() {
        const { getStories } = this.props;

        getStories();
    }

    updateStoryFeed = () => {
        this.storyFeed.updateStoryFeed();
    }

    render() {
        return (
            <div>
                <StoryFeed ref={(cmp) => { if (cmp) { this.storyFeed = cmp } }} />
                <StoryModal updateStoryFeed={this.updateStoryFeed} />
            </div>
        );
    }
}

const mapDispatchToProps = (dispatch) => ({
    getStories: () => dispatch(storyActions.getStories())
});

export default connect(null, mapDispatchToProps)(HomePage);