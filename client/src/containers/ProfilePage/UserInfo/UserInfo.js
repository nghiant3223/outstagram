import React, { Component } from 'react';
import { Button } from 'semantic-ui-react';

import "./UserInfo.css";

class UserInfo extends Component {
    constructor(props) {
        super(props);
        this.state = {
            followed: false
        }
    }

    onFollowClick = () => {
        this.setState((prevState) => ({ followed: !prevState.followed }));
    }

    render() {
        const { followed } = this.state;

        return (
            <div className="InfoContainer">
                <div className="InfoHeader">
                    <div className="InfoHeader__Fullname">Clain Lannister</div>
                    <Button className="InfoHeader__Follow" compact size='tiny' toggle active={followed} onClick={this.onFollowClick}>{followed ? "Following" : "Follow"}</Button>
                </div>

                <div className="InfoItemContainer">
                    <div className="InfoItem">
                        <div className="InfoItem__Title">50</div>
                        <div className="InfoItem__More">Followers</div>
                    </div>
                    <div className="InfoItem">
                        <div className="InfoItem__Title">50</div>
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