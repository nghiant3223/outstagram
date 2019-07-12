import React from 'react';
import PropTypes from 'prop-types';

import { connect } from 'react-redux';
import { Icon } from 'semantic-ui-react';

import "./FeedbackSummary.css";

const REACTOR_DISPLAY_COUNT = 3;

function FeedbackSummary(props) {
    const { commentCount, user, reacted, reactCount, displayCommentCount } = props;
    const reactors = props.reactors || [];
    const displayReactors = [];
    let reactorString = "";

    if (reactors.length > 0) {
        if (reacted) {
            displayReactors.push("You");
            for (let i = reactors[0].id === user.id ? 1 : 0; i < REACTOR_DISPLAY_COUNT - 1 && i < reactors.length; i++) {
                displayReactors.push(reactors[i].fullname);
            }
        } else {
            for (let i = reactors[0].id === user.id ? 1 : 0; i < REACTOR_DISPLAY_COUNT && i < reactors.length; i++) {
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
                {reactorString && <Icon name={"heart"} color="red" inverted />}{reactorString}
            </div>

            <div className="FeedbackSummary__Right">
                {displayCommentCount > 0 && `${displayCommentCount}/${commentCount} ${commentCount > 1 ? "comments" : "comment"}`}
            </div>
        </div >
    )
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

export default connect(mapStateToProps)(FeedbackSummary);