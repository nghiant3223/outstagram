import React from 'react';

import Avatar from '../Avatar/Avatar';

import "./PostHeader.css";

export default function PostHeader({ userID, fullname, createdAt }) {
    return (
        <div className="PostHeader">
            <div className="PostHeader__Avatar">
                <Avatar width="2.75rem" />
            </div>

            <div className="PostHeader__Info">
                <div className="PostHeader__Info__Fullname">{fullname}</div>
                <div className="PostHeader__Info__CreatedAt">{createdAt}</div>
            </div>
        </div>
    )
}
