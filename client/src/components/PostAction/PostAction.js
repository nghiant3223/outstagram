import React from 'react';

import "./PostAction.css";
import { Icon, Grid } from 'semantic-ui-react';

function PostAction(props) {
    const { onReactClick, reacted, onCommentClick } = props;

    return (
        <div className="PostAction">
            <div className="PostAction__Action" onClick={onReactClick}><Icon name={reacted ? "heart" : "heart outline"} color={reacted ? "red" : "black"} inverted />Like</div>
            <div className="PostAction__Action" onClick={onCommentClick}><Icon name="comment outline" color="black" inverted /> Comment</div>
        </div>
    )
}

export default PostAction;