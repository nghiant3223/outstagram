import React, { Component } from 'react';
import { connect } from 'react-redux'
import { Link } from 'react-router-dom'
import { Modal, Icon } from 'semantic-ui-react';
import pluralize from 'pluralize';

import Avatar from "../Avatar/Avatar";

import * as reactorActions from "../../actions/reactor.action";
import * as reactableServices from "../../services/reactable.service";

import "./ReactorModal.css";
import FollowButton from '../FollowButton/FollowButton';
import ClickableText from '../ClickableText/ClickableText';
import Loading from '../Loading/Loading';

const initialState = {
    isLoading: false,
    reactors: [],
    reactCount: 0
}

class ReactorModal extends Component {
    state = { ...initialState }

    componentDidUpdate(prevProps) {
        const { isModalOpen, reactableID } = this.props

        if (isModalOpen && !prevProps.isModalOpen) {
            this.fetchReactions(reactableID);
        }

        if (!isModalOpen && prevProps.isModalOpen) {
            this.setState({ ...initialState })
        }
    }

    async fetchReactions(id) {
        this.setState({ isLoading: true });

        try {
            const { data: { data: { reactors, reactCount } } } = await reactableServices.getReactions(id, 100, 0);
            this.setState({ reactCount, reactors });
        } catch (e) {
            console.log("Fetch reactions failed", e);
        } finally {
            this.setState({ isLoading: false });
        }
    }

    render() {
        const { reactors, reactCount, isLoading } = this.state;
        const { isModalOpen, close, user } = this.props;

        return (
            <Modal
                closeOnEscape
                closeOnDimmerClick
                open={isModalOpen}
                centered={false}
                size="tiny"
                onClose={close}>
                <Modal.Header className="ReactorContainer__Header">
                    <ClickableText><Icon name="heart" color="red" inverted /> {reactCount} {pluralize("Reactions")}</ClickableText>
                </Modal.Header>
                <div className="ReactorContainer">
                    {isLoading ?
                        <div className="ReactorContainer__Loading"><Loading /></div>
                        :
                        <Modal.Content scrolling className="ReactorContainer__ReactorContainer">
                            {reactors.map((reactor) =>
                                <div className="ReactorContainer__Reactor" key={reactor.id}>
                                    <div className="ReactorContainer__Reactor__Container">
                                        <Avatar userID={reactor.id} />
                                        <div><Link to={`/${reactor.username}`} onClick={close}>{reactor.fullname}</Link></div>
                                    </div>
                                    {reactor.id !== user.id && <div><FollowButton userID={reactor.id} followed={reactor.followed} size="tiny" /></div>}
                                </div>)}
                        </Modal.Content>
                    }
                </div>
            </Modal>
        )
    }
}

const mapStateToProps = ({ reactorReducer: { isModalOpen, reactableID }, authReducer: { user } }) => ({ isModalOpen, reactableID, user });

const mapDispatchToProps = (dispatch) => ({
    close: () => dispatch(reactorActions.closeModal())
});

export default connect(mapStateToProps, mapDispatchToProps)(ReactorModal);