import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux'
import { Comment as SemanticReply } from "semantic-ui-react";
import { Link } from 'react-router-dom';
import pruralize from "pluralize";

import Dot from '../Dot/Dot';
import Avatar from '../Avatar/Avatar';
import { getTimeDiffFromNow } from '../../utils/time';
import * as reactorActions from "../../actions/reactor.action";
import * as reactableServices from "../../services/reactable.service";
import * as commentService from "../../services/comment.service";

import "./Reply.css";
import ClickableText from '../ClickableText/ClickableText';
import UserInfoPopup from '../UserInfoPopup/UserInfoPopup';

class Reply extends Component {
    constructor(props) {
        super(props);

        this.state = {
            reply: undefined,
            reacted: props.reacted,
            reactCount: props.reactCount
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

        this.setState((prevState) => ({ reacted: !prevState.reacted, reactCount: prevState.reactCount + (reacted ? -1 : 1) }));
    }

    onReplyClick = () => {
        const { focusReplyInput } = this.props;
        focusReplyInput();
    }

    onNumberOfLikeClick = () => {
        const { openReactorModal, isNew } = this.props;

        if (!isNew) {
            const { reactableID } = this.props;
            openReactorModal(reactableID)
        } else {
            const { reactableID } = this.state.comment;
            openReactorModal(reactableID)
        }
    }

    render() {
        const { reply, reactCount } = this.state;

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
                            <UserInfoPopup username={ownerUsername} trigger={<Link to={`/${ownerUsername}`}><Avatar width="2.75rem" userID={ownerID} /></Link>} />
                        </div>

                        <div style={{ width: "100%" }}>
                            <SemanticReply.Content>
                                <div className="Comment__ContentContainer">
                                    <span className="Comment__AuthorName">
                                        <UserInfoPopup username={ownerUsername} trigger={<Link to={`/${ownerUsername}`} className="Fullname">{ownerFullname}</Link>} />
                                    </span>
                                    {content}
                                </div>

                                {(!isNew || reply !== undefined) &&
                                    <SemanticReply.Actions>
                                        {this.renderReact()}
                                        <SemanticReply.Action><div onClick={this.onReplyClick}>Reply</div></SemanticReply.Action>
                                        <Dot style={{ marginLeft: 0 }} />
                                        {getTimeDiffFromNow(createdAt)}
                                        {reactCount > 0 &&
                                            <Fragment>
                                                <Dot />
                                                <ClickableText onClick={this.onNumberOfLikeClick}> {reactCount} {pruralize("Like", reactCount)}</ClickableText>
                                            </Fragment>}
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

const mapDispatchToProps = (dispatch) => ({
    openReactorModal: (id) => dispatch(reactorActions.openModal(id))
});

export default connect(null, mapDispatchToProps)(Reply);