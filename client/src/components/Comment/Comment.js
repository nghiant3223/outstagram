import React from 'react';
import { Comment as SemanticComment, Icon } from "semantic-ui-react";

import "./Comment.css";
import Dot from '../Dot/Dot';
import Avatar from '../Avatar/Avatar';

function Comment() {
    return (
        <SemanticComment.Group className="Comment">
            <SemanticComment>
                <div className="Comment__Container">
                    <div className="Comment__AvatarContainer">
                        <Avatar />
                    </div>

                    <div>
                        <SemanticComment.Content>
                            <SemanticComment.Author>Tom Lukic</SemanticComment.Author>
                            <SemanticComment.Text>
                                This will be great for business reports. I will definitely download this.
                            </SemanticComment.Text>

                            <SemanticComment.Actions>
                                <SemanticComment.Action>Like</SemanticComment.Action>
                                <SemanticComment.Action>Reply</SemanticComment.Action>
                                <Dot style={{ marginLeft: 0 }} />
                                3 hour
                    </SemanticComment.Actions>
                        </SemanticComment.Content>
                    </div>
                </div>
            </SemanticComment>
        </SemanticComment.Group>
    )
}

export default Comment;