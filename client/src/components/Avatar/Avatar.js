import React from 'react';

import './Avatar.css';
import defaultAvatar from '../../images/x.png';
import { noAuthStatic } from '../../axios';

export default function Avatar({ isActive, userID, size = "small", width = "2.5em", height }) {
    let className = "Avatar";
    if (isActive === true) {
        className += " Avatar--Active";
    } else if (isActive === false) {
        className += " Avatar--Inactive"
    }

    const avatarURL = userID ? noAuthStatic('/images/avatars/' + userID, { size }) : defaultAvatar;
    const style = {
        background: `url(${avatarURL})`,
        backgroundPosition: "50% 50%",
        backgroundSize: "cover",
        width: width, height: height || width
    }

    return <div className={className} style={style} />;
}
