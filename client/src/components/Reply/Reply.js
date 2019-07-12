import React, { Component } from 'react';
import { Comment as SemanticReply, Icon } from "semantic-ui-react";
import { Link } from 'react-router-dom';

import Dot from '../Dot/Dot';
import Avatar from '../Avatar/Avatar';
import { getDiffFromPast } from '../../utils/time';
import * as postServices from "../../services/post.service";
import * as reactableServices from "../../services/reactable.service";
import * as commentService from "../../services/comment.service";
import PostInput from '../PostInput/PostInput';

import "./Reply.css";

class Reply extends Component {
    constructor(props) {
        super(props);

        this.state = {
            reply: undefined,
            reacted: props.reacted
        }
    }

    async componentDidMount() {
        const { isNew, replaceReply, id, postCmtableID, cmtID, content } = this.props;
        if (isNew) {
            try {
                const { data: { data: { reply } } } = await commentService.createReply(postCmtableID, cmtID, content);
                this.setState({ reply });
                replaceReply(id, reply);
            } catch (e) {
                console.log("Cannot create reply", e);
            }
        }
    }

    onLikeClick = () => {
        const reactableID = this.props.reactableID || this.state.reply.reactableID;
        const reacted = this.state.reacted;

        if (reacted) {
            reactableServices.unreact(reactableID);
        } else {
            reactableServices.react(reactableID);
        }

        this.setState((prevState) => ({ reacted: !prevState.reacted }));
    }

    onReplyClick = () => {
        const { focusReplyInput } = this.props;
        focusReplyInput();
    }

    render() {
        const { reply } = this.state;

        // If reply is not submitted
        if (reply === undefined) {
            var { ownerID, ownerFullname, content, createdAt, ownerUsername, isNew } = this.props;
        } else {
            var { ownerID, ownerFullname, content, createdAt, ownerUsername } = this.state.reply;
        }

        return (
            <SemanticReply.Group className="Reply" style={{ paddingTop: 0, paddingBottom: 0 }}>
                <SemanticReply>
                    <div className="Reply__Container">
                        <div className="Reply__AvatarContainer">
                            <Avatar userID={ownerID} width="2em" />
                        </div>

                        <div style={{ width: "100%" }}>
                            <SemanticReply.Content>
                                <div className="Comment__ContentContainer">
                                    <span className="Comment__AuthorName">
                                        <Link to={`/${ownerUsername}`}>{ownerFullname}</Link>
                                    </span>
                                    {content}
                                </div>

                                {(!isNew || reply !== undefined) &&
                                    <SemanticReply.Actions>
                                        {this.renderReact()}
                                        <SemanticReply.Action><div onClick={this.onReplyClick}>Reply</div></SemanticReply.Action>
                                        <Dot style={{ marginLeft: 0 }} />
                                        {getDiffFromPast(createdAt)}
                                    </SemanticReply.Actions>}
                            </SemanticReply.Content>
                        </div>
                    </div>
                </SemanticReply>
            </SemanticReply.Group>
        )
    }

    renderReact() {
        const { reacted } = this.state;
        const style = { color: "var(--primary-color)", fontWeight: "bold", fontSize: "0.95em" }

        if (reacted) {
            return <SemanticReply.Action>
                <div onClick={this.onLikeClick} style={style}>Like</div>
            </SemanticReply.Action>
        }

        return <SemanticReply.Action>
            <div onClick={this.onLikeClick}>Like</div>
        </SemanticReply.Action>
    }
}

export default Reply;