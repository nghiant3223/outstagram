import React, { Component } from 'react';

import GridImageContainer from '../GridImageContainer/GridImageContainer';
import Avatar from '../Avatar/Avatar';
import Comment from '../Comment/Comment';
import FeedbackSummary from '../FeedbackSummary/FeedbackSummary';
import PostAction from '../PostAction/PostAction';
import CommentInput from '../CommentInput/CommentInput';

import "./Post.css";
import PostHeader from '../PostHeader/PostHeader';

class Post extends Component {
    render() {
        const { images } = this.props;

        return (
            <div className="Post">

                <PostHeader fullname="Trọng Nghĩa" createdAt="2 hour ago" />

                <div className="ThreaterContainer__InfoContainer__Description">
                    {/* <p className="ThreaterContainer__InfoContainer__Description__Add">Add description</p> */}
                    <p>This is the description</p>
                </div>

                <GridImageContainer images={images} />


                <div>
                    <FeedbackSummary />
                </div>

                <div>
                    <PostAction />
                </div>

                <div className="ThreaterContainer__InfoContainer__CommentContainer">
                    <Comment />
                    <Comment />
                    <Comment />
                    <Comment />
                    <Comment />
                    <Comment />
                    <Comment />
                </div>

                <div>
                    <CommentInput inverted/>
                </div>



            </div>
        )
    }
}

export default Post;