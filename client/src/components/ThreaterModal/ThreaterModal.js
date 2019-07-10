import React, { Component } from 'react';
import { connect } from "react-redux";
import { Modal, Icon } from 'semantic-ui-react';

import "./ThreaterModal.css";
import AmpImage from '../AmpImage/AmpImage';
import Avatar from '../Avatar/Avatar';
import Comment from '../Comment/Comment';
import FeedbackSummary from '../FeedbackSummary/FeedbackSummary';
import PostAction from '../PostAction/PostAction';
import CommentInput from '../CommentInput/CommentInput';

import * as threaterAction from "../../actions/threater.modal";
import PostHeader from '../PostHeader/PostHeader';

class ThreaterModal extends Component {
    render() {
        const { isModalOpen, closeModal } = this.props;

        return (
            <Modal size="fullscreen" basic className="ThreaterModal"
                closeOnEscape
                centered
                closeOnDimmerClick
                open={isModalOpen}
                onClose={closeModal}>
                <div className="ThreaterContainer">
                    <div className="ThreaterContainer__ImageContainer">
                        <AmpImage src="https://unsplash.it/1000/400" fitType="contain" container="auto" />
                        <div className="ThreaterContainer__ImageContainer__Navigation ThreaterContainer__ImageContainer__Navigation--Prev"><Icon name="chevron left" size="big" color="grey" inverted /></div>
                        <div className="ThreaterContainer__ImageContainer__Navigation ThreaterContainer__ImageContainer__Navigation--Next"><Icon name="chevron right" size="big" color="grey" inverted /></div>
                    </div>

                    <div className="ThreaterContainer__InfoContainer">
                        <PostHeader fullname="Trọng Nghĩa" createdAt={"5 minute ago"} />

                        <div className="ThreaterContainer__InfoContainer__Description">
                            {/* <p className="ThreaterContainer__InfoContainer__Description__Add">Add description</p> */}
                            <p>This is the description</p>
                        </div>

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
                            <CommentInput inverted={false} />
                        </div>
                    </div>
                </div>
            </Modal>
        )
    }
}

const mapStateToProps = ({ threaterReducer: { isModalOpen } }) => ({ isModalOpen });

const mapDispatchToProps = (dispatch) => ({
    closeModal: () => dispatch(threaterAction.closeModal())
});


export default connect(mapStateToProps, mapDispatchToProps)(ThreaterModal);