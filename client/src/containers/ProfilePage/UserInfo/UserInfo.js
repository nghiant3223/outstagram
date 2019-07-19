import React, { Component } from 'react';
import { Button } from 'semantic-ui-react';

import * as userServices from '../../../services/user.service';

import "./UserInfo.css";
import FollowButton from '../../../components/FollowButton/FollowButton';

class UserInfo extends Component {
    constructor(props) {
        super(props);

        this.state = {
            followed: props.user.followed,
            followerCount: props.user.followerCount
        }
    }

    componentDidUpdate(prevProps) {
        if (prevProps.user.id !== this.props.user.id) {
            const { followed, followerCount } = this.props.user;
            this.setState({ followed, followerCount });
        }
    }

    onFollowClick = () => {
        const { followed } = this.state;
        const { user } = this.props;

        if (!followed) {
            userServices.followUser(user.id);
            this.setState((prevState) => ({ followerCount: prevState.followerCount + 1 }));
        } else {
            userServices.unfollowUser(user.id);
            this.setState((prevState) => ({ followerCount: prevState.followerCount - 1 }));
        }

        this.setState((prevState) => ({ followed: !prevState.followed }));
    }

    render() {
        const { followed, followerCount } = this.state;
        const { user } = this.props;

        return (
            <div className="InfoContainer">
                <div className="InfoHeader">
                    <div className="InfoHeader__Fullname">{user.fullname}</div>
                    {!user.isMe &&
                        (<div className="InfoHeader__Button">
                            <FollowButton followed={followed} userID={user.id} />
                        </div>)
                    }
                </div>

                <div className="InfoItemContainer">
                    <div className="InfoItem">
                        <div className="InfoItem__Title">{followerCount}</div>
                        <div className="InfoItem__More">Followers</div>
                    </div>
                    <div className="InfoItem">
                        <div className="InfoItem__Title">{user.followingCount}</div>
                        <div className="InfoItem__More">Followings</div>
                    </div>
                    <div className="InfoItem">
                        <div className="InfoItem__Title">{user.postCount}</div>
                        <div className="InfoItem__More">Posts</div>
                    </div>
                </div>
            </div>
        )
    }
}


export default UserInfo;