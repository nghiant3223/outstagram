import React, { Component } from 'react';
import { withRouter } from 'react-router';
import { connect } from 'react-redux';

import ContactInfo from './ContactInfo/ContactInfo';
import ChatboxContainer from './ChatboxContainer/ChatboxContainer';

import * as roomServices from "../../services/room.service";
import * as userServices from "../../services/user.service";
import ConversationPlaceholder from './Placeholder';
import { groupChatName } from '../../utils/api';
import { genUID } from '../../utils/lang';

class ConversationContainer extends Component {
    state = {
        room: null,
        isLoading: false
    }

    componentDidMount() {
        this.fetchRoom(true);
    }

    componentDidUpdate(prevProps) {
        if (this.props.match.params.roomIdOrUsername !== prevProps.match.params.roomIdOrUsername) {
            this.fetchRoom(false);
        }
    }

    async fetchRoom(shouldAddFakeRoom) {
        const { roomIdOrUsername } = this.props.match.params;
        if (!roomIdOrUsername) return;

        const { user } = this.props;
        if (roomIdOrUsername === user.username) return;

        try {
            const { data: { data: room } } = await roomServices.getRoom(roomIdOrUsername);
            this.setState({ room });
        } catch (e) {
            if (e.response.data !== null) {
                const { data: { type } } = e.response.data;
                if (type === "room_not_created") {
                    userServices.getUser(roomIdOrUsername)
                        .then(({ data: { data: { user } } }) => {
                            const fakeRoom = { id: genUID(), partner: user, type: false, isNew: true, messages: [] }
                            // Only add fake room on first load
                            if (shouldAddFakeRoom) this.props.addFakeRoom(fakeRoom);
                            this.setState({ room: fakeRoom });
                        })
                        .catch((e) => console.log("Cannot fetch user", e));
                }
            }
        } finally {
            this.setState({ isLoading: false });
        }
    }

    onNewRoomCreated = (fakeRoomID, newRoom) => {
        this.setState({ room: newRoom });
        const { replaceFakeRoom } = this.props;
        replaceFakeRoom(fakeRoomID, newRoom);
    }

    render() {
        const { room, isLoading } = this.state
        if (isLoading) return <ConversationPlaceholder />;
        if (!room) return <div className="MessageInfoContainer" />;

        const { user } = this.props;
        const { roomIdOrUsername } = this.props.match.params;
        if (roomIdOrUsername == user.username) return <div className="MessageInfoContainer" />;

        const { updateLastMessage } = this.props;
        const { partner, members, messages, id, type: isGroupChat, isNew } = room;

        return (
            <div className="MessageInfoContainer">
                <ContactInfo header={isGroupChat ? groupChatName(members) : partner.fullname} partner={partner} key={partner.id} />
                <ChatboxContainer messages={messages} roomID={id} partner={partner} onNewRoomCreated={this.onNewRoomCreated} roomIdOrUsername={roomIdOrUsername} updateLastMessage={updateLastMessage} isNew={isNew} />
            </div>
        )
    }
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

export default withRouter(connect(mapStateToProps)(ConversationContainer));