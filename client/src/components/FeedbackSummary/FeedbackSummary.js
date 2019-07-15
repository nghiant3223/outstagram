import React from 'react';

import { connect } from 'react-redux';
import { Icon } from 'semantic-ui-react';

import * as reactorActions from "../../actions/reactor.action";

import "./FeedbackSummary.css";
import ClickableText from '../ClickableText/ClickableText';

const REACTOR_DISPLAY_COUNT = 3;

function FeedbackSummary(props) {
    const { commentCount, user, reacted, reactCount, displayCommentCount, openModal, reactableID } = props;
    const reactors = props.reactors || [];
    const displayReactors = [];
    let reactorString = "";

    if (reactors.length > 0) {
        if (reacted) {
            displayReactors.push("You");
            for (let i = 1; i < REACTOR_DISPLAY_COUNT && i < reactors.length; i++) {
                displayReactors.push(reactors[i].fullname);
            }
        } else {
            for (let i = 0; i < REACTOR_DISPLAY_COUNT && i < reactors.length; i++) {
                displayReactors.push(reactors[i].fullname);
            }
        }

        reactorString = displayReactors.join(", ");
        const restReactorsCount = reactCount - displayReactors.length;

        if (restReactorsCount > 0) {
            reactorString = reactorString + " and " + restReactorsCount + " others";
        }
    }

    return (
        <div className="FeedbackSummary">
            <div className="FeedbackSummary__Left">
                {reactCount > 0 && <Icon name={"heart"} color="red" inverted />} <ClickableText onClick={() => openModal(reactableID)}>{reactorString}</ClickableText>
            </div>

            <div className="FeedbackSummary__Right">
                {displayCommentCount > 0 && `${displayCommentCount}/${commentCount} ${commentCount > 1 ? "comments" : "comment"}`}
            </div>
        </div >
    )
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

const mapDispatchToProps = (dispatch) => ({
    openModal: (id) => dispatch(reactorActions.openModal(id))
})

export default connect(mapStateToProps, mapDispatchToProps)(FeedbackSummary);