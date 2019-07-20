import React, { Component } from 'react';
import { connect } from 'react-redux';
import { Icon } from 'semantic-ui-react'

import GridImageContainer from '../GridImageContainer/GridImageContainer';
import Comment from '../Comment/Comment';
import FeedbackSummary from '../FeedbackSummary/FeedbackSummary';
import PostAction from '../PostAction/PostAction';
import PostInput from '../PostInput/PostInput';

import PostHeader from '../PostHeader/PostHeader';
import { getDiffFromPast } from '../../utils/time';
import { noAuthStatic } from '../../axios';
import * as reactableServices from "../../services/reactable.service";
import * as commentableServices from "../../services/commentable.service";

import "./Post.css";

import { genUID } from '../../utils/lang';
import Loading from '../Loading/Loading';
import ClickableText from '../ClickableText/ClickableText';
import PostDescription from '../PostDescription/PostDescription';

class Post extends Component {
    constructor(props) {
        super(props);

        this.state = {
            reacted: props.reacted,
            reactCount: props.reactCount,
            reactors: props.reactors || [],
            comments: props.comments || [],
            commentCount: props.commentCount,
            isLoadingMoreComment: false,
            description: props.description
        }
    }

    onReactClick = async () => {
        const { reacted } = this.state;
        const { reactableID } = this.props;
        const { id, fullname } = this.props.user;

        if (reacted) {
            reactableServices.unreact(reactableID);
        } else {
            reactableServices.react(reactableID);
        }

        this.setState((prevState) => {
            if (!prevState.reacted) {
                return {
                    reacted: !prevState.reacted,
                    reactCount: prevState.reactCount + 1,
                    reactors: [{ id, fullname }, ...prevState.reactors]
                }
            } else {
                return {
                    reacted: !prevState.reacted,
                    reactCount: prevState.reactCount - 1,
                    reactors: prevState.reactors.filter((reactor) => reactor.id !== id)
                }
            }
        });
    }

    // Create temporary comment
    onCommentSubmit = (content) => {
        const { commentableID } = this.props;
        const { id, fullname } = this.props.user;
        const comment = { content, ownerFullname: fullname, ownerID: id, isNew: true, id: genUID(), commentableID, reacted: false, replyCount: 0, reactCount: 0 }
        this.setState((prevState) => ({ comments: [...prevState.comments, comment], commentCount: prevState.commentCount + 1 }));
    }

    // Replace the temporary comment by the newly created comment
    replaceComment = (uid, newCreatedComment) => {
        const { comments } = this.state;
        const comment = comments.find(comment => comment.id === uid);

        if (!comment) {
            throw new Error("Comment does not exist");
        }

        // Copy property from newCreatedComment to current comment in the state;
        for (var k in newCreatedComment) {
            // IMPORTANT: Ignore id field to prevent changing Comment's key, which cause a new Comment is created
            if (k !== "id") {
                comment[k] = newCreatedComment[k];
            }
        }
    }

    onCommentButtonClick = () => {
        const { isTheater } = this.props;

        if (isTheater !== undefined) {
            this.postInput.scrollTo();
        } else {
            this.postInput.focus();
        }
    }

    onMoreCommentClick = async () => {
        const { commentableID } = this.props;
        const { comments } = this.state;
        this.setState({ isLoadingMoreComment: true });
        try {
            const { data: { data: { comments: moreComments } } } = await commentableServices.getComment(commentableID, 5, comments.length);
            this.setState((prevState) => ({ comments: [...moreComments, ...prevState.comments] }));
        } catch (e) {
            console.log("Fetching more comment failed", e);
        } finally {
            this.setState({ isLoadingMoreComment: false });
        }
    }

    render() {
        const { reacted, reactCount, reactors, comments, commentCount, isLoadingMoreComment } = this.state;
        const { id, images, imageID, ownerID, ownerFullname, ownerUsername, createdAt, content, commentableID, reactableID, viewableID, showImageGrid, user } = this.props;

        return (
            <div className="Post">
                <PostHeader fullname={ownerFullname} createdAt={getDiffFromPast(createdAt)} userID={ownerID} username={ownerUsername} />

                <div className="ThreaterContainer__InfoContainer__Description">
                    {this.renderDescription()}
                </div>

                {showImageGrid &&
                    <div>
                        <GridImageContainer images={images ? images : [{ id, imageID, commentableID, reactableID, viewableID, content }]} />
                    </div>}

                <div>
                    <FeedbackSummary reactableID={reactableID} reactors={reactors} commentCount={commentCount} reacted={reacted} reactCount={reactCount} displayCommentCount={comments.length} />
                </div>

                <div>
                    <PostAction onReactClick={this.onReactClick} reacted={reacted} onCommentClick={this.onCommentButtonClick} />
                </div>

                {
                    comments.length > 0 &&
                    <div className="ThreaterContainer__InfoContainer__CommentContainer">
                        {comments.length < commentCount &&
                            <div onClick={this.onMoreCommentClick} style={{ cursor: "pointer", marginTop: "0.5em" }}>
                                <Icon name="reply" color="blue" className="MoreCommentIcon" />
                                <ClickableText>See more comments</ClickableText>
                                {isLoadingMoreComment && <Loading />}
                            </div>}
                        {comments.map((comment) => <Comment userID={user.id} {...comment} key={comment.id} replaceComment={this.replaceComment} commentableID={commentableID} />)}
                    </div>
                }

                <div>
                    <PostInput userID={user.id} isThreater={true} inverted onSubmit={this.onCommentSubmit} ref={el => this.postInput = el} isCommentInput style={{ fontSize: "1.25em" }} placeholder="Write your comment ..." />
                </div>
            </div >
        );
    }

    renderDescription() {
        const { content, ownerID, user, id, showImageGrid, isPost } = this.props;
        return <PostDescription content={content} ownerID={ownerID} user={user} id={id} isPost={showImageGrid || isPost} />
    }
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

export default connect(mapStateToProps)(Post);