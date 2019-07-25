import React, { Component } from 'react';
import { connect } from 'react-redux';

import { renderMessages } from '../../../utils/dom';
import { genUID } from '../../../utils/lang';

import { Message } from '../../../models/message';
import ContainerContext from './ContainerContext';

import * as roomServices from "../../../services/room.service";

import Loading from '../../Loading/Loading';
import MessageTyping from './MessageTyping/MessageTyping';
import "./ChatboxContainer.css";

import Socket from '../../../Socket';

class ChatboxContainer extends Component {
    constructor(props) {
        super(props);
        this.state = {
            isLoading: false,
            messageContent: '',
            messages: props.messages,
            roomID: props.roomID
        }
    }

    shouldScrollTop = false;
    shouldScrollBottom = true;

    componentWillReceiveProps(nextProps) {
        // When active room changes
        if (this.props.messages !== nextProps.messages) {
            this.setState({ messages: nextProps.messages, roomID: nextProps.roomID });
        }
    }

    componentDidUpdate(prevProps, prevState) {
        // When active rooms change
        if (this.props.roomIdOrUsername !== prevProps.roomIdOrUsername) {
            this.scrollToBottom();
            this.messageInput.focus();
        }

        //When active room does not change and more messages are fetched
        if (this.props.roomIdOrUsername === prevProps.roomIdOrUsername && this.state.messages.length > prevState.messages.length) {
            // IMPORTANT: Dont remove setTimeout
            // Set space between message conatiner top and scrollbar
            if (this.shouldScrollTop) {
                setTimeout(this.scrollToTop, 0);
            }

            if (this.shouldScrollBottom) {
                this.scrollToBottom();
            }
        }
    }

    componentDidMount() {
        this.scrollToBottom();
        this.messageInput.focus();
        Socket.on("ROOM.SERVER.SEND_MESSAGE", this.onMessageReceive);
    }

    componentWillUnmount() {
        Socket.removeListener("ROOM.SERVER.SEND_MESSAGE", this.onMessageReceive);
    }

    onMessageReceive = (message) => {
        const { roomID } = this.state;
        const { data } = message;

        if (roomID != data.targetRoomID) return;

        const receivedMessage = new Message(genUID(), data.authorID, data.content);
        this.setState((prevState) => ({ messages: [...prevState.messages, receivedMessage] }));
        this.shouldScrollBottom = true;
    }

    scrollToBottom() {
        this.chatboxContainer.scrollTop = this.chatboxContainer.scrollHeight;
    }

    onMessageConainerScroll = async () => {
        const { messages } = this.state;
        const { roomIdOrUsername } = this.props;
        const { scrollTop } = this.chatboxContainer;

        if (scrollTop == 0) {
            this.setState({ isLoading: true });
            try {
                const { data: { data: { messages: fetchMessages, roomID } } } = await roomServices.getMessages(roomIdOrUsername, 20, messages.length);
                if (fetchMessages !== null) {
                    this.shouldScrollTop = true;
                    this.setState((prevState) => ({ messages: [...fetchMessages, ...prevState.messages], roomID }));
                }
            } catch (e) {
                console.log("Cannot fetch more message", e);
            } finally {
                this.setState({ isLoading: false });
            }
        }
    }

    onFormSubmit = (e) => {
        e.preventDefault();

        const { user } = this.props;
        const messageContent = this.messageInput.value;
        const newMessage = new Message(genUID(), user.id, messageContent, true);

        this.setState((prevState) => ({ messages: [...prevState.messages, newMessage] }));
        this.messageInput.value = "";
    }

    // Replace the temporary message by the newly created message
    replaceMessage = (uid, newCreatedMessage) => {
        const { messages } = this.state;
        const message = messages.find(message => message.id === uid);

        if (!message) {
            throw new Error("Message does not exist");
        }

        // Copy property from newCreatedMessage to current message in the state;
        for (var k in newCreatedMessage) {
            // IMPORTANT: Ignore id field to prevent changing Message's key, which cause a new Message is created
            if (k !== "id") {
                message[k] = newCreatedMessage[k];
            }
        }

        this.scrollToBottom();
    }

    scrollToTop = () => {
        this.chatboxContainer.scrollTop = 10;
        this.shouldScrollTop = false;
    }

    render() {
        const { user, roomIdOrUsername } = this.props;
        const { messages, isLoading, roomID } = this.state;

        return (
            <ContainerContext.Provider value={{ replaceMessage: this.replaceMessage, roomIdOrUsername: roomIdOrUsername, roomID }}>
                <div className="ChatboxContainer">
                    <div className="ChatboxContainer__ChatboxContainer" onScroll={this.onMessageConainerScroll} ref={el => this.chatboxContainer = el}>
                        {isLoading && <div className="ChatboxContainer__ChatboxContainer__Loader"><Loading /></div>}
                        <div style={{ padding: "0.5em" }}>
                            {renderMessages(messages, user.id)}

                            <MessageTyping />
                        </div>
                    </div>

                    <form className="ChatboxContainer__InputContainer" onSubmit={this.onFormSubmit}>
                        <input className="ChatboxContainer__InputContainer__Input" placeholder="Type message..." ref={el => this.messageInput = el} />
                        <div className="ChatboxContainer__InputContainer__SendBtn">
                            <button>send</button>
                        </div>
                    </form>
                </div>
            </ContainerContext.Provider>
        )
    }
}

const mapStateToProps = ({ authReducer: { user } }) => ({ user });

export default connect(mapStateToProps)(ChatboxContainer);