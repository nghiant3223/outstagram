import React, { useState } from 'react';
import { connect } from 'react-redux'
import { Button, Icon } from 'semantic-ui-react';

import * as userServices from "../../services/user.service";
import * as actionTypes from "../../actions/auth.action";

import "./FollowButton.css";

function FollowButton(props) {
    const [followed, setFollowed] = useState(props.followed)

    const toggleFollow = () => {
        const { userID, updateFollowingCount } = props;

        if (!followed) {
            userServices.followUser(userID);
        } else {
            userServices.unfollowUser(userID);
        }

        setFollowed(!followed);
        updateFollowingCount(followed);
    }

    return (
        followed ?
            <Button icon onClick={toggleFollow} className="FollowButton">
                <Icon.Group>
                    <Icon name='user' />
                    <Icon corner name='check' />
                </Icon.Group>
                <span className="FollowButton__Text">Following</span>
            </Button>
            :
            <Button icon onClick={toggleFollow}>
                <Icon.Group>
                    <Icon name='user' />
                    <Icon corner name='plus' />
                </Icon.Group>
                <span className="FollowButton__Text">Follow</span>
            </Button>

    );
}

const mapDispatchToProps = (dispatch) => ({
    updateFollowingCount: (followed) => dispatch(actionTypes.updateUserFollowingCount(!followed))
});

export default connect(null, mapDispatchToProps)(FollowButton);