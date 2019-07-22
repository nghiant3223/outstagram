import React from 'react';
import { connect } from 'react-redux';
import { withRouter } from 'react-router';
import { Link } from 'react-router-dom';

import Avatar from '../../Avatar/Avatar';
import { toHHMM } from "../../../utils/time";

import "./Contact.css";
import { groupChatName } from '../../../utils/api';

function Contact(props) {
    const { roomIdOrUsername } = props.match.params;
    const { partner, lastMessage, type, user, id, members } = props;
    const isGroupChat = type;

    const usernameMatch = !isGroupChat && partner.username === roomIdOrUsername
    const idMatch = isGroupChat && id === roomIdOrUsername;
    if (usernameMatch || idMatch) {
        var className = "ContactItemContainer__ContactItem ContactItemContainer__ContactItem--Active";
    } else {
        var className = "ContactItemContainer__ContactItem";
    }

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

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

export default withRouter(connect(mapStateToProps)(Contact));