import React from 'react';

import './Avatar.css';
import defaultAvatar from '../../images/avatar.png';

export default function Avatar(props) {
    const { isActive, avatar, style } = props;

    let className = "Avatar";
    if (isActive === true) {
        className += " Avatar--Active";
    } else if (isActive === false) {
        className += " Avatar--Inactive"
    }

    return (
        <div className={className} style={style}>
            <img src={avatar || defaultAvatar} alt="User's avatar" />
        </div>
    )
}
