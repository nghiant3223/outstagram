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

class Post extends Component {
    constructor(props) {
        super(props);

        this.state = {
            reacted: props.reacted,
            reactCount: props.reactCount,
            comments: props.comments || [],
            commentCount: props.commentCount,
            isLoadingMoreComment: false
        }
    }

    onReactClick = async () => {
        const { reacted } = this.state;
        const { reactableID } = this.props;
        if (reacted) {
            reactableServices.unreact(reactableID);
        } else {
            reactableServices.react(reactableID);
        }
        this.setState((prevState) => ({ reacted: !prevState.reacted, reactCount: prevState.reactCount + (prevState.reacted ? -1 : 1) }));
    }

    // Create temporary comment
    onCommentSubmit = (content) => {
        const { commentableID } = this.props;
        const { id, fullname } = this.props.user;
        const comment = { content, ownerFullname: fullname, ownerID: id, isNew: true, id: genUID(), commentableID, reacted: false, replyCount: 0 }
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
        this.postInput.scrollTo();
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
        const { reacted, reactCount, comments, commentCount, isLoadingMoreComment } = this.state;
        const { images, ownerFullname, createdAt, reactors, content, commentableID } = this.props;

        return (
            <div className="Post">
                <PostHeader fullname={ownerFullname} createdAt={getDiffFromPast(createdAt)} />

                <div className="ThreaterContainer__InfoContainer__Description">
                    {content ? <p>{content}</p> : <p className="ThreaterContainer__InfoContainer__Description__Add">Add description</p>}
                </div>

                <GridImageContainer images={images.map(image => noAuthStatic(`/images/others/${image.id}`, { size: "origin" }))} />

                <div>
                    <FeedbackSummary reactors={reactors} commentCount={commentCount} reacted={reacted} reactCount={reactCount} displayCommentCount={comments.length} />
                </div>

                <div>
                    <PostAction onReactClick={this.onReactClick} reacted={reacted} onCommentClick={this.onCommentButtonClick} />
                </div>

                <div className="ThreaterContainer__InfoContainer__CommentContainer">
                    {comments.length < commentCount &&
                        <div onClick={this.onMoreCommentClick} style={{ cursor: "pointer" }}>
                            <Icon name="reply" color="blue" className="MoreCommentIcon" />
                            <ClickableText>See more comments</ClickableText>
                            {isLoadingMoreComment && <Loading />}
                        </div>}
                    {comments.map((comment) => <Comment {...comment} key={comment.id} replaceComment={this.replaceComment} commentableID={commentableID} />)}
                </div>

                <div>
                    <PostInput inverted onSubmit={this.onCommentSubmit} ref={el => this.postInput = el} isCommentInput style={{fontSize: "1.25em"}} placeholder="Write your comment ..." />
                </div>

            </div>
        )
    }
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

export default connect(mapStateToProps)(Post);