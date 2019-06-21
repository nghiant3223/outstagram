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
    const { isMy, isActive, text, latestStory, openModal } = props;

    console.log(latestStory);
    console.log(storyConfigs.STORY_CARD_SIZE);
    console.error(latestStory[storyConfigs.STORY_CARD_SIZE]);
    return (
        <div
            className="StoryCard"
            style={{ backgroundImage: `url("/images/${latestStory[storyConfigs.STORY_CARD_SIZE]}")`, cursor: "pointer" }}
            onClick={openModal}>
            {
                isMy ?
                    <div className="StoryCard__Circle StoryCard__Add">
                        <img src={addIcon} />
                    </div>
                    :
                    <Avatar isActive={isActive} style={{ position: "absolute", top: "0.5em", left: "0.5em" }} />
            }
            <b className="StoryCard__Text">{isMy ? "Add your story" : text}</b>
        </div>
    )
}

const mapDispatchToProps = (dispatch) => ({
    openModal: () => dispatch(uiActions.openStoryModal())
});

export default connect(null, mapDispatchToProps)(StoryCard);