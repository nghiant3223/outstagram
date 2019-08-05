import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import moment from 'moment';
import { List, Header, Popup, Button, Icon } from 'semantic-ui-react'
import pluralize from "pluralize";

import { noAuthStatic } from '../../axios';
import * as userServices from "../../services/user.service";

import Avatar from '../Avatar/Avatar';
import FollowButton from '../FollowButton/FollowButton';
import UserInfoPopUpPlaceholder from './Placeholder';

import "./UserInfoPopup.css";

const UserInfoPopup = (props) => {
    const { username, trigger } = props;
    const [user, setUser] = useState(null)

    const renderSubheader = (user) => {
        let infoList = [];

        if (user.followerCount > 0) {
            infoList.push({ icon: "users", content: `${user.followerCount} ${pluralize("follower")}` });
        }

        if (user.postCount > 0) {
            infoList.push({ icon: "newspaper", content: `${user.postCount} ${pluralize("post")}` });
        }

        infoList.push({ icon: "calendar alternate outline", content: `Join at ${moment(user.createdAt).calendar()}` });

        if (infoList.length > 2) {
            infoList = infoList.slice(0, 2);
        }

        return infoList.map((info, index) => (
            <List.Item key={index}>
                <List.Icon name={info.icon} color="grey" />
                <List.Content>{info.content}</List.Content>
            </List.Item>)
        );
    }

    const onClose = () => {
        setUser(null);
    }

    const onOpen = () => {
        userServices.getUser(username)
            .then(({ data: { data: { user } } }) => setUser(user))
            .catch((e) => console.log(e));
    }

    return (
        <Popup
            wide
            hoverable
            on='hover'
            onOpen={onOpen}
            onClose={onClose}
            trigger={trigger}
            style={{ padding: 0 }}
            popperDependencies={[!!user]}
        >
            {user === null ? <UserInfoPopUpPlaceholder /> : (
                <Link to={`/${user.username}`}>
                    <div className="UserInfoPopUp__Container">
                        <div className="UserInfoPopUp__Container__CoverContainer" style={{ background: `url(${noAuthStatic('/images/others/' + user.coverImageID, { size: "big" })})` }} >
                            <div className="UserInfoPopUp__Container__CoverContainer__Avatar">
                                <Avatar width="100px" userID={user.id} />
                            </div>
                        </div>

                        <div className="UserInfoPopUp__Container__DescriptionContainer">
                            <Header as='h4' content={user.fullname} />
                            <List className="List">
                                {renderSubheader(user)}
                            </List>
                        </div>

                        {!user.isMe && <div className="UserInfoPopup__Container__Actions">
                            <FollowButton followed={user.followed} userID={user.id} size="small" basic={true} />
                            <Link to={`/messages/${user.username}`}><Button basic size="small"><Icon name="facebook messenger" />Message</Button></Link>
                        </div>}
                    </div>
                </Link>
            )}
        </Popup>
    )
}

export default UserInfoPopup;
