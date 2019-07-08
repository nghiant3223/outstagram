import React from 'react';

import './Avatar.css';
import defaultAvatar from '../../images/avatar.png';
import { noAuthStatic } from '../../axios';

export default function Avatar(props) {
    const { isActive, userID, size = "small" } = props;

    let className = "Avatar";
    if (isActive === true) {
        className += " Avatar--Active";
    } else if (isActive === false) {
        className += " Avatar--Inactive"
    }

    const avatarURL = userID ? noAuthStatic('/images/avatars/' + userID, { size }) : defaultAvatar
    return (
        <div className={className} style={{ background: `url(${avatarURL})`, backgroundPosition: "50% 50%", backgroundSize: "cover" }} />
    )
}
