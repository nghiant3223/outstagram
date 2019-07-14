import React, { Component } from 'react';
import { connect } from "react-redux";
import { Modal, Icon } from 'semantic-ui-react';

import "./ThreaterModal.css";
import AmpImage from '../AmpImage/AmpImage';
import Avatar from '../Avatar/Avatar';
import Comment from '../Comment/Comment';
import FeedbackSummary from '../FeedbackSummary/FeedbackSummary';
import PostAction from '../PostAction/PostAction';
import PostInput from '../PostInput/PostInput';

import * as threaterAction from "../../actions/threater.modal";
import PostHeader from '../PostHeader/PostHeader';
import { noAuthStatic } from '../../axios';

class ThreaterModal extends Component {
    state = {
        currentIndex: -1
    }

    componentWillReceiveProps(nextProps) {
        if (this.props.post !== nextProps.post && nextProps.post !== undefined) {
            this.setState({ currentIndex: nextProps.post.currentIndex })
            console.log(nextProps.post)
        }
    }

    onNextClick = () => {
        const { post: { images } } = this.props;
        this.setState((prevState) => ({ currentIndex: (prevState.currentIndex + 1) % images.length }));
    }

    onPrevClick = () => {
        const { post: { images } } = this.props;
        this.setState((prevState) => ({ currentIndex: (prevState.currentIndex - 1) % images.length }));
    }

    render() {
        const { isModalOpen, close, post } = this.props;
        const { currentIndex } = this.state;
        const currentPostImage = post.images[currentIndex]

        return (
            <Modal className="ThreaterModal"
                size="large"
                basic
                closeOnEscape
                centered
                closeOnDimmerClick
                open={isModalOpen}
                onClose={close}>
                <i className="ThreaterModal__CloseIcon close icon" onClick={close}></i>
                {post &&
                    <div className="ThreaterContainer">
                        <div className="ThreaterContainer__ImageContainer" id="x">
                            <AmpImage src={noAuthStatic(`/images/others/${currentPostImage.id}?size=origin`)} fit="contain" container="auto" />

                            <div className="ThreaterContainer__ImageContainer__Navigation ThreaterContainer__ImageContainer__Navigation--Prev"
                                onClick={this.onPrevClick}>
                                <Icon name="chevron left" size="big" color="grey" inverted />
                            </div>
                            <div className="ThreaterContainer__ImageContainer__Navigation ThreaterContainer__ImageContainer__Navigation--Next"
                                onClick={this.onNextClick}>
                                <Icon name="chevron right" size="big" color="grey" inverted />
                            </div>
                        </div>

                        <div className="ThreaterContainer__InfoContainer">
                            <PostHeader fullname="Trọng Nghĩa" createdAt={"5 minute ago"} />

                            <div className="ThreaterContainer__InfoContainer__Description">
                                {currentPostImage.content ?
                                    <p className="ThreaterContainer__InfoContainer__Description__Add">Add description</p>
                                    :
                                    <p>This is the description</p>}
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
                                <PostInput inverted={true} />
                            </div>
                        </div>
                    </div>}
            </Modal>
        )
    }
}

const mapStateToProps = ({ threaterReducer: { isModalOpen, onDisplayPost: post } }) => ({ isModalOpen, post });

const mapDispatchToProps = (dispatch) => ({
    close: () => dispatch(threaterAction.closeModal())
});


export default connect(mapStateToProps, mapDispatchToProps)(ThreaterModal);