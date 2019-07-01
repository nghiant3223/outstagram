import React, { Component } from 'react';
import { Button } from 'semantic-ui-react';

import * as userServices from '../../../services/user.service';

import "./UserInfo.css";

class UserInfo extends Component {
    constructor(props) {
        super(props);

        this.state = {
            followed: props.user.followed
        }
    }

    componentDidUpdate(prevProps) {
        if (prevProps.user.id !== this.props.user.id) {
            this.setState({ followed: this.props.user.followed });
        }
    }

    onFollowClick = () => {
        const { followed } = this.state;
        const { user } = this.props;

        if (!followed) {
            userServices.followUser(user.id);
        } else {
            userServices.unfollowUser(user.id);
        }

        this.setState((prevState) => ({ followed: !prevState.followed }));
    }

    render() {
        const { followed } = this.state;
        const { user } = this.props;

        return (
            <div className="InfoContainer">
                <div className="InfoHeader">
                    <div className="InfoHeader__Fullname">{user.fullname}</div>
                    {!user.isMe &&
                        (<div className="InfoHeader__Button">
                            <Button compact size='tiny' toggle active={followed} onClick={this.onFollowClick}>{followed ? "Following" : "Follow"}</Button>
                        </div>)
                    }
                </div>

                <div className="InfoItemContainer">
                    <div className="InfoItem">
                        <div className="InfoItem__Title">{user.followerCount} </div>
                        <div className="InfoItem__More">Followers</div>
                    </div>
                    <div className="InfoItem">
                        <div className="InfoItem__Title">{user.followingCount}</div>
                        <div className="InfoItem__More">Followings</div>
                    </div>
                    <div className="InfoItem">
                        <div className="InfoItem__Title">50</div>
                        <div className="InfoItem__More">Posts</div>
                    </div>
                </div>
            </div>
        )
    }
}


export default UserInfo;