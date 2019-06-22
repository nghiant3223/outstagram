import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

import * as storyConfigs from "../../../configs/story.config";
import * as uiActions from '../../../actions/ui.action';

import "./StoryCard.css";
import addIcon from '../../../images/add.png';
import defaultAvatar from '../../../images/avatar.png';
import defaultBackground from '../../../images/x.png';
import Avatar from '../../Avatar/Avatar';

function StoryCard(props) {
    const { sbNode, openModal, isMy } = props;

    const sb = sbNode.getValue(),
        text = sb.fullname,
        stories = sb.stories,
        isActive = sb.hasNewStory,
        latestStoryURL = stories !== null ? `url("/images/${sb.stories[0][storyConfigs.STORY_CARD_SIZE]}")` : null;

    const circleIcon = isMy ?
        (<div className="StoryCard__Circle StoryCard__Add">  <img src={addIcon} /> </div>) :
        (<Avatar isActive={isActive} style={{ position: "absolute", top: "0.5em", left: "0.5em" }} />)

    return (
        <div
            className="StoryCard"
            style={{ backgroundImage: latestStoryURL, cursor: "pointer" }}
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