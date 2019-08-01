import React, { useState } from 'react';
import { Link } from 'react-router-dom';
import moment from 'moment';
import { List, Header, Placeholder, Popup, Button, Icon } from 'semantic-ui-react'
import pluralize from "pluralize";

import * as userServices from "../../services/user.service";
import Avatar from '../Avatar/Avatar';

import "./UserInfoPopup.css";
import { noAuthStatic } from '../../axios';
import FollowButton from '../FollowButton/FollowButton';

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
        userServices.getUser(username).then(({ data: { data: { user } } }) => setUser(user)).catch((e) => console.log(e));
    }

    return (
        <Popup
            style={{ padding: 0 }}
            hoverable
            on='hover'
            onClose={onClose}
            onOpen={onOpen}
            popperDependencies={[!!user]}
            trigger={trigger}
            wide
        >
            {user === null ? (
                <div className="UserInfoPopUp__Container">
                    <Placeholder style={{ minWidth: '200px' }}>
                        <Placeholder.Header image>
                            <Placeholder.Line />
                            <Placeholder.Line length='medium' />
                        </Placeholder.Header>
                    </Placeholder>
                </div>

            ) : (
                    <div className="UserInfoPopUp__Container">
                        <div className="UserInfoPopUp__Container__CoverContainer" style={{ backgroundImage: `url(${noAuthStatic('/images/others/' + user.coverImageID, { size: "big" })})` }} >
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
                )}
        </Popup>
    )
}

export default UserInfoPopup
