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
import * as postServices from "../../services/post.service";
import PostHeader from '../PostHeader/PostHeader';
import { noAuthStatic } from '../../axios';
import Post from '../Post/Post';

const initialState = {
    isLoading: false,
    currentIndex: -1,
    postOrPostImage: undefined,
}

class ThreaterModal extends Component {
    state = {
        initialState
    }

    componentWillReceiveProps(nextProps) {
        if (nextProps.post !== undefined) {
            const { currentIndex } = nextProps.post
            this.setState({ currentIndex });
        }
    }

    componentDidUpdate(prevProps, prevState) {
        if (this.props.post === undefined) {
            // Do not update if update (1) occurs
            if (this.state.currentIndex == -1) {
                return;
            }

            // Reset to the initial state (1)
            this.setState(initialState);
            return;
        }

        if (this.props.post !== prevProps.post || this.state.currentIndex !== prevState.currentIndex) {
            const { images } = this.props.post;
            const { currentIndex } = this.state

            this.setState({ isLoading: true });

            if (images.length > 1) {
                postServices.getPostImage(images[currentIndex].id)
                    .then(({ data: { data: posts } }) => {
                        this.setState({ postOrPostImage: posts, isLoading: false });
                    }).catch((e) => {
                        this.setState({ isLoading: false });
                        console.warn("Failed to get post's image", e);
                    });
            } else {
                postServices.getSpecificPost(images[0].id)
                    .then(({ data: { data: posts } }) => {
                        this.setState({ postOrPostImage: posts, isLoading: false });
                    }).catch((e) => {
                        this.setState({ isLoading: false });
                        console.warn("Failed to get post's image", e);
                    });
            }

        }
    }

    onNextClick = () => {
        const { post: { images } } = this.props;
        this.setState((prevState) => ({ currentIndex: (prevState.currentIndex + 1) % images.length }));
    }

    onPrevClick = () => {
        const { post: { images } } = this.props;
        this.setState((prevState) => ({ currentIndex: Math.abs((prevState.currentIndex - 1) % images.length) }));
    }

    render() {
        const { isModalOpen, close } = this.props;
        const { postOrPostImage } = this.state;

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
                {postOrPostImage &&
                    <div className="ThreaterContainer">
                        <div className="ThreaterContainer__ImageContainer" id="x">
                            <AmpImage src={noAuthStatic(`/images/others/${postOrPostImage.imageID}?size=origin`)} fit="contain" container="auto" />

                            <div className="ThreaterContainer__ImageContainer__Navigation ThreaterContainer__ImageContainer__Navigation--Prev"
                                onClick={this.onPrevClick}>
                                <Icon name="chevron left" size="big" color="grey" inverted />
                            </div>
                            <div className="ThreaterContainer__ImageContainer__Navigation ThreaterContainer__ImageContainer__Navigation--Next"
                                onClick={this.onNextClick}>
                                <Icon name="chevron right" size="big" color="grey" inverted />
                            </div>
                        </div>
                        <Post {...postOrPostImage} showImageGrid={false} key={postOrPostImage.id} isPost={this.props.post && this.props.post.images.length == 1} />
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