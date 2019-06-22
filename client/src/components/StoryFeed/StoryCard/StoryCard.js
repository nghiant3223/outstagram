import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

import * as storyConfigs from "../../../configs/story.config";
import * as uiActions from '../../../actions/ui.action';

import "./StoryCard.css";
import addIcon from '../../../images/add.png';
import Avatar from '../../Avatar/Avatar';

function StoryCard(props) {
    const { sbNode, openModal } = props;

    const sb = sbNode.getValue(),
        isMy = sb.isMy,
        text = sb.fullname,
        stories = sb.stories,
        isActive = sb.hasNewStory,
        backgroundStyle = stories === null ?
            { background: "linear-gradient(0deg, rgba(255,255,255,1) 25%, rgba(196,196,196,1) 100%)" } :
            { backgroundImage: `url("/images/${sb.stories[0][storyConfigs.STORY_CARD_SIZE]}")` };

    const circleIcon = isMy ?
        (<div className="StoryCard__Circle StoryCard__Add"> <img src={addIcon} /> </div>) :
        (<div className="StoryCard__Circle"><Avatar isActive={isActive} /></div>)

    return (
        <div
            className="StoryCard"
            style={{ ...backgroundStyle, cursor: "pointer" }}
            onClick={() => openModal(sbNode)} >
            {circleIcon}
            <b className="StoryCard__Text">{isMy ? "Add your story" : text}</b>
        </div>
    )
}

const mapDispatchToProps = (dispatch) => ({
    openModal: (sbNode) => dispatch(uiActions.openStoryModal(sbNode))
});

export default connect(null, mapDispatchToProps)(StoryCard);