import React, { Component } from 'react';
import PropTypes from 'prop-types';
import { connect } from 'react-redux';

import * as uiActions from '../../../actions/ui.action';

import "./StoryCard.css";
import addIcon from '../../../images/add.png';
import defaultAvatar from '../../../images/avatar.png';
import defaultBackground from '../../../images/x.png';
import Avatar from '../../Avatar/Avatar';

function StoryCard(props) {
    const { isMy, isActive, text, avatar, story, openModal } = props;

    return (
        <div
            className="StoryCard"
            style={{ backgroundImage: `url(${story || defaultBackground})`, cursor: "pointer" }}
            onClick={openModal}>
            {
                isMy ?
                    <div className="StoryCard__Circle StoryCard__Add">
                        <img src={addIcon} />
                    </div>
                    :
                    <Avatar isActive={isActive} avatar={avatar} style={{ position: "absolute", top: "0.5em", left: "0.5em" }} />
            }
            <b className="StoryCard__Text">{isMy ? "Add your story" : text}</b>
        </div>
    )
}

const mapDispatchToProps = (dispatch) => ({
    openModal: () => dispatch(uiActions.openStoryModal())
});

export default connect(null, mapDispatchToProps)(StoryCard);