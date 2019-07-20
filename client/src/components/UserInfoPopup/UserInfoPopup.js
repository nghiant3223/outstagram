import React from 'react'
import { Button, Header, Placeholder, Popup } from 'semantic-ui-react'
import pluralize from "pluralize";

import * as userServices from "../../services/user.service";
import Avatar from '../Avatar/Avatar';

import "./UserInfoPopup.css";

const UserInfoPopup = (props) => {
    const { username, trigger } = props;
    const [data, setData] = React.useState(null)

    const renderSubheader = (user) => {
        if (user.followerCount > 0) {
            return `${user.followerCount} ${pluralize("follower")}`
        }

        if (user.postCount > 0) {
            return `${user.postCount} ${pluralize("post")}`
        }

        return `Join on ${new Date(user.createdAt).toLocaleDateString("en-US")}`
    }

    return (
        <Popup
            on='hover'
            position="top left"
            onClose={() => {
                setData(null)
            }}
            onOpen={() => {
                setData(null)
                userServices.getUser(username).then(({ data: { data: user } }) => setData(user)).catch((e) => console.log(e));
            }}
            popperDependencies={[!!data]}
            trigger={trigger}
            wide
        >
            {data === null ? (
                <Placeholder style={{ minWidth: '200px' }}>
                    <Placeholder.Header image>
                        <Placeholder.Line />
                        <Placeholder.Line length='medium' />
                    </Placeholder.Header>
                </Placeholder>
            ) : (
                    <div className="UserInfoPopUp__Container">
                        <div>
                            <Avatar width="2.5em" userID={data.id} />
                        </div>
                        <div>
                            <Header as='h4' content={data.fullname} subheader={renderSubheader(data)} />
                            <p>{data.description}</p>
                        </div>
                    </div>
                )}
        </Popup>
    )
}

export default UserInfoPopup
