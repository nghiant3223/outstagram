
import React, { Component } from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router';
import { Link } from 'react-router-dom';

import Avatar from '../../Avatar/Avatar';
import { toHHMM } from "../../../utils/time";
import { groupChatName } from '../../../utils/api';

import "./Contact.css";
import Socket from '../../../Socket';

class Contact extends Component {
    constructor(props) {
        super(props);

        this.state = {
            newMessage: false
        }
    }

    componentDidUpdate(prevProps) {
        if (this.roomMatch(this.props) && !this.roomMatch(prevProps)) {
            this.setState({ newMessage: false });
        }
    }

    roomMatch(props) {
        const { partner, id } = props;
        const { type: isGroupChat } = props;
        const { roomIdOrUsername } = props.match.params;
        const usernameMatch = !isGroupChat && partner.username === roomIdOrUsername;
        const idMatch = isGroupChat && id === roomIdOrUsername;
        return usernameMatch || idMatch;
    }

    componentDidMount() {
        console.log("mount")
        Socket.on("ROOM.SERVER.SEND_MESSAGE", this.onMessageReceive);
    }

    componentWillUnmount() {
        Socket.removeListener("ROOM.SERVER.SEND_MESSAGE", this.onMessageReceive);
    }

    onMessageReceive = (message) => {
        const { id } = this.props;
        const { data } = message;

        if (data.targetRoomID !== id) return;

        if (!this.roomMatch(this.props)) {
            this.setState({ newMessage: true });
        }
    }

    render() {
        const { newMessage } = this.state;
        const { partner, lastMessage, type, user, id, members } = this.props;
        const isGroupChat = type;

        let className = "ContactItemContainer__ContactItem";
        if (this.roomMatch(this.props)) className = className + " " + "ContactItemContainer__ContactItem ContactItemContainer__ContactItem--Active";
        if (newMessage) className = className + " " + "ContactItemContainer__ContactItem--NewMessage";

        return (
            <Link to={`/messages/${isGroupChat ? id : partner.username}`}>
                <div className={className}>
                    <div>
                        <Avatar userID={partner.id} />
                    </div>
                    <div className="ContactItemContainer__ContactItem__Content">
                        <div className="Fullname">{isGroupChat ? groupChatName(members) : partner.fullname}</div>
                        {lastMessage && <div className="ContactItemContainer__ContactItem__Content__LastMessage">{user.id === lastMessage.authorID && "You:"}&nbsp;{lastMessage.content}</div>}
                    </div>
                    {lastMessage && <div className="ContactItemContainer__ContactItem__Time">{toHHMM(new Date(lastMessage.createdAt))}</div>}
                </ div>
            </Link>
        )
    }
}


const mapStateToProps = ({ authReducer: { user } }) => ({ user });

export default withRouter(connect(mapStateToProps)(Contact));