import React, { Component } from 'react';
import { Modal, Icon } from 'semantic-ui-react';

import "./ThreaterModal.css";
import AmpImage from '../AmpImage/AmpImage';
import Avatar from '../Avatar/Avatar';
import Comment from '../Comment/Comment';
import FeedbackSummary from '../FeedbackSummary/FeedbackSummary';
import PostAction from '../PostAction/PostAction';
import CommentInput from '../CommentInput/CommentInput';

class ThreaterModal extends Component {
    render() {
        return (
            <Modal size="fullscreen" basic className="ThreaterModal"
                closeOnEscape
                centered
                closeOnDimmerClick
                open={true}>
                <div className="ThreaterContainer">
                    <div className="ThreaterContainer__ImageContainer">
                        <AmpImage src="https://unsplash.it/1000/400" fitType="contain" container="auto" />
                        <div className="ThreaterContainer__ImageContainer__Navigation ThreaterContainer__ImageContainer__Navigation--Prev"><Icon name="chevron left" size="big" color="grey" inverted /></div>
                        <div className="ThreaterContainer__ImageContainer__Navigation ThreaterContainer__ImageContainer__Navigation--Next"><Icon name="chevron right" size="big" color="grey" inverted /></div>
                    </div>

                    <div className="ThreaterContainer__InfoContainer">
                        <div className="ThreaterContainer__InfoContainer__Header">
                            <div className="ThreaterContainer__InfoContainer__Header__Avatar">
                                <Avatar width="3em" />
                            </div>

                            <div className="ThreaterContainer__InfoContainer__Header__Info">
                                <div className="ThreaterContainer__InfoContainer__Header__Info__Fullname">Fullname</div>
                                <div className="ThreaterContainer__InfoContainer__Header__Info__CreatedAt">1 minute ago</div>
                            </div>
                        </div>

                        <div className="ThreaterContainer__InfoContainer__Description">
                            {/* <p className="ThreaterContainer__InfoContainer__Description__Add">Add description</p> */}
                            <p>This is the description</p>
                        </div>

                        <FeedbackSummary />

                        <PostAction />

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
                            <CommentInput />
                        </div>
                    </div>
                </div>
            </Modal>
        )
    }
}


export default ThreaterModal;