import React, { Component, Fragment } from 'react';
import { connect } from 'react-redux'
import pruralize from "pluralize";
import { Comment as SemanticComment, Icon } from "semantic-ui-react";
import { Link } from 'react-router-dom';

import Dot from '../Dot/Dot';
import Avatar from '../Avatar/Avatar';
import { getTimeDiffFromNow } from '../../utils/time';
import * as reactorActions from "../../actions/reactor.action";
import * as commentableServices from "../../services/commentable.service";
import * as reactableServices from "../../services/reactable.service";
import * as commentServices from "../../services/comment.service";
import { genUID } from "../../utils/lang";
import PostInput from '../PostInput/PostInput';

import "./Comment.css";
import ClickableText from '../ClickableText/ClickableText';
import Loading from "../Loading/Loading";
import Reply from '../Reply/Reply';
import UserInfoPopup from '../UserInfoPopup/UserInfoPopup';

class Comment extends Component {
    constructor(props) {
        super(props);

        this.state = {
            comment: undefined,
            reacted: props.reacted,
            replyInputShow: false,
            replies: [],
            replyCount: props.replyCount,
            isLoadingMoreReply: false,
            reactCount: props.reactCount
        }
    }

    async componentDidMount() {
        const { isNew, replaceComment, id, commentableID, content } = this.props;
        if (isNew) {
            try {
                const { data: { data: { comment } } } = await commentableServices.commentPost(commentableID, content);
                this.setState({ comment });
                replaceComment(id, comment);
            } catch (e) {
                console.log("Cannot create comment");
            }
        }
    }

    onLikeClick = () => {
        const reactableID = this.props.reactableID || this.state.comment.reactableID;
        const reacted = this.state.reacted;

        if (reacted) {
            reactableServices.unreact(reactableID);
        } else {
            reactableServices.react(reactableID);
        }

        this.setState((prevState) => ({ reacted: !prevState.reacted, reactCount: prevState.reactCount + (reacted ? -1 : 1) }));
    }

    onReplyClick = () => {
        const { replyInputShow } = this.state;

        if (!replyInputShow) {
            this.setState({ replyInputShow: true });
        } else {
            this.focusReplyInput();
        }
    }

    // Create temporary reply
    onReplySubmit = (content) => {
        const { commentableID } = this.props;
        const { id, fullname, username } = this.props.user;
        const reply = { content, ownerFullname: fullname, ownerUsername: username, ownerID: id, isNew: true, id: genUID(), commentableID, reacted: false, replyCount: 0, reactCount: 0 }
        this.setState((prevState) => ({ replies: [...prevState.replies, reply], replyCount: prevState.replyCount + 1 }));
    }

    // Replace the temporary reply by the newly created reply
    replaceReply = (uid, newCreatedReply) => {
        const { replies } = this.state;
        const reply = replies.find(reply => reply.id === uid);

        if (!reply) {
            throw new Error("Reply does not exist");
        }

        // Copy property from newCreatedReply to current reply in the state;
        for (var k in newCreatedReply) {
            // IMPORTANT: Ignore id field to prevent changing Reply's key, which cause a new Reply is created
            if (k !== "id") {
                reply[k] = newCreatedReply[k];
            }
        }
    }

    onMoreReplyClick = async () => {
        const { commentableID, isNew } = this.props;
        const { replies } = this.state;

        if (!isNew) {
            var { id: cmtID } = this.props;
        } else {
            var { id: cmtID } = this.state.comment;
        }

        this.setState({ isLoadingMoreReply: true });

        try {
            const { data: { data: { replies: moreReplies } } } = await commentServices.getReply(commentableID, cmtID, 5, replies.length);
            this.setState((prevState) => ({ replies: [...moreReplies, ...prevState.replies] }));
        } catch (e) {
            console.log("Fetching more reply failed", e);
        } finally {
            this.setState({ isLoadingMoreReply: false });
        }
    }

    focusReplyInput = () => {
        this.replyInput.focus();
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
        const { comment, replyInputShow, replies, replyCount, isLoadingMoreReply, reactCount } = this.state;

        // If comment is not submitted
        if (comment === undefined) {
            var { ownerID, ownerFullname, content, createdAt, ownerUsername, commentableID: postCmtableID, id, isNew, userID } = this.props;
        } else {
            var { ownerID, ownerFullname, content, createdAt, ownerUsername, commentableID: postCmtableID, id } = this.state.comment;
        }

        // `replyCount` is the total number of replies
        // `replies.length` is the number of replies which are displayed
        const notFetchedReplyCount = replyCount - replies.length;

        return (
            <SemanticComment.Group className="Comment">
                <SemanticComment>
                    <div className="Comment__Container">
                        <div className="Comment__AvatarContainer">
                            <UserInfoPopup username={ownerUsername} trigger={<Link to={`/${ownerUsername}`}><Avatar width="2.5rem" userID={ownerID} /></Link>} />
                        </div>

                        <div style={{ width: "100%" }}>
                            <SemanticComment.Content>
                                <div className="Comment__ContentContainer">
                                    <span className="Comment__AuthorName">
                                        <UserInfoPopup username={ownerUsername} trigger={<Link to={`/${ownerUsername}`} className="Fullname Fullname--Medium">{ownerFullname}</Link>} />
                                    </span>
                                    {content}
                                </div>

                                {(!isNew || comment !== undefined) &&
                                    <SemanticComment.Actions className="Comment__ActionContainer">
                                        {this.renderReact()}
                                        <SemanticComment.Action>
                                            <div onClick={this.onReplyClick}>Reply</div>
                                        </SemanticComment.Action>
                                        <Dot style={{ marginLeft: 0 }} />
                                        {getTimeDiffFromNow(createdAt)}
                                        {reactCount > 0 &&
                                            <Fragment>
                                                <Dot />
                                                <ClickableText onClick={this.onNumberOfLikeClick}> {reactCount} {pruralize("Like", reactCount)}</ClickableText>
                                            </Fragment>}
                                    </SemanticComment.Actions>}

                                {notFetchedReplyCount > 0 &&
                                    <div onClick={this.onMoreReplyClick} className="Comment_MoreReplies" style={{ marginTop: "0.5em" }}>
                                        <Icon name="reply" color="blue" className="MoreCommentIcon" />
                                        <ClickableText>{notFetchedReplyCount} {pruralize("Reply", notFetchedReplyCount)}</ClickableText>
                                        {isLoadingMoreReply && <Loading />}
                                    </div>}

                                {replies.map((reply) => <Reply {...reply} postCmtableID={postCmtableID} cmtID={id} key={reply.id} replaceReply={this.replaceReply} focusReplyInput={this.focusReplyInput} />)}
                            </SemanticComment.Content>

                            {(replyInputShow || replies.length > 0) && <PostInput userID={userID} style={{ paddingTop: 0, paddingBottom: 0, marginTop: "0.25em" }} inverted focusOnMount onSubmit={this.onReplySubmit} placeholder="Write your reply ..." ref={el => this.replyInput = el} />}
                        </div>
                    </div>
                </SemanticComment>
            </SemanticComment.Group>
        )
    }

    renderReact() {
        const { reacted } = this.state;
        const style = { color: "var(--primary-dark-color)", fontWeight: "bold", fontSize: "0.95em" }

        if (reacted) {
            return <SemanticComment.Action>
                <div onClick={this.onLikeClick} style={style}>Like</div>
            </SemanticComment.Action>
        }

        return <SemanticComment.Action>
            <div onClick={this.onLikeClick}>Like</div>
        </SemanticComment.Action>
    }
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

const mapDispatchToProps = (dispatch) => ({
    openReactorModal: (id) => dispatch(reactorActions.openModal(id))
});

export default connect(mapStateToProps, mapDispatchToProps)(Comment);