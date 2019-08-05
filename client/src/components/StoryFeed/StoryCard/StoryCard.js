import React from 'react';
import { connect } from 'react-redux';

import * as storyConfigs from "../../../configs/story.config";
import * as creatorActions from "../../../actions/creator.action";
import * as uiActions from '../../../actions/ui.action';

import "./StoryCard.css";
import addIcon from '../../../images/add.png';
import Avatar from '../../Avatar/Avatar';
import { noAuthStatic } from '../../../axios';

function storyCard(props) {
    const { sbNode, openStoryFeedModal, openCreatorModal } = props;

    noAuthStatic("/images/others")

    const sb = sbNode.getValue(),
        isMy = sb.isMy,
        userID = sb.userID,
        text = sb.fullname,
        stories = sb.stories,
        isActive = sb.hasNewStory,
        backgroundStyle = stories === null ?
            { background: "linear-gradient(0deg, rgba(255,255,255,1) 25%, rgba(196,196,196,1) 100%)" } :
            { backgroundImage: `url("/images/${sb.stories[0][storyConfigs.STORY_CARD_SIZE]}")` };

    const circleIcon = isMy ?
        (<div className="StoryCard__Circle StoryCard__Add" onClick={openCreatorModal}> <img src={addIcon} alt="plus icon" /> </div>) :
        (<div className="StoryCard__Circle"><Avatar isActive={isActive} userID={userID} /></div>)

    return (
        <div
            className="StoryCard"
            style={{ ...backgroundStyle }}>
            {circleIcon}
            <div className="StoryCard__Background" onClick={() => openStoryFeedModal(sbNode)} ></div>
            <b className="StoryCard__Text">{isMy ? "Add your story" : text}</b>
        </div>
    )
}

const mapDispatchToProps = (dispatch) => ({
    openStoryFeedModal: (sbNode) => dispatch(uiActions.openStoryModal(sbNode)),
    openCreatorModal: () => dispatch(creatorActions.openModal("STORY"))
});

export default connect(null, mapDispatchToProps)(storyCard);